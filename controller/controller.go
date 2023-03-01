package controller

import (
	"io"
	"net/http"
)

func requestBodyToString(r *http.Request) ([]byte, error) {
	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		return nil, error
	}

	return requestBody, nil
}
