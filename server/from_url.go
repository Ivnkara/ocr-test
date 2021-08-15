package server

import (
	"net/http"
	"github.com/Ivnkara/ocr-test/client"
)

func FromUrl(w http.ResponseWriter, req *http.Request) {
	var err error
	ocr := client.NewOcrClient()

	urls, ok := req.URL.Query()["url"]
	if ok {
		result := map[int]string{}

		for i, url := range urls {
			result[i], err = ocr.FromUrl(url)
			if err != nil {
				writeResponse(w, http.StatusBadRequest, err)

				return
			}
		}

		writeResponse(w, http.StatusOK, result)

		return
	}

	writeResponse(w, http.StatusBadRequest, "Url not found")
}
