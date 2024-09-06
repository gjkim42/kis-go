package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	client    *http.Client
	url       string
	appKey    string
	appSecret string

	header map[string]string
}

type ClientOptions struct {
	Header map[string]string
}

func NewClient(httpclient *http.Client, url, appKey, appSecret string, opts ClientOptions) *Client {
	return &Client{
		client:    httpclient,
		url:       url,
		appKey:    appKey,
		appSecret: appSecret,
		header:    opts.Header,
	}
}

func (r *Client) Get() *request {
	req := newRequest(r)
	req.method = http.MethodGet
	return req
}

func (r *Client) Post() *request {
	req := newRequest(r)
	req.method = http.MethodPost
	return req
}

type request struct {
	client    *http.Client
	appKey    string
	appSecret string

	err error

	url      string
	urlQuery string

	method string
	body   []byte
	header map[string]string
}

func newRequest(c *Client) *request {
	return &request{
		client:    c.client,
		appKey:    c.appKey,
		appSecret: c.appSecret,
		header:    c.header,

		url: c.url,
	}
}

func (r *request) At(path string) *request {
	r.url += "/" + path
	return r
}

func (r *request) Header(key, value string) *request {
	if r.header == nil {
		r.header = make(map[string]string)
	}
	r.header[key] = value
	return r
}

func (r *request) Headers(headers map[string]string) *request {
	if r.header == nil {
		r.header = make(map[string]string)
	}
	for k, v := range headers {
		r = r.Header(k, v)
	}
	return r
}

func (r *request) Query(key, value string) *request {
	if r.urlQuery == "" {
		r.urlQuery = key + "=" + value
	} else {
		r.urlQuery += "&" + key + "=" + value
	}
	return r
}

func (r *request) Queries(queries map[string]string) *request {
	for k, v := range queries {
		r.Query(k, v)
	}
	return r
}

func (r *request) Body(obj any) *request {
	body, err := json.Marshal(obj)
	if err != nil {
		r.err = err
		return r
	}

	r.body = body

	return r
}

func (r *request) Do(ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, r.method, r.url+"?"+r.urlQuery, bytes.NewReader(r.body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("appKey", r.appKey)
	req.Header.Set("appSecret", r.appSecret)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")
	for k, v := range r.header {
		req.Header.Set(k, v)
	}

	return r.client.Do(req)
}
