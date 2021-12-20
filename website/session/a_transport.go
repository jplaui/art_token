package session

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

// set in middleware, read in service method AfterEndpointCall
type SessionIdContextKey string

// wrong input data
var ErrBadRequest = errors.New("Bad request")

// handler not found
var ErrNotFound = errors.New("Resource not found")

// handler not found
var ErrInternalServer = errors.New("Internal server error")

// custom error interface which allows to pass service logic errors back to the client
type customError interface {
	error() error
}

// encodes customError to JSON and writes error to responseWriter
func encodeError(_ context.Context, err error, w http.ResponseWriter) {

	errorType := errorStatusCode(err)
	w.WriteHeader(errorType)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})

}

// get error status code for predefined error types
// returns notFound if error cannot be mapped to bad request or internal server
func errorStatusCode(err error) int {

	switch err {

	case ErrBadRequest:
		return http.StatusBadRequest

	case ErrInternalServer:
		return http.StatusInternalServerError

	default:
		return http.StatusNotFound
	}
}

// does not decode request
func decodeNone(_ context.Context, r *http.Request) (interface{}, error) { return r, nil }

// decode create session
func decodeCreateSession(ctx context.Context, r *http.Request) (interface{}, error) {

	// create session request, defined in endpoints file
	var cs CreateSessionRequest

	// JSON decoding
	err := json.NewDecoder(r.Body).Decode(&cs)
	if err != nil {
		return nil, ErrBadRequest
	}

	return cs, nil
}

// encodes response after endpoint handlers are done with writing to responseWriter
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if err, ok := response.(customError); ok && err.error() != nil {
		encodeError(ctx, err.error(), w)
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

// defines multiplexer routes
// maps endpoint handlers to routes
// add middleware functions on transport layer
// info: make sure to attach session routes first in main.go as they check if users have active sessions
func AttachRoutes(router *mux.Router, sc *securecookie.SecureCookie, ctx context.Context, s Service, logger log.Logger) http.Handler {

	e := MakeEndpoints(s)
	options := []httptransport.ServerOption{

		// executed after enpoint invocation/call
		// function executed on response writer
		// service.AfterEndpointCall sets or revokes cookies
		httptransport.ServerAfter(s.AfterEndpointCall),

		// allows more fine grained error handling and logging than httptransport.ServerErrorHandler
		httptransport.ServerErrorEncoder(encodeError),
	}

	createSessionHandler := httptransport.NewServer(
		e.CreateSession,
		decodeCreateSession,
		encodeResponse,
		options...,
	)

	deleteSessionHandler := httptransport.NewServer(
		e.DeleteSession,
		decodeNone,
		encodeResponse,
		options...,
	)

	router.Handle("/login", createSessionHandler).Methods("POST")
	router.Use(jsonMDW)
	router.Use(setSessionIdContextMDW(ctx, s, sc, logger))
	router.Handle("/logout", deleteSessionHandler).Methods("GET")

	return router
}

// makes sure that login request and responses of session calls are encoded in JSON
func jsonMDW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// decodes cookie and sets session identifier into context
// required by after endpoint call function to identify user
func setSessionIdContextMDW(ctx context.Context, s Service, ss *securecookie.SecureCookie, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// log level
			logger := log.With(logger, "middleware", "setSessionIdContextMDW")

			// check if request contains session_cookie
			if cookie, err := r.Cookie("session_cookie"); err == nil {

				// cookie decoding
				value := make(map[string]string)
				if err = ss.Decode("session_cookie", cookie.Value, &value); err == nil {

					// extract session identifier from cookie and write to context
					var sessionId string
					sessionId = value["session_id"]
					key := SessionIdContextKey("session_id")
					context.WithValue(ctx, key, sessionId)

					// update session expiry time if session active
					// info: applies to all active session requests hitting the page
					if session, err := s.ReadSession(sessionId); err == nil {

						// check expiry
						if time.Now().Before(session.ExpiresAt) {

							// update
							s.UpdateSession(sessionId, session)
							level.Info(logger).Log("session MDW", "updated")
						}
					}

				}
			}

			// if no cookie or no session id exist, do nothing and pass request along
			next.ServeHTTP(w, r)
			return
		})
	}
}
