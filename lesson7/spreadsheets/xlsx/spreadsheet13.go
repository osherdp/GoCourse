package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test13.xlsx"

func main() {
	// konstrukce struktury typu File
	worksheet := xlsx.NewFile()

	fmt.Println(worksheet)

	// přidání listu do sešitu
	sheet, err := worksheet.AddSheet("Tabulka1")
	if err != nil {
		panic(err)
	}
	defer sheet.Close()

	fmt.Println(sheet)

	// přidání řádku do tabulky
	row := sheet.AddRow()

	// přidání buňky na řádek
	cell := row.AddCell()

	// naplnění buňky hodnotou
	cell.SetInt(42)

	// pokus o uložení sešitu
	err = worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
