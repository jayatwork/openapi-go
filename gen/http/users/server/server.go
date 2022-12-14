// Code generated by goa v3.8.4, DO NOT EDIT.
//
// users HTTP server
//
// Command:
// $ goa gen openapi-go/design

package server

import (
	"context"
	"net/http"
	users "openapi-go/gen/users"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the users service endpoint HTTP handlers.
type Server struct {
	Mounts              []*MountPoint
	Create              http.Handler
	All                 http.Handler
	GenHTTPOpenapi3JSON http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the users service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *users.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	fileSystemGenHTTPOpenapi3JSON http.FileSystem,
) *Server {
	if fileSystemGenHTTPOpenapi3JSON == nil {
		fileSystemGenHTTPOpenapi3JSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Create", "POST", "/users"},
			{"All", "GET", "/users"},
			{"./gen/http/openapi3.json", "GET", "/openapi3.json"},
		},
		Create:              NewCreateHandler(e.Create, mux, decoder, encoder, errhandler, formatter),
		All:                 NewAllHandler(e.All, mux, decoder, encoder, errhandler, formatter),
		GenHTTPOpenapi3JSON: http.FileServer(fileSystemGenHTTPOpenapi3JSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "users" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Create = m(s.Create)
	s.All = m(s.All)
}

// Mount configures the mux to serve the users endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCreateHandler(mux, h.Create)
	MountAllHandler(mux, h.All)
	MountGenHTTPOpenapi3JSON(mux, goahttp.Replace("", "/./gen/http/openapi3.json", h.GenHTTPOpenapi3JSON))
}

// Mount configures the mux to serve the users endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountCreateHandler configures the mux to serve the "users" service "create"
// endpoint.
func MountCreateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/users", f)
}

// NewCreateHandler creates a HTTP handler which loads the HTTP request and
// calls the "users" service "create" endpoint.
func NewCreateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateRequest(mux, decoder)
		encodeResponse = EncodeCreateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create")
		ctx = context.WithValue(ctx, goa.ServiceKey, "users")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountAllHandler configures the mux to serve the "users" service "all"
// endpoint.
func MountAllHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/users", f)
}

// NewAllHandler creates a HTTP handler which loads the HTTP request and calls
// the "users" service "all" endpoint.
func NewAllHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeAllResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "all")
		ctx = context.WithValue(ctx, goa.ServiceKey, "users")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapi3JSON configures the mux to serve GET request made to
// "/openapi3.json".
func MountGenHTTPOpenapi3JSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi3.json", h.ServeHTTP)
}
