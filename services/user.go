package services

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"time"
)

func AddUser(firstname string, lastname string, password string, email string, role string, generatedId uuid.UUID) {

	currentTime := time.Now().Format(time.RFC3339)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	checkError(err)

	hashedPassword := string(hash)

	addUserQuery := fmt.Sprintf("INSERT INTO USERS (FIRSTNAME,LASTNAME,EMAIL,PASSWORD,ROLE,CARTID,CREATEDAT,UPDATEDAT) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s')", firstname, lastname, email, hashedPassword, role, generatedId, currentTime, currentTime)
	_, er := db.Exec(addUserQuery)
	checkError(er)
}

func AuthenticateUser(email string, password string) string {

	selectUser := fmt.Sprintf("SELECT PASSWORD FROM USERS WHERE EMAIL='%s'", email)
	r, er := db.Query(selectUser)
	checkError(er)
	defer r.Close()

	var hashedPassword string

	for r.Next() {
		r.Scan(&hashedPassword)
	}

	checkFail := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if checkFail == nil {
		randomNumber := rand.Intn(101)
		authToken := generateAuthToken(fmt.Sprint(randomNumber), os.Getenv("secretKey"))
		return authToken
	} else {
		return "failed"
	}
}

func GetUserDetails(email string, data string) string {

	selectUser := fmt.Sprintf("SELECT %s FROM USERS WHERE EMAIL='%s'", data, email)
	r, er := db.Query(selectUser)
	checkError(er)
	defer r.Close()

	var resData string

	for r.Next() {
		r.Scan(&resData)
	}

	return resData
}
