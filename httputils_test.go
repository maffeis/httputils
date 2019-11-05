package httputils

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAddRequestHandler(t *testing.T) {
	AddRequestHandler("/m/{msg}", "GET", handleMessage)
}

func TestListenHTTP(t *testing.T) {
	errorchan := make(chan error)

	ListenHTTP("1.2.3.4:9080", 0, 0, errorchan)

	err := <-errorchan

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestListenHTTPS(t *testing.T) {
	errorchan := make(chan error)

	ListenHTTPS("1.2.3.4:9080", 0, 0, "", "", errorchan)

	err := <-errorchan

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGetSet(t *testing.T) {
	if !reflect.DeepEqual(CorsGetAllowedHeaders(), []string{"Authorization"}) {
		t.Errorf("not equals")
	}

	CorsSetAllowedHeaders([]string{"X"})

	if !reflect.DeepEqual(CorsGetAllowedHeaders(), []string{"X"}) {
		t.Errorf("not equals")
	}

	if !reflect.DeepEqual(CorsGetAllowedOrigins(), []string{"*"}) {
		t.Errorf("not equals")
	}

	CorsSetAllowedOrigins([]string{"*"})

	if !reflect.DeepEqual(CorsGetAllowedOrigins(), []string{"*"}) {
		t.Errorf("not equals")
	}

	if !reflect.DeepEqual(CorsGetAllowedMethods(), []string{"GET", "POST", "OPTIONS"}) {
		t.Errorf("not equals")
	}

	CorsSetAllowedHeaders([]string{"X"})

	if !reflect.DeepEqual(CorsGetAllowedHeaders(), []string{"X"}) {
		t.Errorf("not equals: " + CorsGetAllowedHeaders()[0])
	}
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
}
