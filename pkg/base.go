package pkg

import (
	"database/sql"
	"fmt"
	"mantrade/models"
)

func Base(db *sql.DB) ([]models.Sql, error) {

	var rez []models.Sql
	rows, err := db.Query("select * from mantrade")
	if err != nil {
		return nil, fmt.Errorf("Ошибка Select при работе с базой %v", err)
	}

	var tmp models.Sql
	for rows.Next() {
		err := rows.Scan(&tmp.Id, &tmp.Oem, &tmp.Detail_name, &tmp.Make_name, &tmp.Cost, &tmp.Qnt,
			&tmp.Comment, &tmp.Date)
		if err != nil {
			return nil, fmt.Errorf("Ошибка Scan при работе с базой %v", err)
		}
		//fmt.Printf("Скан с базы %v\n", tmp)
		rez = append(rez, tmp)

	}
	defer rows.Close()
	return rez, nil
}
