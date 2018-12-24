package models


type Account struct {
	id   int    `db:"id"`
	username string `db:"username"`
	password string `db:"password"`
	salt     string `db:"salt"`
	status    string `db:"status"`
	createdAt int `db:"createdAt"`
	updatedAt int `db:"updatedAt"`
}