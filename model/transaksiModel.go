package model


type TransaksiModel struct{
	Id 				int 	`json:"id" validate:"required"`
	IdUser 			string 	`json:"idUser" validate:"required"`
	KodeMakanan 	string 	`json:"KodeMakanan" validate:"required"`
	JumlahMakanan 	string 	`json:"jumlahMakanan" validate:"required"`
	CreatedAt 		int 	`json:"createdAt" validate:"required"`
	JumlahUang 		string 	`json:"jumlahUang" validate:"required"`
}