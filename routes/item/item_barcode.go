package routes

import (
	barcode2 "beeapi/barcode"
	"beeapi/database"
	"beeapi/models"
	"beeapi/response"
	"github.com/gin-gonic/gin"
	"image/png"
	"strconv"
)

type ItemBarcodeRoute struct{}

func (route ItemBarcodeRoute) GetPath() string {
	return "/item/barcode"
}

func (route ItemBarcodeRoute) GetMethod() string {
	return "GET"
}

func (route ItemBarcodeRoute) Prehandle(context *gin.Context) *response.Response {
	return nil
}

func (route ItemBarcodeRoute) Handle(context *gin.Context) *response.Response {
	if idString, exists := context.GetQuery("id"); exists {
		id, err := strconv.Atoi(idString)
		if err != nil {
			return response.CreateFatalResponse("ID is not valid.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "id is not a uint", err)}, 400)
		}

		if id < 0 {
			return response.CreateFatalResponse("ID is not valid.", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "id is not a uint", err)}, 400)
		}

		var item = &models.Item{}
		resp := database.GetDB().Find(item, "id = ?", id)

		if resp.Error != nil {
			return response.CreateFatalResponse("Failed to query item.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to query item.", resp.Error)}, 500)
		}

		if resp.RowsAffected == 0 {
			return response.CreateFatalResponse("No item found with ID "+idString, []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No item found.", nil)}, 404)
		}

		code, err := barcode2.GenerateItemBarcode(item)
		if err != nil {
			return response.CreateFatalResponse("Failed to generate barcode.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to generate barcode.", err)}, 500)
		}

		err = png.Encode(context.Writer, code)
		if err != nil {
			return response.CreateFatalResponse("Failed to encode barcode.", []*response.Error{response.CreateError(response.ErrorCodeInternalServerError, "Failed to encode barcode.", err)}, 500)
		}

		context.Data(200, "image/png", nil)
		return nil
	} else {
		return response.CreateFatalResponse("No ID specified", []*response.Error{response.CreateError(response.ErrorCodeInvalidRequest, "No ID specified", nil)}, 400)
	}
}
