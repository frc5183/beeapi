package routes

import (
	"beeapi/database"
	"beeapi/models"
	"beeapi/response"
	"github.com/gin-gonic/gin"
)

type ItemCreateRoute struct{}

func (route ItemCreateRoute) GetPath() string {
	return "/item"
}

func (route ItemCreateRoute) GetMethod() string {
	return "POST"
}

func (route ItemCreateRoute) Prehandle(context *gin.Context) *response.Response {
	return nil
}

func (route ItemCreateRoute) Handle(context *gin.Context) *response.Response {
	body, err := context.GetRawData()
	if err != nil {
		return response.CreateFatalResponse("Failed to read request body.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "Failed to read request body.", err)}, 400)
	}

	var item = &models.Item{}

	err = item.Unmarshal(body)
	if err != nil {
		return response.CreateFatalResponse("Failed to parse request body.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "Failed to parse request body.", err)}, 400)
	}

	verify := item.Verify()
	if verify != nil {
		return response.CreateFatalResponse("Failed to parse request body.", []*response.Error{verify}, 400)
	}

	database.GetDB().Save(item)

	data, err := item.Marshal()
	if err != nil {
		return response.CreateWarningResponse("Successfully created item.", nil, []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to marshal item", err)}, 201)
	}

	return response.CreateSuccessResponse("Successfully created item.", data, 201)
}
