package session

import (
	"context"
	"crypto/md5"
	"net/http"
	"strings"
	"time"

	"website/storage"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/badoux/checkmail"
	"github.com/gorilla/securecookie"
)

type ContextKey2 string

// service interface defining all required methods
type Service interface {
	CreateSession(ctx context.Context, email string, password string) error
	ReadSession(sessionId string) (Session, error)
	UpdateSession(sessionId string, session Session) error
	DeleteSession(sessionId string) error
	AfterEndpointCall(ctx context.Context, w http.ResponseWriter) context.Context
}

// service struct implementing service interface with attributes
type service struct {
	userStore     storage.UserStore
	sessionStore  SessionStore
	secretSession *securecookie.SecureCookie
	logger        log.Logger
}

// service struct create session method
func (s *service) CreateSession(ctx context.Context, email string, password string) error {

	// logger level
	logger := log.With(s.logger, "method", "CreateSession")

	// remove whitespaces
	trimmedEmail := strings.TrimSpace(email)
	trimmedPassword := strings.TrimSpace(password)

	// verify email
	if err := checkmail.ValidateFormat(trimmedEmail); err != nil {
		level.Error(logger).Log("checkmail.ValidateFormat:", err)
		return err
	}

	// load user from file
	user, err := s.userStore.ReadUser(trimmedEmail)
	if err != nil {
		level.Error(logger).Log("s.userStore.Readuser:", err)
		return err
	}

	// verify password
	err = storage.CheckPasswordHash(trimmedPassword, user.Password)
	if err != nil {
		level.Error(logger).Log("CheckPasswordHash:", err)
		return err
	}

	// create session identifier, hash user uuid
	hash := md5.Sum([]byte(user.Email))
	fileHash := string(hash[:])
	hash = md5.Sum([]byte(user.UUID))
	sessionId := string(hash[:])

	// create session
	err = s.sessionStore.CreateSession(sessionId, fileHash)
	if err != nil {
		level.Error(logger).Log("s.sesionStore.CreateSession:", err)
		return err
	}

	// if session creation successfull, set session identifier in context
	// if not set, service.AfterEndpointCall cannot read session identifier and set cookie
	key := SessionIdContextKey("session_id")
	context.WithValue(ctx, key, sessionId)

	return nil

}

// service struct read session method
func (s *service) ReadSession(sessionId string) (session Session, err error) {

	// logger level
	logger := log.With(s.logger, "method", "ReadSession")

	// read session from store
	session, err = s.sessionStore.ReadSession(sessionId)
	if err != nil {
		level.Error(logger).Log("s.sessionStore.ReadSession:", err)
	}

	return session, nil
}

// service struct update session method
func (s *service) UpdateSession(sessionId string, session Session) error {

	// logger level
	logger := log.With(s.logger, "method", "UpdateSession")

	// update session with new times
	err := s.sessionStore.UpdateSession(sessionId, session)
	if err != nil {
		level.Error(logger).Log("s.sessionStore.UpdateSession:", err)
		return err
	}

	return nil
}

// service struct delete session method
func (s *service) DeleteSession(sessionId string) error {

	// logger level
	logger := log.With(s.logger, "method", "DeleteSession")

	// delete session
	err := s.sessionStore.DeleteSession(sessionId)
	if err != nil {
		level.Error(logger).Log("s.sessionStore.DeleteSession:", err)
		return err
	}

	return nil
}

// initialization function to return service struct
// this function should is called in main.go
func NewService(userStore storage.UserStore, sessionStore SessionStore, secretSession *securecookie.SecureCookie, logger log.Logger) Service {
	return &service{
		userStore:     userStore,
		sessionStore:  sessionStore,
		secretSession: secretSession,
		logger:        logger,
	}
}

// cookie management
// function executes after enpoint call is done
func (s *service) AfterEndpointCall(ctx context.Context, w http.ResponseWriter) context.Context {

	// create cookie based on session state
	var cookie *http.Cookie

	// access session identifier from context
	// make sure session_id context set in createSession if request does not contains a cookie
	k := SessionIdContextKey("session_id")
	sessionId, ok := ctx.Value(k).(string)

	// only set activate cookie if session_id allows to search for active session
	if ok {

		// fetch session
		session, err := s.ReadSession(sessionId)
		if err == nil {

			// encode session value for cookie
			value := map[string]string{"session_id": sessionId}
			valueEncoded, err := s.secretSession.Encode("session_cookie", value)
			if err == nil {

				// session found, check if session valid
				if time.Now().Before(session.ExpiresAt) {

					// session not expired and still active
					cookie = generateCookie(true, valueEncoded)
					http.SetCookie(w, cookie)

					// return if cookie has been set
					return ctx
				}
			}
		}
	}

	// remove session cookie
	cookie = generateCookie(false, "")
	http.SetCookie(w, cookie)
	return ctx
}

// creates cookie
func generateCookie(set bool, value string) *http.Cookie {

	// activate cookie if set==true
	var cookie *http.Cookie

	if set {

		// set cookie value
		cookie = &http.Cookie{
			Name:     "session_cookie",
			Value:    value,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}

	} else {

		// remove cookie from browser client
		cookie = &http.Cookie{
			Name:     "session_cookie",
			Value:    "",
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			MaxAge:   -1,
			Expires:  time.Now().Add(-10 * time.Hour),
		}
	}

	return cookie
}
