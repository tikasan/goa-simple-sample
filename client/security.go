// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa simple sample": security Resource Client
//
// Command:
// $ goagen
// --design=github.com/pei0804/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/pei0804/goa-simple-sample
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// SecuritySecurityPath computes a request path to the security action of security.
func SecuritySecurityPath() string {

	return fmt.Sprintf("/api/v1/securiy/")
}

// セキュリティの例です
func (c *Client) SecuritySecurity(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSecuritySecurityRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSecuritySecurityRequest create the request corresponding to the security action endpoint of the security resource.
func (c *Client) NewSecuritySecurityRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.UserTokenSigner != nil {
		if err := c.UserTokenSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
