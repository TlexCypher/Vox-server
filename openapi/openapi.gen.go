// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"

	"github.com/labstack/echo/v4"
)

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	// Role role for created user.
	Role *string `json:"role,omitempty"`

	// Username name for created user.
	Username *string `json:"username,omitempty"`
}

// CreateUserResponse defines model for CreateUserResponse.
type CreateUserResponse struct {
	// Code status code
	Code *int                `json:"code,omitempty"`
	Data *valueobject.UserId `json:"data,omitempty"`
}

// Error defines model for Error.
type Error struct {
	// Code status code.
	Code *int `json:"code,omitempty"`

	// Message error message.
	Message *string `json:"message,omitempty"`
}

// GetAllUsersResponse defines model for GetAllUsersResponse.
type GetAllUsersResponse struct {
	// Code status code.
	Code *int `json:"code,omitempty"`

	// Data list of all users.
	Data *[]entity.User `json:"data,omitempty"`

	// Message additional message for response.
	Message *string `json:"message,omitempty"`
}

// CreateNewUserJSONRequestBody defines body for CreateNewUser for application/json ContentType.
type CreateNewUserJSONRequestBody = CreateUserRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Health check
	// (GET /healthcheck)
	HealthCheck(ctx echo.Context) error
	// Get all users
	// (GET /users)
	GetAllUsers(ctx echo.Context) error
	// Create a new user
	// (POST /users)
	CreateNewUser(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// HealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.HealthCheck(ctx)
	return err
}

// GetAllUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAllUsers(ctx)
	return err
}

// CreateNewUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateNewUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateNewUser(ctx)
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

	router.GET(baseURL+"/healthcheck", wrapper.HealthCheck)
	router.GET(baseURL+"/users", wrapper.GetAllUsers)
	router.POST(baseURL+"/users", wrapper.CreateNewUser)

}