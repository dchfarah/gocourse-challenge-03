package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []routeEntry
	WebServerPort string
}

type routeEntry struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]routeEntry, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, routeEntry{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, route := range s.Handlers {
		s.Router.Method(route.Method, route.Path, route.Handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
