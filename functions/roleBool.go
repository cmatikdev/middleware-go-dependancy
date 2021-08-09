package functions

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func ValidateRole(id uint32, role string) (bool, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open(os.Getenv("DB_DRIVER"), DBURL)
	defer db.Close()
	if err != nil {
		fmt.Printf("No se puede conectar a la base de datos", os.Getenv("DB_DRIVER"))
		return false, err
	} else {
		fmt.Printf("Conexión a  %s database", os.Getenv("DB_DRIVER"))
	}
	var per string
	rows, err := db.Table("users").
		Select("p2.name").
		Joins("inner join roles r2 on (r2.id = ru.role_id)").
		Joins("inner join permission_roles pr on (pr.role_id = r2.id)").
		Joins("inner join permissions p2 on (p2.id = pr.permission_id)").
		Where("where ru.user_id = ?;", id).
		Rows()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false, err
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&per)
		fmt.Printf("per: %v\n", per)
		if per == role {
			return true, nil
		}
	}
	return false, errors.New("no existe el rol en DBB")

}
