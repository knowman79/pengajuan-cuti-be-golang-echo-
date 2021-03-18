package models

type ResponseLogin struct {
	User    UserModel `json:"user"`
	Token   string    `json:"token"`
	Message string    `json:"message"`
}
