package applemusic

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Apple Music client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a applemusic.Client that is configured to talk to that test server.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Apple Music client configured to use test server
	client = NewClient(nil)
	u, _ := url.Parse(server.URL)
	client.BaseURL = u
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Set(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func Test_makeIdsOptions(t *testing.T) {
	testCases := []struct {
		ids  []string
		opt  *Options
		want IdsOptions
	}{
		{
			ids:  []string{},
			opt:  nil,
			want: IdsOptions{},
		},
		{
			ids: []string{"1", "2", "3"},
			opt: nil,
			want: IdsOptions{
				Ids: "1,2,3",
			},
		},
		{
			ids: []string{"1", "2", "3"},
			opt: &Options{
				Language: "en-us",
			},
			want: IdsOptions{
				Ids: "1,2,3",
				Options: Options{
					Language: "en-us",
				},
			},
		},
	}
	for k, tc := range testCases {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			got := makeIdsOptions(tc.ids, tc.opt)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("makeIdsOptions() is %v, want %v", got, tc.want)
			}
		})
	}
}

// test we don't override the params already given when another param is specified in the url
func Test_addOptions_noParamsOverride(t *testing.T) {
	actualParams := "?offset=100"

	want := "/v1/me/library/playlists/p.2P6WgVAuVeYx3OB/tracks?l=fr&offset=100"
	got, _ := addOptions("/v1/me/library/playlists/p.2P6WgVAuVeYx3OB/tracks"+actualParams, Options{Language: "fr"})

	if got != want {
		t.Errorf("Url is %s, want %s", got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, defaultUserAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	type foo struct {
		Name string `json:"name"`
		Type string `json:"type,omitempty"`
	}

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &foo{Name: "Tester"}, `{"name":"Tester"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	// test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestNewRequest_invalidJSON(t *testing.T) {
	c := NewClient(nil)

	type T struct {
		A map[interface{}]interface{}
	}
	_, err := c.NewRequest("GET", "/", &T{})

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Errorf("Expected a JSON error; got %#v.", err)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

func TestNewRequest_emptyUserAgent(t *testing.T) {
	c := NewClient(nil)
	c.UserAgent = ""
	req, err := c.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if _, ok := req.Header["User-Agent"]; ok {
		t.Fatal("constructed request contains unexpected User-Agent header")
	}
}

func TestNewRequest_emptyBody(t *testing.T) {
	c := NewClient(nil)
	req, err := c.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Fatalf("constructed request contains a non-nil Body")
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	client.Do(context.Background(), req, body)

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(context.Background(), req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

func TestDo_noContent(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(context.Background(), req, &body)
	if err != nil {
		t.Fatalf("Do returned unexpected error: %v", err)
	}
}

func TestCheckResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(strings.NewReader(`{
  "errors": [
    {
      "id": "ID",
      "title": "Invalid Parameter Value",
      "detail": "Invalid language tag 'zh-tw'",
      "status": "400",
      "code": "40005",
      "source": {
        "parameter": "l"
      }
    }
  ]
}`)),
	}

	err := CheckResponse(res).(*ErrorResponse)
	if err == nil {
		t.Errorf("Expected error response.")
	}

	want := &ErrorResponse{
		Response: res,
		Errors: []Error{
			{
				Id:     "ID",
				Title:  "Invalid Parameter Value",
				Detail: "Invalid language tag 'zh-tw'",
				Status: "400",
				Code:   "40005",
				Source: Source{
					Parameter: "l",
				},
			},
		},
	}
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Error = %#v, want %#v", err, want)
	}
}

// ensure that we properly handle API errors that do not contain a response body.
func TestCheckResponse_noBody(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}

	err := CheckResponse(res).(*ErrorResponse)
	if err == nil {
		t.Errorf("Expected error response.")
	}

	want := &ErrorResponse{
		Response: res,
	}
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Error = %#v, want %#v", err, want)
	}
}

func TestCheckResponse_statusUnauthorized(t *testing.T) {
	u, _ := url.Parse("/")
	res := &http.Response{
		Request: &http.Request{
			Method: "GET",
			URL:    u,
		},
		StatusCode: http.StatusUnauthorized,
		Body:       ioutil.NopCloser(strings.NewReader("Unauthorized")),
	}

	err := CheckResponse(res)
	if err == nil {
		t.Error("Expected error response.")
	}
	if got, want := err.(*UnauthorizedError).Message, "Unauthorized"; got != want {
		t.Errorf("Error = %v, want %v", got, want)
	}
	if got, want := err.Error(), "GET /: 401 Unauthorized"; got != want {
		t.Errorf("Error = %v, want %v", got, want)
	}
}

func TestCheckResponse_statusTooManyRequests(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusTooManyRequests,
		Body:       ioutil.NopCloser(strings.NewReader(`{"message":"API capacity exceeded"}`)),
	}

	err := CheckResponse(res)
	if err == nil {
		t.Error("Expected error response.")
	}
	if got, want := err.(*TooManyRequestsError).Message, "API capacity exceeded"; got != want {
		t.Errorf("Error = %v, want %v", got, want)
	}
}

func TestTransport(t *testing.T) {
	setup()
	defer teardown()

	token := "TOKEN"
	musicUserToken := "MUSIC_USER_TOKEN"

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", token); got != want {
			t.Errorf("request contained token %s, want %s", got, want)
		}
		if got, want := r.Header.Get("Music-User-Token"), musicUserToken; got != want {
			t.Errorf("request contained music user token %s, want %s", got, want)
		}
	})

	tp := &Transport{
		Token:          token,
		MusicUserToken: musicUserToken,
	}
	c := NewClient(tp.Client())
	c.BaseURL = client.BaseURL
	req, _ := c.NewRequest("GET", "/", nil)
	c.Do(context.Background(), req, nil)
}

func TestTransport_transport(t *testing.T) {
	// default transport
	tp := &Transport{}
	if tp.transport() != http.DefaultTransport {
		t.Errorf("Expected http.DefaultTransport to be used.")
	}

	// custom transport
	tp = &Transport{
		Transport: &http.Transport{},
	}
	if tp.transport() == http.DefaultTransport {
		t.Errorf("Expected custom transport to be used.")
	}
}
