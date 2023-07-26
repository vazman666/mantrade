package pkg

import (
	"mantrade/models"
)

func CheckNew(shipped []models.Sql, shippedBase []models.Sql) []models.Sql {
	var result []models.Sql
	for _, j := range shipped {
		check := true
		for _, i := range shippedBase {
			if j.Id == i.Id {
				check = false
				//fmt.Printf("уже было в базе %v\n", i)
				break
			}

		}
		if check {
			result = append(result, j)
		}
	}
	return result
}
