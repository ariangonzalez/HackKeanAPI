package middleware

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func TransactionHandler(db *dbr.Session) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			tx, _ := db.Begin()

			c.Set("tx", tx)

			if err := next(c); err != nil {
				tx.Rollback()
				log.Debug("Transction Rollback: ", err)
				return err
			}
			log.Debug("Transaction Commit")
			tx.Commit()

			return nil
		})
	}
}
