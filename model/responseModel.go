package model


type ResponseModel struct {
	Code	int		`json:"code" validate:"required"`
	Message	string	`json:"message" validate:"required"`
}