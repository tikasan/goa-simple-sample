// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "goa simple sample": bottles Resource Client
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

// AddBottlesPath computes a request path to the add action of bottles.
func AddBottlesPath() string {

	return fmt.Sprintf("/api/v1/bottles")
}

// 追加
func (c *Client) AddBottles(ctx context.Context, path string, accountID int, name string, quantity int) (*http.Response, error) {
	req, err := c.NewAddBottlesRequest(ctx, path, accountID, name, quantity)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddBottlesRequest create the request corresponding to the add action endpoint of the bottles resource.
func (c *Client) NewAddBottlesRequest(ctx context.Context, path string, accountID int, name string, quantity int) (*http.Request, error) {
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

// DeleteBottlesPath computes a request path to the delete action of bottles.
func DeleteBottlesPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/bottles/%s", param0)
}

// 削除
func (c *Client) DeleteBottles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteBottlesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteBottlesRequest create the request corresponding to the delete action endpoint of the bottles resource.
func (c *Client) NewDeleteBottlesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBottlesPath computes a request path to the list action of bottles.
func ListBottlesPath() string {

	return fmt.Sprintf("/api/v1/bottles")
}

// 複数
func (c *Client) ListBottles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBottlesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBottlesRequest create the request corresponding to the list action endpoint of the bottles resource.
func (c *Client) NewListBottlesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListRelationBottlesPath computes a request path to the listRelation action of bottles.
func ListRelationBottlesPath() string {

	return fmt.Sprintf("/api/v1/bottles/relation")
}

// 複数(リレーション版)
func (c *Client) ListRelationBottles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListRelationBottlesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListRelationBottlesRequest create the request corresponding to the listRelation action endpoint of the bottles resource.
func (c *Client) NewListRelationBottlesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowBottlesPath computes a request path to the show action of bottles.
func ShowBottlesPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/bottles/%s", param0)
}

// 単数
func (c *Client) ShowBottles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBottlesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBottlesRequest create the request corresponding to the show action endpoint of the bottles resource.
func (c *Client) NewShowBottlesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateBottlesPath computes a request path to the update action of bottles.
func UpdateBottlesPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/bottles/%s", param0)
}

// 更新
func (c *Client) UpdateBottles(ctx context.Context, path string, name *string, quantity *int) (*http.Response, error) {
	req, err := c.NewUpdateBottlesRequest(ctx, path, name, quantity)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateBottlesRequest create the request corresponding to the update action endpoint of the bottles resource.
func (c *Client) NewUpdateBottlesRequest(ctx context.Context, path string, name *string, quantity *int) (*http.Request, error) {
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
