package db

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/gommon/log"
	_ "github.com/go-sql-driver/mysql"
)

func Init() *dbr.Session {

	session := getSession()

	return session
}

func getSession() *dbr.Session {

	db, err := dbr.Open("mysql","root:@tcp(localhost:3306)/HackKean",
		nil)

	if err != nil {
		log.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
