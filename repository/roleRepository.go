package repository

import (
	"fmt"

	"github.com/my/repo/driver"
	"github.com/my/repo/models"
)

func ReadAll() []models.RoleModel {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.RoleModel

	items, err := db.Query("select role_id,role from role")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.RoleModel{}
		var err = items.Scan(&each.Id, &each.Role)

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
