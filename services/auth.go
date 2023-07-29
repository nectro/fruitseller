package services

import (
	"fmt"
)

func AddAuth(token string, email string) bool {

	checkAuth := fmt.Sprintf("SELECT USEREMAIL FROM AUTH WHERE USEREMAIL='%s'", email)
	R, er := db.Query(checkAuth)
	checkError(er)
	defer R.Close()

	var useremail string

	for R.Next() {
		R.Scan(&useremail)
	}

	if useremail == "" {
		addAuth := fmt.Sprintf("INSERT INTO AUTH (USEREMAIL,AUTHTOKEN) VALUES ('%s','%s')", email, token)
		_, er := db.Exec(addAuth)
		checkError(er)
	} else {
		addAuth := fmt.Sprintf("UPDATE AUTH SET AUTHTOKEN='%s' WHERE USEREMAIL='%s'", token, email)
		_, er := db.Exec(addAuth)
		checkError(er)
	}

	if er == nil {
		return true
	} else {
		return false
	}
}

func VerifyAuthToken(token string) string {

	checkAuth := fmt.Sprintf("SELECT USEREMAIL FROM AUTH WHERE AUTHTOKEN='%s'", token)
	R, er := db.Query(checkAuth)
	checkError(er)
	defer R.Close()

	var useremail string

	for R.Next() {
		R.Scan(&useremail)
	}

	return useremail
}
