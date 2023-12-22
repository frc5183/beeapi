package routes

import (
	"beeapi/config"
	"beeapi/response"
	itemRoutes "beeapi/routes/item"
	otherRoutes "beeapi/routes/other"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var router *gin.Engine

var routes = []Route{
	otherRoutes.HealthRoute{},

	itemRoutes.ItemGetRoute{},
	itemRoutes.ItemGetAllRoute{},
	itemRoutes.ItemSearchRoute{},
	itemRoutes.ItemCreateRoute{},
	itemRoutes.ItemUpdateRoute{},
	itemRoutes.ItemDeleteRoute{},
	itemRoutes.ItemBarcodeRoute{},
}

func Init() {
	if !config.GetConfig().Development {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.Default()
}

func RegisterRoutes() {
	for _, route := range routes {
		router.Handle(route.GetMethod(), route.GetPath(), func(context *gin.Context) {
			for _, route := range routes {
				if route.GetPath() == context.Request.URL.Path && route.GetMethod() == context.Request.Method {
					resp := route.Prehandle(context)
					if resp != nil {
						respond(context, resp)
						return
					}

					resp = route.Handle(context)
					if resp != nil {
						respond(context, resp)
					}
				}
			}
		})
	}
}

func respond(context *gin.Context, response *response.Response) {
	if config.GetConfig().Development {
		for _, err := range response.Errors {
			if err.NativeError != nil {
				log.Info().Msgf("Error: %s", err.NativeError.Error())
				err.Message = err.Message + " * This error carries a native error and development mode is enabled, check the logs for more information."
			}
		}
	}

	context.JSON(response.HTTPCode, response)
}

func GetRouter() *gin.Engine {
	return router
}
