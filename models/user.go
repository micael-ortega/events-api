package models

type User struct {
	ID           int    `json:""`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}
