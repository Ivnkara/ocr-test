package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Ivnkara/ocr-test/client"
)

func FromFile(w http.ResponseWriter, req *http.Request) {
	filepath, _ := getFilePathFromRequest(req)
	defer removeFile(filepath)

	ocr := client.NewOcrClient()

	result, err := ocr.FromFile(filepath)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)

		return
	}

	writeResponse(w, http.StatusOK, result)
}

func getFilePathFromRequest(req *http.Request) (string, error) {
	req.ParseMultipartForm(1 << 20)

	file, handler, err := req.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	filepath := "./temp-images/" + handler.Filename

	err = ioutil.WriteFile(filepath, fileBytes, 0644)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func removeFile(file string) error {
	err := os.Remove(file)
	if err != nil {
		fmt.Println(665, err)
		return err
	}

	return nil
}
