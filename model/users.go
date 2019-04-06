package model

import (
	"github.com/gocraft/dbr"
)

type User struct {
	Uid   int64  `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}


func NewUser(username string, email string, phone string) *User {
	return &User{
		Username:username,
		Email:email,
		Phone:phone,
	}
}

func (u *User) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("Users").
		Columns( "username", "email", "phone").
		Record(u).
		Exec()

	return err
}

func (u *User) Load(tx *dbr.Tx, username string) error {

	return tx.Select("*").
		From("Users").
		Where("username = ?", username).LoadOne(u)
}



type Users []User

func (u *Users) Load(tx *dbr.Tx) error {

	_, err := tx.Select("*").
		From("Users").Load(u)
	return err
}

