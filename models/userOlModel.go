package models

type UserOlModel struct {
    UserId   	int    	`json:"user_id" validate:"required"`
    Name     	string 	`json:"name" validate:"required"`
    Role   		string  `json:"role" validate:"required"`
    Division   	string  `json:"division" validate:"required"`
}