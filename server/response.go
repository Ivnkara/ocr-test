package server

import (
	"fmt"
	"net/http"
)

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.WriteHeader(code)
	fmt.Fprintf(writer, "%s", data)
}
