package httputils

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
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

type CoreV1TestObject struct {
	// add a Mock object instance
	mock.Mock

	// other fields go here as normal
}

// CoreV1 retrieves the CoreV1Client
func (o *CoreV1TestObject) CoreV1() corev1.CoreV1Interface {
	args := o.Called()
	return args.Get(0).(corev1.CoreV1Interface)
}

type CoreV1InterfaceTestObject struct {
	// add a Mock object instance
	mock.Mock

	// other fields go here as normal
}

// CoreV1 retrieves the CoreV1Client
func (o *CoreV1InterfaceTestObject) Secrets(nameSpace string) corev1.SecretInterface {
	args := o.Called(nameSpace)
	return args.Get(0).(corev1.SecretInterface)
}

// Here's our mock which just contains some variables that will be filled for running assertions on them later on
type mockedHTTPUtils struct {
	err      error
	messages []string
}

func TestLoadSslCert(t *testing.T) {
	// create an instance of our test object
	testObj := new(CoreV1TestObject)

	//clientSet := new(CoreV1TestObject)

	//httpUtils := HTTPUtils{clientSet}

	//_ = httpUtils

	testObj.On("CoreV1", mock.Anything).Return(new(CoreV1InterfaceTestObject))

	//LoadSslCert(&clientSet, "", "", "")
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
}

// URI interface
type URI interface {
	GetURL() string
}

// MessageSender interface
type MessageSender interface {
	SendMessage(message string) error
}

// This one is the "object" that our users will call to use this package functionalities
type API struct {
	baseURL  string
	endpoint string
}

// Here we make API implement implicitly the URI interface
func (api *API) GetURL() string {
	return api.baseURL + api.endpoint
}

// Here we make API implement implicitly the MessageSender interface
// Again we're just WRAPPING the sendMessage function here, nothing fancy
func (api *API) SendMessage(message string) error {
	return sendMessage(api, message)
}

// We want to test this method but it calls SendMessage which makes a real HTTP request!
// Again we're just WRAPPING the sendDataSynchronously function here, nothing fancy
func (api *API) SendDataSynchronously(data []string) error {
	return sendDataSynchronously(api, data)
}

// this would make a real HTTP request
func sendMessage(uri URI, message string) error {
	fmt.Println("This function won't get called because we will mock it")
	return nil
}

// this is the function we want to test :)
func sendDataSynchronously(sender MessageSender, data []string) error {
	for _, text := range data {
		err := sender.SendMessage(text)

		if err != nil {
			return err
		}
	}

	return nil
}

// TEST CASE BELOW

// Here's our mock which just contains some variables that will be filled for running assertions on them later on
type mockedSender struct {
	err      error
	messages []string
}

// We make our mock implement the MessageSender interface so we can test sendDataSynchronously
func (sender *mockedSender) SendMessage(message string) error {
	// let's store all received messages for later assertions
	sender.messages = append(sender.messages, message)

	return sender.err // return error for later assertions
}

func TestSendsAllMessagesSynchronously(t *testing.T) {
	mockedMessages := make([]string, 0)
	sender := mockedSender{nil, mockedMessages}

	messagesToSend := []string{"one", "two", "three"}
	err := sendDataSynchronously(&sender, messagesToSend)

	if err == nil {
		fmt.Println("All good here we expect the error to be nil:", err)
	}

	expectedMessages := fmt.Sprintf("%v", messagesToSend)
	actualMessages := fmt.Sprintf("%v", sender.messages)

	if expectedMessages == actualMessages {
		fmt.Println("Actual messages are as expected:", actualMessages)
	}
}
