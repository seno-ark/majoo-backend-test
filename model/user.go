package model

type UserLogin struct {
	Token string `json:"token"`
	User
}

type User struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	UserName string `db:"user_name" json:"user_name"`
}

type Account struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
}
