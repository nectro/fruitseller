package services

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func DbConnect() {
	//setting environmental variables for the password
	SetEnv()

	// connecting the database
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, os.Getenv("password"), dbname)

	db, err = sql.Open("postgres", psqlconn)
	if err == nil {
		fmt.Printf("database Connected Successfully!")
	} else {
		fmt.Printf("Connection Failed!")
	}
}

func DbClose() {
	db.Close()
}
