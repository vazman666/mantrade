package models

var Login = "29771"
var Pass = "23y1sm"

type LoginType struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
	Errors  string `json:"errors"`
	Result  struct {
		Content string `json:"content"`
		Params  struct {
			Success bool   `json:"success"`
			Token   string `json:"token"`
		} `json:"params"`
	} `json:"result"`
}

type Logout struct {
	Success bool `json:"success"`
}
type SearchProducts struct {
	Success bool `json:"success"`
	Result  struct {
		Params struct {
			Success         bool `json="success"`
			HistoryProducts []struct {
				Uuid         string  `json="uuid"`
				OrderUuid    string  `json=orderUuid"`
				CartItemId   string  `json="cart_item_id"`
				Brand        string  `json="brand"`
				BaseNumber   string  `json="baseNumber"`
				Name         string  `json="name"`
				Info         string  `json="info"`
				ExtraNumber  string  `json="extraNumber"`
				Price        float32 `json="price"`
				PricesString string  `json="pricesString"`
				Quantity     int     `json="quantity"`
				Statuses     []struct {
					Name         string `json="name"`
					Description  string `json="description"`
					TimeDelivery string `json="timeDelivery"`
					Quantity     int    `json="quantity"`
					Date         string `json="date"`
				} `json="statuses"`
			} `json'"historyProducts"`
		} `json="params"`
	} `json="result"`
}
type Parse struct {
	Status string `json="status"`
	Orders []struct {
		Order_items []struct {
			Id          int     `json="id"`
			Make_name   string  `json="make_name"` //производитель
			Oem         string  `json="oem"`       //номер детали
			Detail_name string  `json="detail_name"`
			Cost        float32 `json="cost"`    //цена
			Qnt         int     `json="qnt"`     //количество
			Status      string  `json="status"`  //статус заказа
			Comment     string  `json="comment"` //наш коммент

		} `json="order_items"`
	} `json="orders"`
}
type Sql struct {
	Id          int
	Oem         string
	Detail_name string
	Make_name   string
	Cost        float32
	Qnt         int
	Comment     string
	Date        string
}

//MySQL > DELETE FROM `mytable` WHERE `date` < DATE_SUB(NOW(), INTERVAL 1 MONTH);  удалить все записи старше одного месяца

//INSERT INTO table(data) VALUES(NOW(. ;   добавить текущую дату в таблицу
//Тип столбца DATE

//Структура таблицы japautozap/mantrade
// id int
// oem string (15) номер детали
// deatil_name (40) название детали
// make_name string (15) производитель
// cost float      цена
// qnt int  количество
// comment string(40) примечание
// date DATE время записи заказа

// создаём таблицу
// mysql -u vazman -p
// create database japautozap;  - если ещё не создана
// use japautozap
/* create table mantrade (id int,
oem varchar(15),
detail_name varchar(40),
make_name varchar(15),
cost float,
qnt TINYINT UNSIGNED,
comment varchar(40),
date DATE);*/

//ALTER TABLE ivers modify column name varchar(50);   изменяем столбец
//ALTER TABLE ivers modify column number varchar(30);
//mysqldump -u vazman -p japautozap ivers > japautozap_ivers.sql
//mysql -uvazman -prbhgbxb1 japautozap < japautozap_ivers.sql
// update ivers set status=false;
//ALTER TABLE ivers ADD COLUMN quanity2  VARCHAR(10) AFTER status;
//ALTER TABLE ivers ADD COLUMN status2  bool  AFTER quanity2;
//ALTER TABLE ivers DROP COLUMN status2;
//ALTER TABLE ivers ADD COLUMN subid  TINYINT UNSIGNED AFTER id;
//update ivers set status=false where number="C-110";
