package api

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"hackkean/model"
	"net/http"
)

func GetUsers() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		tx := c.Get("tx").(*dbr.Tx)

		users := new(model.Users)
		if err = users.Load(tx); err != nil {
			log.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, "No Users")
		}

		return c.JSON(http.StatusOK, users)

	}

}

func RegisterUser() echo.HandlerFunc{

	return func(c echo.Context) error {

		u := new(model.User)
		c.Bind(&u)

		tx := c.Get("tx").(*dbr.Tx)

		newUser := model.NewUser(u.Username, u.Email, u.Phone)
		username := newUser.Username
		if err := newUser.Save(tx); err != nil {
			log.Debug(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		user := new(model.User)
		if err1 := user.Load(tx,username); err1 != nil {
			log.Debug(err1)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to register user")
		}

		return c.JSON(http.StatusCreated, user)

	}
}

//func getFriends() echo.HandlerFunc {
//
//
//	return func(c echo.Context) (err error) {
//
//
//	}
//}
//
//func addFriends() echo.HandlerFunc {
//
//	return func(c echo.Context) (err error) {
//
//
//	}
//
//}
//
//func getPosts() echo.HandlerFunc{
//
//	return func(c echo.Context) (err error){
//
//
//	}
//
//}
//
//func submitPosts() echo.HandlerFunc{
//
//	return func(c echo.Context) (err error){
//
//
//	}
//
//}

