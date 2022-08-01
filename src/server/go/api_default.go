/*
 * Basin RPC API
 *
 * Basin RPC API
 *
 * API version: 1.0.11
 * Contact: sestinj@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"Read",
			strings.ToUpper("Get"),
			"/api/v3/read",
			c.Read,
		},
		{
			"Subscribe",
			strings.ToUpper("Post"),
			"/api/v3/subscribe",
			c.Subscribe,
		},
		{
			"Write",
			strings.ToUpper("Put"),
			"/api/v3/write",
			c.Write,
		},
	}
}

// Read - Read Basin resource
func (c *DefaultApiController) Read(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	urlParam := query.Get("url")
	result, err := c.service.Read(r.Context(), urlParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Subscribe - Request subscription to Basin resource
func (c *DefaultApiController) Subscribe(w http.ResponseWriter, r *http.Request) {
	subscribeRequestParam := SubscribeRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscribeRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscribeRequestRequired(subscribeRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Subscribe(r.Context(), subscribeRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Write - Write Basin resource
func (c *DefaultApiController) Write(w http.ResponseWriter, r *http.Request) {
	writeRequestParam := WriteRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&writeRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertWriteRequestRequired(writeRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Write(r.Context(), writeRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
