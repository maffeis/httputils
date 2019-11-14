package httputils

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

var router = mux.NewRouter()

var srvHTTP http.Server = http.Server{}

var srvHTTPS http.Server = http.Server{}

var corsAllowedHeaders []string = []string{"Authorization"}

var corsAllowedOrigins []string = []string{"*"}

var corsAllowedMethods []string = []string{"GET", "POST", "OPTIONS"}

// AddRequestHandler registers a HTTP request handler under a given URL path and HTTP method.
func AddRequestHandler(path string, method string, f func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(path, f).Methods(method)
}

// ListenHTTP binds a HTTP server socket and listen for HTTP requests. addr is of the form "ip:port", where ip denotes
// the IP address of a network interface, or 0.0.0.0 for binding the server socket to all interfaces.
func ListenHTTP(addr string, writeTimeoutSec int, readTimeoutSec int, errorchan chan error) {
	srvHTTP.Handler = corsHelper()
	srvHTTP.Addr = addr
	srvHTTP.WriteTimeout = time.Second * time.Duration(writeTimeoutSec)
	srvHTTP.ReadTimeout = time.Second * time.Duration(readTimeoutSec)

	go func() {
		error := srvHTTP.ListenAndServe()

		if error != nil {
			log.Print(error)
			errorchan <- error
		} else {
			errorchan <- nil
		}
	}()
}

// ListenHTTPS binds a HTTPS server socket and listen for HTTPS requests. addr is of the form "ip:port", where ip denotes
// the IP address of a network interface, or 0.0.0.0 for binding the server socket to all interfaces.
func ListenHTTPS(addr string, writeTimeoutSec int, readTimeoutSec int, sslCertFile string, sslKeyFile string, errorchan chan error) {
	srvHTTPS.Handler = corsHelper()
	srvHTTPS.Addr = addr
	srvHTTPS.WriteTimeout = time.Second * time.Duration(writeTimeoutSec)
	srvHTTPS.ReadTimeout = time.Second * time.Duration(readTimeoutSec)

	go func() {
		error := srvHTTPS.ListenAndServeTLS(sslCertFile, sslKeyFile)

		if error != nil {
			log.Print(error)
			errorchan <- error
		} else {
			errorchan <- nil
		}
	}()
}

// CloseHTTP ls the HTTP service
func CloseHTTP() {
	srvHTTP.Close()
}

// CloseHTTPS ls the HTTP service
func CloseHTTPS() {
	srvHTTPS.Close()
}

// CorsSetAllowedHeaders sets the CORS allowed headers
func CorsSetAllowedHeaders(val []string) {
	corsAllowedHeaders = val
}

// CorsGetAllowedHeaders gets the CORS allowed headers
func CorsGetAllowedHeaders() []string {
	return corsAllowedHeaders
}

// CorsSetAllowedOrigins sets the CORS allowed headers
func CorsSetAllowedOrigins(val []string) {
	corsAllowedOrigins = val
}

// CorsGetAllowedOrigins gets the CORS allowed headers
func CorsGetAllowedOrigins() []string {
	return corsAllowedOrigins
}

// CorsSetAllowedMethods sets the CORS allowed headers
func CorsSetAllowedMethods(val []string) {
	corsAllowedMethods = val
}

// CorsGetAllowedMethods gets the CORS allowed headers
func CorsGetAllowedMethods() []string {
	return corsAllowedMethods
}

func corsHelper() http.Handler {
	headersOk := handlers.AllowedHeaders(corsAllowedHeaders)
	originsOk := handlers.AllowedOrigins(corsAllowedOrigins)
	methodsOk := handlers.AllowedMethods(corsAllowedMethods)

	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}
