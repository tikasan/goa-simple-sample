// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": validation TestHelpers
//
// The content of this file is auto-generated, DO NOT MODIFY

package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/tikasan/goa-simple-sample/app"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// ValidationValidationBadRequest runs the method Validation of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ValidationValidationBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ValidationController, id *int, default_ string, email *string, enum *string, integer *int, string_ *string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if id != nil {
		sliceVal := []string{strconv.Itoa(*id)}
		query["ID"] = sliceVal
	}
	{
		sliceVal := []string{default_}
		query["default"] = sliceVal
	}
	if email != nil {
		sliceVal := []string{*email}
		query["email"] = sliceVal
	}
	if enum != nil {
		sliceVal := []string{*enum}
		query["enum"] = sliceVal
	}
	if integer != nil {
		sliceVal := []string{strconv.Itoa(*integer)}
		query["integer"] = sliceVal
	}
	if string_ != nil {
		sliceVal := []string{*string_}
		query["string"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v1/validation"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if id != nil {
		sliceVal := []string{strconv.Itoa(*id)}
		prms["ID"] = sliceVal
	}
	{
		sliceVal := []string{default_}
		prms["default"] = sliceVal
	}
	if email != nil {
		sliceVal := []string{*email}
		prms["email"] = sliceVal
	}
	if enum != nil {
		sliceVal := []string{*enum}
		prms["enum"] = sliceVal
	}
	if integer != nil {
		sliceVal := []string{strconv.Itoa(*integer)}
		prms["integer"] = sliceVal
	}
	if string_ != nil {
		sliceVal := []string{*string_}
		prms["string"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ValidationTest"), rw, req, prms)
	validationCtx, _err := app.NewValidationValidationContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Validation(validationCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// ValidationValidationOK runs the method Validation of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ValidationValidationOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ValidationController, id *int, default_ string, email *string, enum *string, integer *int, string_ *string) (http.ResponseWriter, *app.Validationtype) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if id != nil {
		sliceVal := []string{strconv.Itoa(*id)}
		query["ID"] = sliceVal
	}
	{
		sliceVal := []string{default_}
		query["default"] = sliceVal
	}
	if email != nil {
		sliceVal := []string{*email}
		query["email"] = sliceVal
	}
	if enum != nil {
		sliceVal := []string{*enum}
		query["enum"] = sliceVal
	}
	if integer != nil {
		sliceVal := []string{strconv.Itoa(*integer)}
		query["integer"] = sliceVal
	}
	if string_ != nil {
		sliceVal := []string{*string_}
		query["string"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v1/validation"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if id != nil {
		sliceVal := []string{strconv.Itoa(*id)}
		prms["ID"] = sliceVal
	}
	{
		sliceVal := []string{default_}
		prms["default"] = sliceVal
	}
	if email != nil {
		sliceVal := []string{*email}
		prms["email"] = sliceVal
	}
	if enum != nil {
		sliceVal := []string{*enum}
		prms["enum"] = sliceVal
	}
	if integer != nil {
		sliceVal := []string{strconv.Itoa(*integer)}
		prms["integer"] = sliceVal
	}
	if string_ != nil {
		sliceVal := []string{*string_}
		prms["string"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ValidationTest"), rw, req, prms)
	validationCtx, _err := app.NewValidationValidationContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Validation(validationCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Validationtype
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Validationtype)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Validationtype", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
