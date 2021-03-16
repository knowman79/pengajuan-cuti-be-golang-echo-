package model


type MakananModel struct {
	Id 			int 	`json:"id" validate:"required"`
	KodeMakanan string 	`json:"kodeMakanan" validate:"required"`
	NamaMakanan string 	`json:"nama_makanan" validate:"required"`
	Harga int 			`json:"harga" validate:"required"`
	Stok int 			`json:"stok" validate:"required"`
	CreatedAt string 	`json:"createdAt" `
	UpdatedAt string 	`json:"updateAt" `
	ExpiredAt string 	`json:"expiredAt" validate:"required"`
	UpdateBy string 	`json:"updateBy"`
	CreateBy string 	`json:"createdBy"`

}