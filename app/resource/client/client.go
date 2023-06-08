package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// inversion
type HTTPClientProvider interface {
	GetKey() string
	GetWithContext(ctx context.Context, param Parameter) (*http.Response, error)
}

type HTTPClient struct {
	HTTPClientProvider
	host string
	key  string
}

type Parameter struct {
	Path        string
	QueryParams map[string]string
	Body        interface{}
}

func (hc *HTTPClient) GetWithContext(ctx context.Context, param Parameter) (*http.Response, error) {

	var (
		fullUrl string
		body    io.Reader
		err     error
	)

	fullUrl = hc.generateUrl(param)
	body = http.NoBody
	if param.Body != nil {
		body, err = generateBody(param.Body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullUrl, body)
	if err != nil {
		return nil, err
	}

	// add queries
	addQueryString(req, param.QueryParams)

	// TODO: make more robust; for a MVP now accepted enc is in gzip only
	req.Header.Add("Accept-Encoding", "gzip")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (hc *HTTPClient) GetKey() string {
	return hc.key
}

func (hc *HTTPClient) generateUrl(param Parameter) string {
	fullPath := hc.host
	if len(param.Path) > 0 && !strings.HasPrefix(param.Path, "/") {
		fullPath = fullPath + "/"
	}
	fullPath = fullPath + param.Path

	return fullPath
}

func addQueryString(req *http.Request, queryStrings map[string]string) {
	q := req.URL.Query()
	for key, value := range queryStrings {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
}

func generateBody(bodyPlain interface{}) (io.Reader, error) {
	jsonBody, err := json.Marshal(bodyPlain)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonBody), nil
}
