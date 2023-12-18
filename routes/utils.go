package routes

import (
	"beeapi/config"
	"beeapi/response"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func respond(context *gin.Context, response *response.Response) {
	if config.GetConfig().Development {
		for _, err := range response.Errors {
			if err.NativeError != nil {
				log.Info().Msgf("Error: %s", err.NativeError.Error())
				err.Message = err.Message + " Note: This error is native and development mode is enabled, check the logs for more information: " + err.NativeError.Error()
			}
		}
	}

	context.JSON(response.HTTPCode, response)
}
