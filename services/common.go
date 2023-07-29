package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/google/uuid"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "fruitsellerBackend"
)

var db *sql.DB
var err error

type Product struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	ImageId     string    `json:"imageId"`
	Description string    `json:"description"`
}

func generateAuthToken(data, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
