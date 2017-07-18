package applemusic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion   = "0.0.1"
	defaultBaseURL   = "https://api.music.apple.com/"
	defaultUserAgent = "go-applemusic/" + libraryVersion
)

// A Client manages communication with the Apple Music API.
type Client struct {
	client *http.Client

	BaseURL   *url.URL
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.
}

type service struct {
	client *Client
}

type Options struct {
	// The localization to use, specified by a language tag.
	// Any supported language tag may be used here, if one is not specified then en-us is used.
	Language string `url:"l,omitempty"`

	// Additional relationships to include in the fetch.
	Include string `url:"include,omitempty"`
}

// PageOptions specifies the optional parameters to support pagination of the objects.
type PageOptions struct {
	// The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Limit int `url:"limit,omitempty"`

	// The next page or group of objects to fetch.
	Offset int `url:"offset,omitempty"`

	Options
}

// addOptions adds the parameters in opt as URL query parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns a new Apple Music API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		UserAgent: defaultUserAgent,
		BaseURL:   baseURL,
	}
	c.common.client = c

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a Apple Music API response.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// Do sends an API request and returns the API response.
//
// The provided ctx must be non-nil. If it is canceled or time out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, nil
}

// Source represents the source of an error.
type Source struct {
	Parameter string      `json:"parameter"`
	Pointer   interface{} `json:"pointer"` // JSON pointer, A pointer to the associated entry in the request document.
}

// Error contains information about an error that occurred while processing a request.
type Error struct {
	Id     string      `json:"id"`
	About  string      `json:"about"`
	Status string      `json:"status"`
	Code   string      `json:"code"`
	Title  string      `json:"title"`
	Detail string      `json:"detail"`
	Source Source      `json:"source"`
	Meta   interface{} `json:"meta"`
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response
	Errors   []Error `json:"errors"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.Errors)
}

// CheckResponse checks the API response for errors.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	if r.StatusCode == http.StatusUnauthorized {
		return errors.New(http.StatusText(http.StatusUnauthorized))
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	switch r.StatusCode {
	default:
		return errorResponse
	}
}
