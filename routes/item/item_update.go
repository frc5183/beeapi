package routes

import (
	"beeapi/database"
	"beeapi/models"
	"beeapi/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ItemUpdateRoute struct{}

func (route ItemUpdateRoute) GetPath() string {
	// todo: change this to /item/:id, current router does not allow for it
	return "/item"
}

func (route ItemUpdateRoute) GetMethod() string {
	return "PATCH"
}

func (route ItemUpdateRoute) Prehandle(context *gin.Context) *response.Response {
	return nil
}

func (route ItemUpdateRoute) Handle(context *gin.Context) *response.Response {
	var item = &models.Item{}

	resp := getItem(context, item)
	if resp != nil {
		return resp
	}

	body, err := context.GetRawData()
	if err != nil {
		return response.CreateFatalResponse("Failed to read request body.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "Failed to read request body.", err)}, 400)
	}

	err = item.Unmarshal(body)
	if err != nil {
		return response.CreateFatalResponse("Failed to parse request body.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "Failed to parse request body.", err)}, 400)
	}

	verify := item.Verify()
	if verify != nil {
		return response.CreateFatalResponse("Failed to parse request body.", []*response.Error{verify}, 400)
	}

	dbResp := database.GetDB().Save(item)
	if dbResp.Error != nil {
		return response.CreateFatalResponse("Failed to update item.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to update item.", dbResp.Error)}, 500)
	}

	data, err := item.Marshal()
	if err != nil {
		return response.CreateWarningResponse("Successfully updated item.", nil, []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to marshal item", err)}, 201)
	}

	return response.CreateSuccessResponse("Successfully updated item.", data, 201)
}

func getItem(context *gin.Context, item *models.Item) *response.Response {
	if idString, exists := context.GetQuery("id"); exists {
		id, err := strconv.Atoi(idString)
		if err != nil {
			return response.CreateFatalResponse("ID is not valid.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "id is not a uint", err)}, 400)
		}

		if id < 0 {
			return response.CreateFatalResponse("ID is not valid.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "id is not a uint", err)}, 400)
		}

		resp := database.GetDB().Find(item, "id = ?", id)

		if resp.Error != nil {
			return response.CreateFatalResponse("Failed to query item.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to query item.", resp.Error)}, 500)
		}

		if resp.RowsAffected == 0 {
			return response.CreateFatalResponse("No item found with ID "+idString, []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No item found.", nil)}, 404)
		}

		return nil
	} else {
		return response.CreateFatalResponse("No ID specified", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No ID specified", nil)}, 400)
	}
}
