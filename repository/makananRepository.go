package repository

import (
	"fmt"
	"time"

	"github.com/my/repo/model"
	"github.com/my/repo/driver"
)

func ReadAllMakanan() []model.MakananModel {
	db, err := driver.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	var result []model.MakananModel
	items, err := db.Query("select id,kode_makanan,nama_makanan,harga,stok,created_at,updated_at,expired_at,update_by,create_by from tb_makanan")
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	for items.Next() {
		var row = model.MakananModel{}
		err = items.Scan(&row.Id, &row.KodeMakanan, &row.NamaMakanan, &row.Harga, &row.Stok, &row.CreatedAt, &row.UpdatedAt, &row.ExpiredAt, &row.UpdateBy, &row.CreateBy)
		if err != nil {
			fmt.Print(err.Error())
			return nil
		}
		result = append(result, row)
	}
	return result
}

func SaveMakanan(request *model.MakananModel) model.ResponseModel {
	db, err := driver.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	response := model.ResponseModel{400, "Bad Request"}
	_, err = db.Exec("insert into tb_makanan values(?,?,?,?,?,?,?,?,?,?)", request.Id, request.KodeMakanan, request.NamaMakanan, request.Harga, request.Stok, time.Now(), time.Now(), request.ExpiredAt, request.UpdateBy, request.CreateBy)
	if err != nil {
		fmt.Println(err.Error())
		response = model.ResponseModel{400, "Failed save Data"}
		return response
	}
	fmt.Println("insert success!")
	response = model.ResponseModel{200, "Success save Data"}
	return response
}

func UpdateMakanan(request *model.MakananModel) model.ResponseModel {
	db, err := driver.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	response := model.ResponseModel{400, "Bad Request"}
	_, err = db.Exec("update tb_makanan set  kode_makanan = ?,nama_makanan = ?, harga = ?,stok = ?, created_at = ?, updated_at = ?,expired_at = ?,update_by = ?,create_by = ? where id = ?",
		request.KodeMakanan, request.NamaMakanan, request.Harga, request.Stok, time.Now(), time.Now(), request.ExpiredAt, request.UpdateBy, request.CreateBy, request.Id)
	if err != nil {
		fmt.Println(err.Error())
		response = model.ResponseModel{400, "Failed update Data"}
		return response
	}
	fmt.Println("update success!")
	response = model.ResponseModel{200, "Success update Data"}
	return response
}
func DeleteMakanan(id int) model.ResponseModel {
	db, err := driver.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	response := model.ResponseModel{400, "Bad Request"}
	fmt.Print("ID : ", id)
	_, err = db.Exec("delete from tb_makanan where id = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		response = model.ResponseModel{400, "Failed delete Data"}
		return response
	}
	fmt.Println("dlete success!")
	response = model.ResponseModel{200, "Success delete Data"}
	return response
}
