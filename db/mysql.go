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
//root:@tcp(localhost:3306)/mydatabase
func getSession() *dbr.Session {

	//db, err := dbr.Open("mysql",
	//	conf.USER+":"+conf.PASSWORD+"@tcp("+conf.HOST+":"+conf.PORT+")/"+conf.DB,
	//	nil)
	db, err := dbr.Open("mysql","admin:@tcp(localhost:3306)/test",
		nil)

	if err != nil {
		log.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
