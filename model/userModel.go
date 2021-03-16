package model


type UserModel struct {
	Id 			int 	`json:"id" validate:"required" sql:"AUTO_INCREMENT"`
	NamaUser 	string 	`json:"namaUser" validate:"required"`
	Username 	string 	`json:"username" validate:"required"`
	Password 	string 	`json:"password" validate:"required"`
	IdRole 		int 	`json:"idRole" validate:"required"`
	Phone 		string 	`json:"phone"`
	Email 		string 	`json:"email"`

}