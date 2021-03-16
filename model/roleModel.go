package model




type RoleModel struct {
	Id 	int 	`json:"id" validate:"required"`
	Role string `json:"role" validate:"required"`

}