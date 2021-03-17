package models

type UserModel struct {
	UserId   int    `json:"user_id" validate:"required" sql:"AUTO_INCREMENT"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `json:"role_id" validate:"required"`
	Division string `json:"division" validate:"required"`
}
