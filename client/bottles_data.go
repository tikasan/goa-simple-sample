// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa simple sample": bottles_data Resource Client
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
	"strconv"
)

// AddBottlesDataPath computes a request path to the add action of bottles_data.
func AddBottlesDataPath() string {

	return fmt.Sprintf("/api/v1/bottles_data/")
}

// 追加
func (c *Client) AddBottlesData(ctx context.Context, path string, accountID int, name string, quantity int) (*http.Response, error) {
	req, err := c.NewAddBottlesDataRequest(ctx, path, accountID, name, quantity)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddBottlesDataRequest create the request corresponding to the add action endpoint of the bottles_data resource.
func (c *Client) NewAddBottlesDataRequest(ctx context.Context, path string, accountID int, name string, quantity int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	tmp36 := strconv.Itoa(accountID)
	values.Set("account_id", tmp36)
	values.Set("name", name)
	tmp37 := strconv.Itoa(quantity)
	values.Set("quantity", tmp37)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteBottlesDataPath computes a request path to the delete action of bottles_data.
func DeleteBottlesDataPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/bottles_data/%s", param0)
}

// 削除
func (c *Client) DeleteBottlesData(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteBottlesDataRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteBottlesDataRequest create the request corresponding to the delete action endpoint of the bottles_data resource.
func (c *Client) NewDeleteBottlesDataRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListBottlesDataPath computes a request path to the list action of bottles_data.
func ListBottlesDataPath() string {

	return fmt.Sprintf("/api/v1/bottles_data/")
}

// 複数
func (c *Client) ListBottlesData(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBottlesDataRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBottlesDataRequest create the request corresponding to the list action endpoint of the bottles_data resource.
func (c *Client) NewListBottlesDataRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowBottlesDataPath computes a request path to the show action of bottles_data.
func ShowBottlesDataPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/bottles_data/%s", param0)
}

// 単数
func (c *Client) ShowBottlesData(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBottlesDataRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBottlesDataRequest create the request corresponding to the show action endpoint of the bottles_data resource.
func (c *Client) NewShowBottlesDataRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateBottlesDataPath computes a request path to the update action of bottles_data.
func UpdateBottlesDataPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/bottles_data/%s", param0)
}

// 更新
func (c *Client) UpdateBottlesData(ctx context.Context, path string, name *string, quantity *int) (*http.Response, error) {
	req, err := c.NewUpdateBottlesDataRequest(ctx, path, name, quantity)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateBottlesDataRequest create the request corresponding to the update action endpoint of the bottles_data resource.
func (c *Client) NewUpdateBottlesDataRequest(ctx context.Context, path string, name *string, quantity *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if name != nil {
		values.Set("name", *name)
	}
	if quantity != nil {
		tmp38 := strconv.Itoa(*quantity)
		values.Set("quantity", tmp38)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
