package main

import (
	"context"
	"net/http"
	"os"
	"website/configs"
	"website/redirect"
	"website/session"
	"website/spa"
	"website/storage"

	"github.com/fvbock/endless"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"golang.org/x/sync/errgroup"
)

var (
	logger        log.Logger
	g             errgroup.Group
	secretSession *securecookie.SecureCookie
	config        *configs.Config
)

func main() {

	// configs read in
	config = configs.GetConfig()

	// logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "art_token backend webserver",
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	level.Info(logger).Log("msg", "server start")
	defer level.Info(logger).Log("msg", "server stop")

	// HTTP handler
	mux1 := redirect.NewRouter(log.With(logger, "service", "HTTP server"), config.HttpsPort)
	httpAddr := config.HttpAddr + ":" + config.HttpPort
	serverHTTP := endless.NewServer(httpAddr, mux1)

	// run redirect server to enable default HTTPS
	g.Go(func() error {

		level.Info(logger).Log("msg", "1. server starting at address: ", httpAddr)
		err := serverHTTP.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Log("exit", "closing HTTP server; error:"+err.Error())
		}
		return err

	})

	// context for storage handlers
	ctx1 := context.Background()
	ctx1, cancel := context.WithCancel(ctx1)
	defer cancel()

	// context for session handlers
	ctx2 := context.Background()
	ctx2, cancel2 := context.WithCancel(ctx2)
	defer cancel2()

	// encrypted session cookie
	sEncryptCookie := securecookie.GenerateRandomKey(32)
	var blockKey = []byte(sEncryptCookie)
	secretSession = securecookie.New([]byte(config.HttpSessionSecret), blockKey)

	// session service
	var svc2 session.Service
	{
		storageConfig := storage.UserStoreConfig{UsersPath: config.UsersPath}
		userStore := storage.NewUserStore(storageConfig, log.With(logger, "client", "storage"))
		sessionStore := session.NewSessionStore()
		svc2 = session.NewService(userStore, sessionStore, secretSession, log.With(logger, "service", "session"))
	}

	// // storage service
	// var svc1 storage.Service
	// {

	// }

	// HTTPS handler
	mux2 := mux.NewRouter()

	// attach services
	session.AttachRoutes(mux2, secretSession, ctx2, svc2, log.With(logger, "transport", "session"))
	// storage.AttachRoutes(mux2, ctx1, svc1, log.With(logger, "transport", "storage"))
	spa.AttachRoutes(mux2, config.StaticAssetsDir, log.With(logger, "transport", "spa"))

	// configure server
	httpsAddr := config.HttpsAddr + ":" + config.HttpsPort
	serverHTTPS := endless.NewServer(httpsAddr, mux2)

	// run HTTPS server and serve content
	g.Go(func() error {

		level.Info(logger).Log("msg", "2. server starting at address: ", httpsAddr)
		err := serverHTTPS.ListenAndServeTLS(config.TlsCertPath, config.TlsKeyPath)
		if err != nil && err != http.ErrServerClosed {
			logger.Log("exit", "closing HTTPS server; error: "+err.Error())
		}
		return err

	})

	if err := g.Wait(); err != nil {
		level.Info(logger).Log("msg", "server stopped after error grouping wait")
	}

}
