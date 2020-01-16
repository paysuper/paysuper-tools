package http

import (
	"bytes"
	"context"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultHttpClientTimeout = 10
	defaultResponseBodyLimit = 512
)

type Transport struct {
	Transport         http.RoundTripper
	Logger            *zap.SugaredLogger
	ResponseBodyLimit int
}

type contextKey struct {
	name string
}

func NewLoggedHttpClient(logger *zap.SugaredLogger) *http.Client {
	return &http.Client{
		Transport: &Transport{
			Logger:            logger,
			ResponseBodyLimit: defaultResponseBodyLimit,
		},
		Timeout: time.Duration(defaultHttpClientTimeout * time.Second),
	}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := context.WithValue(req.Context(), &contextKey{name: "RequestStart"}, time.Now())
	req = req.WithContext(ctx)

	var reqBody []byte

	if req.Body != nil {
		reqBody, _ = ioutil.ReadAll(req.Body)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

	resp, err := t.transport().RoundTrip(req)
	if err != nil {
		return resp, err
	}

	t.log(req.URL.Path, req.Header, reqBody, resp)

	return resp, err
}

func (t *Transport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}

func (t *Transport) log(reqUrl string, reqHeader http.Header, reqBody []byte, resp *http.Response) {
	if t.Logger == nil {
		return
	}

	var resBody []byte

	if resp.Body != nil {
		resBody, _ = ioutil.ReadAll(resp.Body)
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(resBody))

	data := []interface{}{
		"request_headers", t.httpHeadersToString(reqHeader),
		"request_body", string(reqBody),
		"response_status", resp.StatusCode,
		"response_headers", t.httpHeadersToString(resp.Header),
		"response_body", t.cutResponseBody(resBody),
	}

	t.Logger.Infow(reqUrl, data...)
}

func (t *Transport) cutResponseBody(body []byte) string {
	sBody := string(body)
	r := []rune(sBody)

	if len(r) >= t.ResponseBodyLimit {
		return string(r[:t.ResponseBodyLimit])
	}

	return sBody
}

func (t *Transport) httpHeadersToString(headers map[string][]string) string {
	var out string

	for k, v := range headers {
		out += k + ":" + v[0] + "\n "
	}

	return out
}
