package errors

import (
	"encoding/json"
	"net/http"
)

type ProductNotFound struct {
	Message string `json:"message"`
}

func NewProductNotFound(err error) *ProductNotFound {
	return &ProductNotFound{
		Message: err.Error(),
	}
}

func (*ProductNotFound) GetError() int {
	return http.StatusNotFound
}

func (foundError *ProductNotFound) ToJSON() string {
	json, err := json.Marshal(foundError)
	if err != nil {
		return ""
	}

	return string(json)
}
