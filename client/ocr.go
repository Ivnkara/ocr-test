package client

import (
	"os"

	ocr "github.com/ranghetto/go_ocr_space"
)

type OcrClient struct {
    ApiKey     string
	Lang       string
	OcrService ocr.Config
}

func NewOcrClient() *OcrClient {
	app := &OcrClient{
		ApiKey: os.Getenv("API_KEY"),
		Lang: os.Getenv("APP_LANG"),
	}

	app.OcrService = ocr.InitConfig(app.ApiKey, app.Lang)

	return app
}

func (c *OcrClient) FromFile(filePath string) (result string, err error) {
	response, err := c.OcrService.ParseFromLocal(filePath)
	if err != nil {
		return "", err
	}

	result = response.JustText()
	
	return
}

func (c *OcrClient) FromUrl(url string) (result string, err error) {
	response, err := c.OcrService.ParseFromUrl(url)
	if err != nil {
		return "", err
	}

	result = response.JustText()
	
	return
}
