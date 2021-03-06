// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, server, port string) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()
	matrix := New(server, port)

	mux.Handle(http.MethodGet, "/test", readiness)
	mux.Handle(http.MethodGet, "/.well-known/matrix/server", matrix.delegate)

	return mux
}
