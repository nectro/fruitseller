package handlers

import (
	"github.com/google/uuid"
)

type AddRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProductAddRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ImageId     string  `json:"imageId"`
	Description string  `json:"description"`
}

type ProductDeleteRequest struct {
	Id uuid.UUID `json:"id"`
}

type ProductUpdateRequest struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	ImageId     string    `json:"imageId"`
	Description string    `json:"description"`
}

type ProductToCartRequest struct {
	ProdId uuid.UUID `json:"prodId"`
}
