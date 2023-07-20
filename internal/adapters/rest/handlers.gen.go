// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package rest

import (
	"github.com/labstack/echo/v4"
)

// Error defines model for Error.
type Error struct {
	Detail *string `json:"detail,omitempty"`
	Status int     `json:"status"`
	Title  string  `json:"title"`
}

// ReportData defines model for ReportData.
type ReportData struct {
	Author    string                  `json:"author"`
	Body      string                  `json:"body"`
	Course    int                     `json:"course"`
	Group     string                  `json:"group"`
	LabNumber string                  `json:"labNumber"`
	Static    *map[string]interface{} `json:"static,omitempty"`
	Teacher   string                  `json:"teacher"`
	Title     string                  `json:"title"`
	WorkType  string                  `json:"workType"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse = Error

// GenerateReportJSONRequestBody defines body for GenerateReport for application/json ContentType.
type GenerateReportJSONRequestBody = ReportData

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Generate report in pdf format.
	// (POST /generateReport)
	GenerateReport(ctx echo.Context) error
	// Healthcheck
	// (GET /ping)
	Ping(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GenerateReport converts echo context to params.
func (w *ServerInterfaceWrapper) GenerateReport(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GenerateReport(ctx)
	return err
}

// Ping converts echo context to params.
func (w *ServerInterfaceWrapper) Ping(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Ping(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/generateReport", wrapper.GenerateReport)
	router.GET(baseURL+"/ping", wrapper.Ping)

}
