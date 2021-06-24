package cmd

import (
	"net/http"
)

type transport struct {
	underlyingTransport http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-hasura-admin-secret", adminSecret)

	return t.underlyingTransport.RoundTrip(req)
}
