package repository

import (
	"example/driver"
	"example/models"
	"fmt"
)

type ResponseModel struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func ReadAllUser() []models.UserModel {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.UserModel

	items, err := db.Query("select user_id, name, username, password, role_id from user")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.UserModel{}
		var err = items.Scan(&each.UserId, &each.Name, &each.Username, &each.Password, &each.RoleId)

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

func CreateUser(U *models.UserModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec(`INSERT INTO "user" ("user_id", "name", "username", "password", "role_id") VALUES ($1, $2, $3, $4, $5)`, U.UserId, U.Name, U.Username, U.Password, U.RoleId)

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("insert success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func DeleteUser(userId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("delete from user where user_id = $1", userId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Delete success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func UpdateUser(U *models.UserModel, userId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("update user set name = $1, username = $2, password = $3 , role_id = $4 where user_id = $5", U.Name, U.Username, U.Password, U.RoleId, userId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}
