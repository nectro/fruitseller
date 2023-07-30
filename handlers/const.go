package handlers

import (
	"github.com/google/uuid"
	"strings"
)

type AddRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

// {
//     "firstname":"Root",
//     "lastname":"Admin",
//     "email":"demo@gmail.com",
//     "password":"Demo@1234",
//     "role":"ADMIN"
// }

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// {
//     "email":"demo@gmail.com",
//     "password":"Demo@1234"
// }

// Add the token received to the header file under the variable name of "key"

type ProductAddRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ImageId     string  `json:"imageId"`
	Description string  `json:"description"`
}

// {
//     "name":"lichi",
//     "price":30.00,
//     "imageId":"image151",
//     "description":"its sweet or sour, brown shell and white flesh"
// }

// Only admin has the privileges to add products

type ProductDeleteRequest struct {
	Id uuid.UUID `json:"id"`
}

// {
//     "id":"178881ac-2e24-4778-bc86-f551909cdc64"
// }

// Only admin has the privileges to delete products

type ProductUpdateRequest struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	ImageId     string    `json:"imageId"`
	Description string    `json:"description"`
}

// {
//     "id":"09d32e72-bd9d-4a2a-bcf6-d740a1dd4190",
//     "imageId":"image151ds"
// }
// OR
// {
//     "id":"09d32e72-bd9d-4a2a-bcf6-d740a1dd4190",
//     "price":28.00102
// }

// only give those fields which are to be changed

// Only admin has the privileges to update products

type ProductToCartRequest struct {
	ProdId uuid.UUID `json:"prodId"`
}

// {
//     "prodId":"9752d222-882b-4d2e-a0c0-2a298d5fcb6d"
// }

// FOR PASSWORD CHECKING

func CheckSpecialCharacters(str string) bool {
	specialChars := "!@#$%^&*()_+~"

	for chars := range specialChars {
		if strings.ContainsRune(str, rune(specialChars[chars])) {
			return true
		}
	}
	return false
}

func CheckUpperCase(str string) bool {
	specialChars := "ABCDEFGHIJKLMNOPQRSTUVQXYZ"

	for chars := range specialChars {
		if strings.ContainsRune(str, rune(specialChars[chars])) {
			return true
		}
	}
	return false
}

func CheckLowerCase(str string) bool {
	specialChars := "abcdefghijklmnopqrstuvwxyz"

	for chars := range specialChars {
		if strings.ContainsRune(str, rune(specialChars[chars])) {
			return true
		}
	}
	return false
}

func CheckItem(list []string, item string) bool {
	for idx := range list {
		if item == list[idx] {
			return true
		}
	}
	return false
}
