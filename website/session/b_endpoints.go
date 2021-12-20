package session

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type CreateSessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BoolResponse struct {
	Data bool  `json:"data"`
	Err  error `json:"errors"`
}

// have BoolResponse follow the customError interface defined in a_transport.go
func (r BoolResponse) error() error { return r.Err }

// Endpoints struct groups all endpoint handlers and could be used to store further attributes such as logger, entpoint state etc.
// endpoint handlers call and manage service logic
type Endpoints struct {
	CreateSession endpoint.Endpoint
	DeleteSession endpoint.Endpoint
}

// init Endpoints struct and maps functions to endpoint handlers
// this function is called in a_transport.go
func MakeEndpoints(s Service) Endpoints {

	return Endpoints{
		CreateSession: epCreateSession(s),
		DeleteSession: epDeleteSession(s),
	}
}

// create session endpoint
func epCreateSession(s Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		// convert request interface
		req := request.(CreateSessionRequest)

		// call service method
		err := s.CreateSession(ctx, req.Email, req.Password)
		if err != nil {
			return BoolResponse{Data: false, Err: err}, ErrInternalServer
		}

		return BoolResponse{Data: true, Err: nil}, nil
	}

}

// delete session endpoint
func epDeleteSession(s Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		// access cookie
		// r := request.(*http.Request)
		// cookie, err := r.Cookie("session_cookie")

		k := SessionIdContextKey("session_id")
		sessionId, ok := ctx.Value(k).(string)
		if !ok {
			return BoolResponse{Data: false, Err: errors.New("context conversion failed")}, ErrBadRequest
		}

		// call service method
		err := s.DeleteSession(sessionId)
		if err != nil {
			return BoolResponse{Data: false, Err: err}, ErrInternalServer
		}

		return BoolResponse{Data: true, Err: err}, nil
	}

}
