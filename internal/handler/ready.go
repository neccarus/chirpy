package handler

import (
	"fmt"
	"net/http"
)

func Ready(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "text/plain; charset=utf-8")
	_, err := writer.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}
