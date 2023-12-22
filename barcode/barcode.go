package barcode

import (
	"beeapi/config"
	"beeapi/models"
	_ "embed"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code39"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"image/draw"
	"strconv"
)

// todo: custom font support
//
//go:embed roboto.ttf
var fontData []byte

func GenerateItemBarcode(item *models.Item) (image.Image, error) {
	c39, err := code39.Encode(strconv.Itoa(int(item.ID)), true, true)
	if err != nil {
		return nil, err
	}

	code, err := barcode.Scale(c39, int(config.GetConfig().Barcode.BarcodeWidth), int(config.GetConfig().Barcode.BarcodeHeight))
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rect(0, 0, int(config.GetConfig().Barcode.ImageWidth), int(config.GetConfig().Barcode.ImageHeight)))

	draw.Draw(img, img.Rect, image.NewUniform(color.White), image.Point{}, draw.Src)
	draw.Draw(img, img.Rect, code, image.Point{}, draw.Over)
	err = addLabel(img)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func addLabel(img *image.RGBA) error {
	font, err := truetype.Parse(fontData)
	if err != nil {
		return err
	}

	context := freetype.NewContext()
	context.SetDPI(72)
	context.SetFont(font)
	context.SetFontSize(float64(config.GetConfig().Barcode.LabelSize))
	context.SetClip(img.Bounds())
	context.SetDst(img)
	context.SetSrc(image.Black)

	_, err = context.DrawString(config.GetConfig().Barcode.Label, freetype.Pt(int(config.GetConfig().Barcode.LabelX), int(config.GetConfig().Barcode.LabelY)))
	if err != nil {
		return err
	}

	return nil
}
