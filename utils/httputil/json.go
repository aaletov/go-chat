package httputil

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

const (
	maxBodySize = 1048576
)

func ValidateContentType(w http.ResponseWriter, r *http.Request, ctype string) (status int, msg string) {
	if r.Header.Get("Content-Type") == "" {
		msg = "Content-Type is empty"
		status = http.StatusUnsupportedMediaType
		return
	}
	value := r.Header.Get("Content-Type")
	if value != ctype {
		msg = fmt.Sprintf("Content-Type header is not ctype", ctype)
		status = http.StatusUnsupportedMediaType
		return 
	}

	return http.StatusOK, ""
}

func Unmarshal[T any](w http.ResponseWriter, r *http.Request, obj *T) (status int, msg string) {
	reader := http.MaxBytesReader(w, r.Body, maxBodySize)
	body, err := ioutil.ReadAll(reader)

	if err != nil {
		msg = fmt.Sprintf("Unable to read body: %v", err)
		status = http.StatusUnprocessableEntity
		return
	}

	err = json.Unmarshal(body, obj)

	if err != nil {
		msg = fmt.Sprintf("Invalid request body: %v", err)
		status = http.StatusUnprocessableEntity
		return
	}

	return http.StatusOK, ""
}