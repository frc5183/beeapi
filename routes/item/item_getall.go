package routes

import (
	"beeapi/database"
	"beeapi/models"
	"beeapi/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ItemGetAllRoute struct{}

func (route ItemGetAllRoute) GetPath() string {
	return "/item/all"
}

func (route ItemGetAllRoute) GetMethod() string {
	return "GET"
}

func (route ItemGetAllRoute) Prehandle(context *gin.Context) *response.Response {
	return nil
}

func (route ItemGetAllRoute) Handle(context *gin.Context) *response.Response {
	var err error
	var items []*models.Item

	limitString := context.DefaultQuery("limit", "-1")
	offsetString := context.DefaultQuery("offset", "-1")

	var limit int
	var offset int

	limit, err = strconv.Atoi(limitString)
	if err != nil {
		return response.CreateFatalResponse("Limit is not valid", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "limit is not an int", err)}, 400)
	}

	offset, err = strconv.Atoi(offsetString)
	if err != nil {
		return response.CreateFatalResponse("Offset is not valid", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "offset is not an int", err)}, 400)
	}

	if limit == -1 && offset > -1 {
		return response.CreateFatalResponse("Limit is not valid", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "limit must be set with offset", nil)}, 400)
	}

	resp := database.GetDB().Limit(limit).Offset(offset).Find(&items)

	if resp.Error != nil {
		return response.CreateFatalResponse("Failed to query items.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to query items.", resp.Error)}, 500)
	}

	if len(items) < 0 {
		return response.CreateSuccessResponse("Successfully queried items.", nil, 200)
	}

	var marshalledItems []*interface{}
	err = marshalItems(&items, &marshalledItems)
	if err != nil {
		return response.CreateFatalResponse("Failed to marshal items", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to marshal items", err)}, 500)
	}

	return response.CreateSuccessResponse("Successfully queried items.", marshalledItems, 200)
}

func marshalItems(items *[]*models.Item, marshalledItems *[]*interface{}) error {
	for _, item := range *items {
		data, err := item.Marshal()
		if err != nil {
			return err
		}

		*marshalledItems = append(*marshalledItems, &data)
	}

	return nil
}
