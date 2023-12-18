package routes

import (
	"beeapi/response"
	"github.com/gin-gonic/gin"
)

type Route interface {
	// GetPath returns the path for the route.
	GetPath() string

	// GetMethod returns the HTTP method for the route.
	GetMethod() string

	// Prehandle is called before Handle and is used to check if the request is valid.
	// If the request is invalid, return an error.
	// TODO: create a system which allows for simple conditional checking of the request.
	Prehandle(context *gin.Context) *response.Response

	// Handle is called to handle the request.
	// Only called if Prehandle returns no errors.
	Handle(context *gin.Context) *response.Response
}
