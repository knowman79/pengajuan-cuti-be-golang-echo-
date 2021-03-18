package repository

import (
	"example/driver"
	"example/models"
	"fmt"
)

func ReadUserByUsername(username string) []models.UserModel {
	var result []models.UserModel
	db, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return result
	}

	defer db.Close()

	items, err := db.Query(`SELECT tb_user.user_id, tb_user.username, tb_user.password, tb_role.role_id, tb_user.name, tb_user.division FROM tb_user JOIN tb_role ON tb_user.role_id = tb_role.role_id WHERE tb_user.username = $1 `, username)
	if err != nil {
		fmt.Println(err.Error())
		return result
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.UserModel{}
		var err = items.Scan(&each.UserId, &each.Username, &each.Password, &each.RoleId, &each.Name, &each.Division)

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
