package model

import (
	"net/http"
)

type ProductRequest struct {
	Name          string  `json:"name" validate:"required"`
	Description    string  `json:"description" validate:"required"`
	ImageUrl          string  `json:"image_url"`
	Price          float64 `json:"price" validate:"required"`
	
}

func (c *ProductRequest) Bind(r *http.Request) error{
		return nil
}

type ProductUpdateRequest struct {
 	Name          *string  `json:"name,omitempty"`
 	Description    *string  `json:"description,omitempty"`
 	ImageUrl          *string  `json:"image_url,omitempty"`
 	Price          *float64 `json:"price,omitempty"`
}

func (b *ProductUpdateRequest) Bind(r *http.Request) error {
	return nil
}

type ProductResponse struct {
	Id         uint    `json:"id"`
	Name      string  `json:"name"`
	Description string `json:"description"`
	ImageUrl      string  `json:"image_url"`
	Price      float64 `json:"price"`
}



