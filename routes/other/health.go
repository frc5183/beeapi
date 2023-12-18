package routes

import (
	"beeapi/response"
	"github.com/gin-gonic/gin"
)

type HealthRoute struct{}

func (route HealthRoute) GetPath() string {
	return "/health"
}

func (route HealthRoute) GetMethod() string {
	return "GET"
}

func (route HealthRoute) Prehandle(context *gin.Context) *response.Response {
	return nil
}

func (route HealthRoute) Handle(context *gin.Context) *response.Response {
	return response.CreateSuccessResponse("OK", nil, 200)
}
