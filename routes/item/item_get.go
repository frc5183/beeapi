package routes

import (
	"beeapi/database"
	"beeapi/models"
	"beeapi/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ItemGetRoute struct{}

func (route ItemGetRoute) GetPath() string {
	// todo: change this to /item/:id, current router does not allow for it
	return "/item"
}

func (route ItemGetRoute) GetMethod() string {
	return "GET"
}

func (route ItemGetRoute) Prehandle(context *gin.Context) *response.Response {
	return nil
}

func (route ItemGetRoute) Handle(context *gin.Context) *response.Response {
	if idString, exists := context.GetQuery("id"); exists {
		id, err := strconv.Atoi(idString)
		if err != nil {
			return response.CreateFatalResponse("ID is not valid.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "id is not a uint", err)}, 400)
		}

		if id < 0 {
			return response.CreateFatalResponse("ID is not valid.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "id is not a uint", err)}, 400)
		}

		var item = &models.Item{}
		var resp *gorm.DB
		if includeDeleted, exists := context.GetQuery("includeDeleted"); exists && includeDeleted == "true" {
			resp = database.GetDB().Unscoped().Find(item, "id = ?", id)
		} else {
			resp = database.GetDB().Find(item, "id = ?", id)
		}

		if resp.Error != nil {
			return response.CreateFatalResponse("Failed to query item.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to query item.", resp.Error)}, 500)
		}

		if resp.RowsAffected == 0 {
			return response.CreateFatalResponse("No item found with ID "+idString, []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No item found.", nil)}, 404)
		}

		data, err := item.Marshal()
		if err != nil {
			return response.CreateFatalResponse("Failed to marshal item.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to marshal item.", err)}, 500)
		}

		return response.CreateSuccessResponse("Item found successfully.", data, 200)
	} else {
		return response.CreateFatalResponse("No ID specified", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No ID specified", nil)}, 400)
	}
}
