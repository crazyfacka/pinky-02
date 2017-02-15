package models

import (
	"encoding/json"
)

type ProductData struct {
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       string `json:"price"`
}

type Product struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       string `json:"price"`
	Name        string `json:"name"`
}

func NewProductData() (data *ProductData) {
	return &ProductData{}
}

func NewProduct(id, name, description, price string, stock int) (product *Product) {
	return &Product{
		Id:          id,
		Name:        name,
		Price:       price,
		Stock:       stock,
		Description: description,
	}
}

func (product *Product) ToJSON() string {
	json, err := json.Marshal(product)
	if err != nil {
		return ""
	}

	return string(json)
}
