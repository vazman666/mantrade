package pkg

import (
	"mantrade/models"

	"github.com/tealeg/xlsx/v3"
)

func Exel(data []models.Sql) {
	//fmt.Printf("exeldata=%v\n", data)
	type xlsxstruct struct {
		Firm       string
		PartNumber string
		Name       string
		Price      float32
		Quantity   int
		Summ       float32
		Remark     string
	}
	var xlsdata []xlsxstruct
	for _, i := range data {
		var tmp xlsxstruct
		tmp.Firm = i.Make_name
		tmp.PartNumber = i.Oem
		tmp.Name = i.Detail_name
		/*value, _ := strconv.ParseFloat(i.Price, 32)
		tmp.Price = float32(value)*/
		tmp.Price = i.Cost
		//tmp.Quantity, _ = strconv.Atoi(i.Quantity)
		tmp.Quantity = i.Qnt
		tmp.Summ = tmp.Price * float32(tmp.Quantity)
		tmp.Remark = i.Comment
		xlsdata = append(xlsdata, tmp)
	}

	wb := xlsx.NewFile() //создаём новый экскиз экселя

	sheetTest, err := wb.AddSheet("Sheet") //добавляем страничку
	if err != nil {
		panic(err)
	}
	sheetTest.SetColWidth(1, 1, 8)  // c 1 по 1 ширина 8
	sheetTest.SetColWidth(2, 2, 16) // с 2 по 3 ширина 16
	sheetTest.SetColWidth(3, 3, 55)
	sheetTest.SetColWidth(4, 9, 13)

	row1 := sheetTest.AddRow()

	cell := row1.AddCell()
	cell.SetValue("Firm")
	cell = row1.AddCell()
	cell.SetValue("PartNumber")
	cell = row1.AddCell()
	cell.SetValue("Name")
	cell = row1.AddCell()
	cell.SetValue("Цена")
	cell = row1.AddCell()
	cell.SetValue("Количество")
	cell = row1.AddCell()
	cell.SetValue("Сумма")
	cell = row1.AddCell()
	cell.SetValue("Примечание")
	row1.SetHeight(15)
	row1 = sheetTest.AddRow()
	cell = row1.AddCell()
	cell.SetValue("")

	for _, value := range xlsdata {
		//fmt.Printf("%v\n", value)
		row1 = sheetTest.AddRow()        //добавляем строку
		_ = row1.WriteStruct(&value, -1) //и вставляе в эту строку строку из прайс
		row1.SetHeight(15)

	}
	err = wb.Save("a.xlsx")
	if err != nil {
		panic(err)
	}

}
