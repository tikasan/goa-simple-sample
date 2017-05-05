// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ActionsController is the controller interface for the Actions actions.
type ActionsController interface {
	goa.Muxer
	ID(*IDActionsContext) error
	Hello(*HelloActionsContext) error
	Ping(*PingActionsContext) error
}

// MountActionsController "mounts" a Actions resource controller on the given service.
func MountActionsController(service *goa.Service, ctrl ActionsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v1/actions/:ID", ctrl.MuxHandler("preflight", handleActionsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v1/actions/hello", ctrl.MuxHandler("preflight", handleActionsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v1/actions/ping", ctrl.MuxHandler("preflight", handleActionsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewIDActionsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ID(rctx)
	}
	h = handleActionsOrigin(h)
	service.Mux.Handle("GET", "/api/v1/actions/:ID", ctrl.MuxHandler("ID", h, nil))
	service.LogInfo("mount", "ctrl", "Actions", "action", "ID", "route", "GET /api/v1/actions/:ID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHelloActionsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Hello(rctx)
	}
	h = handleActionsOrigin(h)
	service.Mux.Handle("GET", "/api/v1/actions/hello", ctrl.MuxHandler("Hello", h, nil))
	service.LogInfo("mount", "ctrl", "Actions", "action", "Hello", "route", "GET /api/v1/actions/hello")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPingActionsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Ping(rctx)
	}
	h = handleActionsOrigin(h)
	service.Mux.Handle("GET", "/api/v1/actions/ping", ctrl.MuxHandler("Ping", h, nil))
	service.LogInfo("mount", "ctrl", "Actions", "action", "Ping", "route", "GET /api/v1/actions/ping")
}

// handleActionsOrigin applies the CORS response headers corresponding to the origin.
func handleActionsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// ArrayController is the controller interface for the Array actions.
type ArrayController interface {
	goa.Muxer
	Array(*ArrayArrayContext) error
}

// MountArrayController "mounts" a Array resource controller on the given service.
func MountArrayController(service *goa.Service, ctrl ArrayController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v1/array", ctrl.MuxHandler("preflight", handleArrayOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewArrayArrayContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Array(rctx)
	}
	h = handleArrayOrigin(h)
	service.Mux.Handle("GET", "/api/v1/array", ctrl.MuxHandler("Array", h, nil))
	service.LogInfo("mount", "ctrl", "Array", "action", "Array", "route", "GET /api/v1/array")
}

// handleArrayOrigin applies the CORS response headers corresponding to the origin.
func handleArrayOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// JsController is the controller interface for the Js actions.
type JsController interface {
	goa.Muxer
	goa.FileServer
}

