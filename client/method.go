// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "goa simple sample": method Resource Client
//
// Command:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// EtcMethodPath computes a request path to the etc action of method.
func EtcMethodPath(id int, type_ int) string {
	param0 := strconv.Itoa(id)
	param1 := strconv.Itoa(type_)

	return fmt.Sprintf("/api/v1/method/users/%s/follow/%s", param0, param1)
}

// ちょっと特殊ケース
func (c *Client) EtcMethod(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewEtcMethodRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewEtcMethodRequest create the request corresponding to the etc action endpoint of the method resource.
func (c *Client) NewEtcMethodRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FollowMethodPath computes a request path to the follow action of method.
func FollowMethodPath() string {

	return fmt.Sprintf("/api/v1/method/users/follow")
}

// FollowMethodPath2 computes a request path to the follow action of method.
func FollowMethodPath2() string {

	return fmt.Sprintf("/api/v1/method/users/follow")
}

// フォロー操作
func (c *Client) FollowMethod(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFollowMethodRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFollowMethodRequest create the request corresponding to the follow action endpoint of the method resource.
func (c *Client) NewFollowMethodRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListMethodPath computes a request path to the list action of method.
func ListMethodPath() string {

	return fmt.Sprintf("/api/v1/method/list")
}

// ListMethodPath2 computes a request path to the list action of method.
func ListMethodPath2() string {

	return fmt.Sprintf("/api/v1/method/list/new")
}

// ListMethodPath3 computes a request path to the list action of method.
func ListMethodPath3() string {

	return fmt.Sprintf("/api/v1/method/list/topic")
}

// リストを返す
func (c *Client) ListMethod(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListMethodRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListMethodRequest create the request corresponding to the list action endpoint of the method resource.
func (c *Client) NewListMethodRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// MethodMethodPath computes a request path to the method action of method.
func MethodMethodPath() string {

	return fmt.Sprintf("/api/v1/method/get")
}

// MethodMethodPath2 computes a request path to the method action of method.
func MethodMethodPath2() string {

	return fmt.Sprintf("/api/v1/method/post")
}

// MethodMethodPath3 computes a request path to the method action of method.
func MethodMethodPath3() string {

	return fmt.Sprintf("/api/v1/method/delete")
}

// MethodMethodPath4 computes a request path to the method action of method.
func MethodMethodPath4() string {

	return fmt.Sprintf("/api/v1/method/put")
}

// HTTPメソッド
func (c *Client) MethodMethod(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewMethodMethodRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMethodMethodRequest create the request corresponding to the method action endpoint of the method resource.
func (c *Client) NewMethodMethodRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
