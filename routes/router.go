package routes

import (
	otherRoutes "beeapi/routes/other"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

var routes = []Route{
	otherRoutes.HealthRoute{},
}

func RegisterRoutes() {
	for _, route := range routes {
		router.Handle(route.GetMethod(), route.GetPath(), func(context *gin.Context) {
			response := route.Prehandle(context)
			if response != nil {
				respond(context, response)
				return
			}

			response = route.Handle(context)
			if response != nil {
				respond(context, response)
			}
		})
	}
}

func GetRouter() *gin.Engine {
	return router
}
