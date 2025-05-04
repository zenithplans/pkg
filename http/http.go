package http

import (
	"context"
	"fmt"
	"net/http"
)

type HttpServer struct {
	driver *http.Server
}

// Start listens on the TCP network address s.Addr and then
// calls [Serve] to handle requests on incoming connections.
func (s *HttpServer) Start() error {
	return s.driver.ListenAndServe()
}

func (s *HttpServer) GetMux() *http.ServeMux {
	// NOTE: [http.ServeMux] is the only supported mux
	// at the moment. It is set as soon as [HttpServer] is created
	// with [New]. Hence, no assertion error handling.
	return s.driver.Handler.(*http.ServeMux)
}

// Close gracefully shuts down the server without interrupting any
// active connections. Shutdown works by first closing all open
// listeners, then closing all idle connections, and then waiting
// indefinitely for connections to return to idle and then shut down.
// If the provided context expires before the shutdown is complete,
// Shutdown returns the context's error, otherwise it returns any
// error returned from closing the [Server]'s underlying Listener(s).
func (s *HttpServer) Close(ctx context.Context) error {
	return s.driver.Shutdown(ctx)
}

func New(conf Config) *HttpServer {
	return &HttpServer{
		driver: &http.Server{
			Addr: fmt.Sprintf(
				"%s:%d",
				conf.Host,
				conf.Port,
			),
			Handler:           http.NewServeMux(),
			ReadTimeout:       conf.ReadTimeout,
			WriteTimeout:      conf.WriteTimeout,
			ReadHeaderTimeout: conf.ReadHeaderTimeout,
		},
	}
}
