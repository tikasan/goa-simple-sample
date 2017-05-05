// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": validation Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// ValidationValidationPath computes a request path to the validation action of validation.
func ValidationValidationPath() string {

	return fmt.Sprintf("/api/v1/validation")
}

// Validation
func (c *Client) ValidationValidation(ctx context.Context, path string, id int, defaultType string, email string, enumType string, integerType int, reg string, stringType string) (*http.Response, error) {
	req, err := c.NewValidationValidationRequest(ctx, path, id, defaultType, email, enumType, integerType, reg, stringType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewValidationValidationRequest create the request corresponding to the validation action endpoint of the validation resource.
func (c *Client) NewValidationValidationRequest(ctx context.Context, path string, id int, defaultType string, email string, enumType string, integerType int, reg string, stringType string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	tmp9 := strconv.Itoa(id)
	values.Set("ID", tmp9)
	values.Set("defaultType", defaultType)
	values.Set("email", email)
	values.Set("enumType", enumType)
	tmp10 := strconv.Itoa(integerType)
	values.Set("integerType", tmp10)
	values.Set("reg", reg)
	values.Set("stringType", stringType)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
