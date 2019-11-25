package stapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RequestData struct {
	request *http.Request
	body    string
}

var requestChan = make(chan *RequestData, 1)
var responseBody = ""
var httpServer *httptest.Server

func setup() {
	httpServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		requestChan <- &RequestData{r, string(data)}
		fmt.Fprintf(w, responseBody)
	}))
}

func teardown() {
	if httpServer != nil {
		httpServer.Close()
	}
}

func TestClientRequest(t *testing.T) {

	cl := New("key", nil)

	req, err := cl.newRequest(http.MethodGet, BuildURL("example"), struct {
		Example string `url:"example"`
	}{"test"})

	if err != nil {
		t.Error(err)
		return
	}

	if req.Method != http.MethodGet {
		t.Errorf("wrong http method, expected %s, got %s\n", http.MethodGet, req.Method)
	}

	if req.URL.Query().Get("example") != "test" {
		t.Errorf("invalid query param, expected %s, got %s", "test", req.URL.Query().Get("example"))
	}

	if req.URL.String() != "http://stapi.co/api/v1/rest/example?apiKey=key&example=test" {
		t.Errorf("invalid URL, expected %s, got %s\n",
			"http://stapi.co/api/v1/rest/example?apiKey=key&example=test",
			req.URL.String(),
		)
	}

	if req.UserAgent() != UserAgent {
		t.Errorf("invalid user agent, expected %s, got %s", req.UserAgent(), UserAgent)
	}

	req, err = cl.newRequest(http.MethodPost, BuildURL("example"), struct {
		Example string `url:"example"`
	}{"test"})
	if err != nil {
		t.Error(err)
		return
	}

	// POST method should send options as application/x-www-form-urlencoded
	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Errorf("invalid header for http post, expected %s, got %s",
			"application/x-www-form-urlencoded",
			req.Header.Get("Content-Type"),
		)
	}
}

func TestClient(t *testing.T) {
	setup()
	defer teardown()

	cl := New("key", httpServer.Client())

	var resData struct {
		OK bool `json:"ok"`
	}
	responseBody = `{"ok": true}`
	err := cl.Get(httpServer.URL, nil, &resData)
	<-requestChan
	if err != nil {
		t.Error(err)
	}
	if !resData.OK {
		t.Error("invalid struct parsing in GET")
	}

	err = cl.Post(httpServer.URL, nil, nil)
	<-requestChan
	if err != nil {
		t.Error(err)
	}
}
