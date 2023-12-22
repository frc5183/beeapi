package routes

import (
	"beeapi/database"
	"beeapi/models"
	"beeapi/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ItemDeleteRoute struct{}

func (route ItemDeleteRoute) GetPath() string {
	// todo: change this to /item/:id, current router does not allow for it
	return "/item"
}

func (route ItemDeleteRoute) GetMethod() string {
	return "DELETE"
}

func (route ItemDeleteRoute) Prehandle(context *gin.Context) *response.Response {
	return nil
}

func (route ItemDeleteRoute) Handle(context *gin.Context) *response.Response {
	if idString, exists := context.GetQuery("id"); exists {
		id, err := strconv.Atoi(idString)
		if err != nil {
			return response.CreateFatalResponse("ID is not valid.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "id is not a uint", err)}, 400)
		}

		if id < 0 {
			return response.CreateFatalResponse("ID is not valid.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "id is not a uint", err)}, 400)
		}

		var item = &models.Item{}
		resp := database.GetDB().Delete(item, "id = ?", id)

		if resp.Error != nil {
			return response.CreateFatalResponse("Failed to delete item.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to delete item.", resp.Error)}, 500)
		}

		if resp.RowsAffected == 0 {
			return response.CreateFatalResponse("No item found with ID "+idString, []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No item found.", nil)}, 404)
		}

		return response.CreateSuccessResponse("Item deleted successfully.", nil, 200)
	} else {
		return response.CreateFatalResponse("No ID specified", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No ID specified", nil)}, 400)
	}
}
