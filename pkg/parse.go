package pkg

/*Делаем запрос к api сайта . Получаем максимум(сколько дадут) наших заказов, отбираем те, что откгружены
возвращаем их в виде массива */
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mantrade/models"
	"net/http"
	"strconv"
)

func Parse(token string) ([]models.Sql, error) {
	var rezult []models.Sql

	client := &http.Client{}
	for i := 1; i < 4; i++ {
		url := "https://man-trade.ru/api/v2/orders?show_archive=1&per_page=150&page=" + strconv.Itoa(i) + "&extended_response=true"
		//fmt.Printf("%v\n", url)
		//request, err := http.NewRequest("GET", "https://man-trade.ru/api/v2/orders?show_archive=1&per_page=150&page=1&extended_response=true", nil)
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}

		request.Header.Set("Cookie", "customer_id=d1b3f96c-6d71-06a7-d8e1-1291f459bcfb; login=XJ-122; _www_session=BAh7CUkiD3Nlc3Npb25faWQGOgZFVEkiJTAxYzJmNzA4ZDQ4ZmU5MjdjMTMwNjA0YzE3ZDI0NDA0BjsAVEkiEF9jc3JmX3Rva2VuBjsARkkiMEQ4S1FTc251VFJ4NE9ldFVXNGs3SWswdFNOTDBTZ2VyX3ZpY3RUOWRMbDQGOwBGSSIZY3VzdG9tZXJfY3JlZGVudGlhbHMGOwBUSSIBgDhmMDMyOTM0NGU4YTc3ZTExZDc4ZDU1NTE3MTQ2OTE2YzZkY2JhYjIxYTc1ZDEzNDcxMTdiODM1YTU5MjA1YzJjYjIxNWQ2ODYxNDYxNWIzZmIyZTJiMzI2MmUzOTczMmE3NDMwMmFjNTRmZDYwZDRlNDY1YTRlOTI3N2QyN2EyBjsAVEkiHGN1c3RvbWVyX2NyZWRlbnRpYWxzX2lkBjsAVGlw--0ca5e0ca144e4343dd30e0f0be3eca10ba87e5fd; region_id=1; customer_credentials=8f0329344e8a77e11d78d55517146916c6dcbab21a75d1347117b835a59205c2cb215d68614615b3fb2e2b3262e39732a74302ac54fd60d4e465a4e9277d27a2%3A%3A107%3A%3A2023-09-05T15%3A45%3A47%2B03%3A00")

		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		targets := models.Parse{}
		err = json.Unmarshal(body, &targets)
		if err != nil {
			return nil, fmt.Errorf("Ошибка unmarshal результатов парсинга сайта %v", err)
		}
		for _, j := range targets.Orders {
			if j.Order_items[0].Status == "vydano" {
				var tmp models.Sql
				tmp.Id = j.Order_items[0].Id
				tmp.Oem = j.Order_items[0].Oem
				tmp.Detail_name = j.Order_items[0].Detail_name
				tmp.Make_name = j.Order_items[0].Make_name
				tmp.Cost = j.Order_items[0].Cost
				tmp.Qnt = j.Order_items[0].Qnt
				tmp.Comment = j.Order_items[0].Comment
				rezult = append(rezult, tmp)
			}
		}
	}
	return rezult, nil
}
