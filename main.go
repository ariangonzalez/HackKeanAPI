package main

import (
	"fmt"
	"github.com/labstack/echo"
	"hackkean/db"
	mw "hackkean/middleware"
	"hackkean/api"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Use(mw.TransactionHandler(db.Init()))


	e.GET("/getAllUsers", api.GetUsers())
	e.POST("/registerUser", api.RegisterUser())
	e.Logger.Fatal(e.Start(":1323"))

}

