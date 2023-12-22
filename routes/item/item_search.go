package routes

import (
	"beeapi/database"
	"beeapi/models"
	"beeapi/response"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"gorm.io/gorm"
)

type ItemSearchRoute struct{}

func (route ItemSearchRoute) GetPath() string {
	return "/item/search"
}

func (route ItemSearchRoute) GetMethod() string {
	return "GET"
}

func (route ItemSearchRoute) Prehandle(context *gin.Context) *response.Response {
	return nil
}

// todo: find a way to do this without loading in all items. also implement limit, offset, and searching by fields
func (route ItemSearchRoute) Handle(context *gin.Context) *response.Response {
	var items []*models.Item
	var matchedItems []*models.Item

	searchTerm, exists := context.GetQuery("search")
	if !exists {
		return response.CreateFatalResponse("No search term specified", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No search term specified", nil)}, 400)
	}

	var resp *gorm.DB
	if includeDeleted, exists := context.GetQuery("includeDeleted"); exists && includeDeleted == "true" {
		resp = database.GetDB().Unscoped().Find(&items)
	} else {
		resp = database.GetDB().Find(&items)
	}

	if resp.Error != nil {
		return response.CreateFatalResponse("Failed to get items.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to get items.", resp.Error)}, 500)
	}

	for _, item := range items {
		if fuzzy.MatchNormalizedFold(searchTerm, item.Name) || fuzzy.MatchNormalizedFold(searchTerm, item.Description) || fuzzy.MatchNormalizedFold(searchTerm, item.Retailer) || fuzzy.MatchNormalizedFold(searchTerm, item.PartNumber) || fuzzy.MatchNormalizedFold(searchTerm, item.Location) {
			matchedItems = append(matchedItems, item)
		}
	}

	if len(matchedItems) == 0 {
		return response.CreateSuccessResponse("Successfully found items.", nil, 200)
	}

	var marshalledItems []*interface{}
	for _, item := range matchedItems {
		data, err := item.Marshal()
		if err != nil {
			return response.CreateFatalResponse("Failed to marshal item.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to marshal item.", err)}, 500)
		}
		marshalledItems = append(marshalledItems, &data)
	}

	return response.CreateSuccessResponse("Successfully found items.", marshalledItems, 200)
}