// MountJsController "mounts" a Js resource controller on the given service.
func MountJsController(service *goa.Service, ctrl JsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/js/*filepath", ctrl.MuxHandler("preflight", handleJsOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/js/*filepath", "/js/")
	h = handleJsOrigin(h)
	service.Mux.Handle("GET", "/js/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Js", "files", "/js/", "route", "GET /js/*filepath")

	h = ctrl.FileHandler("/js/", "/js/index.html")
	h = handleJsOrigin(h)
	service.Mux.Handle("GET", "/js/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Js", "files", "/js/index.html", "route", "GET /js/")
}

// handleJsOrigin applies the CORS response headers corresponding to the origin.
func handleJsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// MethodController is the controller interface for the Method actions.
type MethodController interface {
	goa.Muxer
	Method(*MethodMethodContext) error
}

// MountMethodController "mounts" a Method resource controller on the given service.
func MountMethodController(service *goa.Service, ctrl MethodController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v1/method/get", ctrl.MuxHandler("preflight", handleMethodOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v1/method/post", ctrl.MuxHandler("preflight", handleMethodOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v1/method/delete", ctrl.MuxHandler("preflight", handleMethodOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v1/method/put", ctrl.MuxHandler("preflight", handleMethodOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewMethodMethodContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Method(rctx)
	}
	h = handleMethodOrigin(h)
	service.Mux.Handle("GET", "/api/v1/method/get", ctrl.MuxHandler("Method", h, nil))
	service.LogInfo("mount", "ctrl", "Method", "action", "Method", "route", "GET /api/v1/method/get")
	service.Mux.Handle("POST", "/api/v1/method/post", ctrl.MuxHandler("Method", h, nil))
	service.LogInfo("mount", "ctrl", "Method", "action", "Method", "route", "POST /api/v1/method/post")
	service.Mux.Handle("DELETE", "/api/v1/method/delete", ctrl.MuxHandler("Method", h, nil))
	service.LogInfo("mount", "ctrl", "Method", "action", "Method", "route", "DELETE /api/v1/method/delete")
	service.Mux.Handle("PUT", "/api/v1/method/put", ctrl.MuxHandler("Method", h, nil))
	service.LogInfo("mount", "ctrl", "Method", "action", "Method", "route", "PUT /api/v1/method/put")
}

// handleMethodOrigin applies the CORS response headers corresponding to the origin.
func handleMethodOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SecurityController is the controller interface for the Security actions.
type SecurityController interface {
	goa.Muxer
	Security(*SecuritySecurityContext) error
}

// MountSecurityController "mounts" a Security resource controller on the given service.
func MountSecurityController(service *goa.Service, ctrl SecurityController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v1/securiy", ctrl.MuxHandler("preflight", handleSecurityOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSecuritySecurityContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Security(rctx)
	}
	h = handleSecurityOrigin(h)
	service.Mux.Handle("GET", "/api/v1/securiy", ctrl.MuxHandler("Security", h, nil))
	service.LogInfo("mount", "ctrl", "Security", "action", "Security", "route", "GET /api/v1/securiy")
}

// handleSecurityOrigin applies the CORS response headers corresponding to the origin.
func handleSecurityOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger/*filepath", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger/*filepath", "public/swagger/")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swagger/", "route", "GET /swagger/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger/", "public/swagger/index.html")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swagger/index.html", "route", "GET /swagger/")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// ValidationController is the controller interface for the Validation actions.
type ValidationController interface {
	goa.Muxer
	Validation(*ValidationValidationContext) error
}

// MountValidationController "mounts" a Validation resource controller on the given service.
func MountValidationController(service *goa.Service, ctrl ValidationController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v1/validation", ctrl.MuxHandler("preflight", handleValidationOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewValidationValidationContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Validation(rctx)
	}
	h = handleValidationOrigin(h)
	service.Mux.Handle("GET", "/api/v1/validation", ctrl.MuxHandler("Validation", h, nil))
	service.LogInfo("mount", "ctrl", "Validation", "action", "Validation", "route", "GET /api/v1/validation")
}

// handleValidationOrigin applies the CORS response headers corresponding to the origin.
func handleValidationOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// ViewController is the controller interface for the View actions.
type ViewController interface {
	goa.Muxer
	View(*ViewViewContext) error
}

// MountViewController "mounts" a View resource controller on the given service.
func MountViewController(service *goa.Service, ctrl ViewController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/v1/view/default", ctrl.MuxHandler("preflight", handleViewOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/v1/view/tiny", ctrl.MuxHandler("preflight", handleViewOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewViewViewContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.View(rctx)
	}
	h = handleViewOrigin(h)
	service.Mux.Handle("GET", "/api/v1/view/default", ctrl.MuxHandler("View", h, nil))
	service.LogInfo("mount", "ctrl", "View", "action", "View", "route", "GET /api/v1/view/default")
	service.Mux.Handle("GET", "/api/v1/view/tiny", ctrl.MuxHandler("View", h, nil))
	service.LogInfo("mount", "ctrl", "View", "action", "View", "route", "GET /api/v1/view/tiny")
}

// handleViewOrigin applies the CORS response headers corresponding to the origin.
func handleViewOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
