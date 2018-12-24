package models


type Account struct {
	Id   int    `db:"id", json:"id"`
	Username string `db:"username", json:"username"`
	Password string `db:"password", json:"password"`
	Salt     string `db:"salt", json:"salt"`
	Status    string `db:"status", json:"status"`
	CreatedAt int `db:"createdAt", json:"createdAt"`
	UpdatedAt int `db:"updatedAt", json:"updatedAt"`
}