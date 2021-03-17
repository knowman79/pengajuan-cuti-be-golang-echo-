package models

type RoleModel struct {
	RoleId int    `json:"role_id" validate:"required"`
	Role   string `json:"role" validate:"required"`
}
