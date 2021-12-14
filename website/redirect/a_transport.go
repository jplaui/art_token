package redirect

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type redirectContext struct {
	HttpsPort string
	Logger    log.Logger
}

// httpsRedirect redirects http requests to https
func (rc *redirectContext) httpsRedirect(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ":")[0]
	level.Info(rc.Logger).Log("msg", "redirect http to https")
	http.Redirect(
		w, r,
		fmt.Sprintf("https://%s:%s%s", host, rc.HttpsPort, r.URL.String()),
		http.StatusMovedPermanently,
	)
}

// router very simple, no endpoint definition in this case
func NewRouter(logger log.Logger, httpsPort string) http.Handler {
	redirectContext := &redirectContext{Logger: logger, HttpsPort: httpsPort}
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", redirectContext.httpsRedirect)
	return mux1
}
