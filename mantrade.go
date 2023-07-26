package main

import (
	"fmt"
	"log"
	"mantrade/pkg"
)

//"bytes"

func main() {
	/*token, err := pkg.Login() //логинимся, получаем token
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("token=%v\n", token)*/
	//c, err := pkg.Parse(token) //парсим сайт, получаем список последних заказов
	token := "aa"
	shipped, err := pkg.Parse(token) //парсим сайт, получаем список отгруженных заказов
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Parse=%v\n", shipped)
	db, err := pkg.SqlLogin() //Открываем базу
	if err != nil {
		log.Fatal(err)
	}
	/*err = pkg.AddNew(c, db) //Загружаем в базу новые и изменённые
	if err != nil {
		log.Fatal(err)
	}*/
	shippedBase, err := pkg.Base(db) //получаем из базы все ранее неотгруженные в []models.Sql
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Unshipped %v", shippedBase)
	new := pkg.CheckNew(shipped, shippedBase)
	if len(new) < 1 {
		fmt.Printf("Ничего нового отгруженного \n")

	} else {
		fmt.Printf("New:=%v\n", new)
		pkg.Exel(new)
		err = pkg.AddNew(new, db) //Загружаем в базу новые и изменённые
		if err != nil {
			log.Fatal(err)
		}
	}
	/*
		shippedunshipped, err := pkg.Shippedunshipped(c, unshipped, token) //Получаем отгруженные на сайте, но неотгруженные в базе
			if err != nil {
				log.Fatal(err)
			}
			if len(shippedunshipped) != 0 {
				fmt.Printf("Пауза 10 мин\n")
				//time.Sleep(10 * time.Minute)
				shippedunshipped, err := pkg.Shippedunshipped(c, unshipped, token)
				if err != nil {
					log.Fatal(err)
				}
				pkg.Exel(shippedunshipped)             //выгружаем отгруженные в exel
				pkg.ChangeStatus(db, shippedunshipped) //меняем статус у вновь отгруженных в базе

			}
			//fmt.Printf("рез=%v\n", shippedunshipped)
			err = pkg.ClearOld(db)
			if err != nil {
				log.Fatal(err)
			}

			/*for _, i := range b {
				if !pkg.Find(i.Id) { //если для элемента списка нет записи в базе - записываем в базу
					pkg.Add(i)
					fmt.Printf("нет в базе\n")
				}
			}
			d := pkg.NotReady() //грузим из базы все со статусом неотгружено
			//pkg.Separate(a,d) //проверяем все неотгруженные из базы на разделение
			//fmt.Printf("из базы все со статусом неотгружено %v\n", d)
			f, err := pkg.SearchProducts(a, d) //f=заказы, которые в базе не отгружен, а на сайте отгружен
			if err != nil {
				log.Fatal(err)
			}
			if len(f) != 0 { //если есть те, что на сайте отгружены а в базе нет
				//time.Sleep(10 * time.Minute)
				f, err := pkg.SearchProducts(a, d) //f=заказы, которые в базе не отгружен, а на сайте отгружен
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Меняем статус, выгружаем в ексель %v\n", f)
				pkg.ChangeStatus(f) //меняем стаус у вновь отгруженных в базе
				pkg.Exel(f)         //выгружаем отгруженные в exel

			}*/

	//pkg.ClearOld(db) //удаляем записи старше месяца из базы
	pkg.SqlLogout(db)
	/*err = pkg.Logout(token) //отлогиниваемся
	if err != nil {
		log.Fatal(err)
	}*/

}

/*запускать каждый час
читаем по апи заказы за два последних дня
проверяем каждый из них, если заказа нет в нашей базе - парсим сайт чтобы найти примечание и добавляем
заказ в базу со статусом не отгружено

Для каждого неотгруженного из базы проверяем по апи, не изменился ли с татус на выполнен
Если изменился - ждём 10 минут
Для каждого неотгруженного из базы проверяем по апи, не изменился ли статус на выполнен
Если изменился - закидываем его в накладную и меняем его в базе на отгружено
отправляем накладную на почту. Удаляем в базе заказы старше 2 месяца

пробегаем по всем Неотгруженным (разделить могут только их)заказам в базе, сверяем количество.
Если изменилось - то заказ разделили, меняем в базе func separate




Количество брать из Статуса*/
