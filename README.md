# Go module `httputils`

Module `github.com/maffeis/httputils` provides utility functions for developing HTTP server applications in go.
The following example sets up a go server which listens on port 8080 for HTTP requests and on 8443 for HTTPS requests. The HTTPS socket is secured with a certificate and with a private key specified to the `ListenHTTPS` function as file paths. A REST endpoint `GET /m/{msg}` is exposed, which returns a JSON response document.

```go
import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maffeis/httputils"
	log "github.com/sirupsen/logrus"
)

func main() {
	// channel for synchronizing with the HTTP/HTTPS event loop:
	errorchan := make(chan error)

	log.Infof("opening port 8080/8443")

	httputils.AddRequestHandler("/m/{msg}", "GET", handleMessage)

	httputils.ListenHTTP("127.0.0.1:8080", 30, 30, errorchan)
	httputils.ListenHTTPS("127.0.0.1:8443", 30, 30, "ssl/gosrv.cert", "ssl/gosrv.key", errorchan)

	// wait for the HTTP/HTTPS event loop to terminate:
	err := <-errorchan
	if err != nil {
		log.Fatalf("ListenHTTP/S failed: %s", err.Error())
	} else {
		log.Infof("HTTP/S terminated")
	}

	// wait for the HTTP/HTTPS event loop to terminate:
	err = <-errorchan
	if err != nil {
		log.Fatalf("ListenHTTP/S failed: %s", err.Error())
	} else {
		log.Infof("HTTP/S terminated")
    }
    
    log.Infof("server terminated")
}

// handler function for "/m/{msg}" REST endpoint:
func handleMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]
	data1 := map[string]string{}
	data1["message"] = message
	response := map[string]map[string]string{}
	response["response"] = data1

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
```