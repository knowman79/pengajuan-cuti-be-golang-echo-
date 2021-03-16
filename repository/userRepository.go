package repository

import (
	"example/driver"
	"example/model"
	"fmt"
)

func ReadAll() []model.UserModel {
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []model.UserModel

	items, err := db.Query("select id,nama_user,username,password,id_role,phone,email from tb_user")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = model.UserModel{}
		var err = items.Scan(&each.Id, &each.NamaUser, &each.Username, &each.Password, &each.IdRole, &each.Phone, &each.Email)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result

}

func SaveUser(U *model.UserModel) *model.ResponseModel {
	Res := &model.ResponseModel{500, "Internal Server Error"}
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("insert into tb_user values (?, ?, ?, ?, ?, ?, ?)", U.Id, U.NamaUser, U.Username, U.Password, U.IdRole, U.Phone, U.Email)
	if err != nil {
		fmt.Println(err.Error())
		Res = &model.ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("insert success!")
	Res = &model.ResponseModel{200, "Success save Data"}
	return Res
}

func UpdateUser(U *model.UserModel) *model.ResponseModel {
	Res := &model.ResponseModel{500, "Internal Server Error"}
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("update tb_user set nama_user = ?, username = ?, password = ? , id_role = ? , phone = ?, email = ? where id = ?", U.NamaUser, U.Username, U.Password, U.IdRole, U.Phone, U.Email, U.Id)
	if err != nil {
		fmt.Println(err.Error())
		Res = &model.ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &model.ResponseModel{200, "Success save Data"}
	return Res
}

func DeleteUser(Id int) *model.ResponseModel {
	Res := &model.ResponseModel{500, "Internal Server Error"}
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("delete from tb_user where id = ?", Id)
	if err != nil {
		fmt.Println(err.Error())
		Res = &model.ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Delete success!")
	Res = &model.ResponseModel{200, "Success save Data"}
	return Res
}
