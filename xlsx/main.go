package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"reflect"
	"time"
)

func WriteXlsx(sheet string, records interface{}) *excelize.File {
	xlsx := excelize.NewFile()    // new file
	index := xlsx.NewSheet(sheet) // new sheet
	xlsx.SetActiveSheet(index)    // set active (default) sheet
	firstCharacter := 65          // start from 'A' line
	t := reflect.TypeOf(records)

	if t.Kind() != reflect.Slice {
		return xlsx
	}

	s := reflect.ValueOf(records)

	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i).Interface()
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		for j := 0; j < elemType.NumField(); j++ {
			field := elemType.Field(j)
			tag := field.Tag.Get("xlsx")
			name := tag
			column := string(firstCharacter + j)
			if tag == "" {
				continue
			}
			// 设置表头
			if i == 0 {
				xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+1), name)
			}
			// 设置内容
			xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+2), elemValue.Field(j).Interface())
		}
	}
	return xlsx
}

/*********************** usage ***********************/

type Data struct {
	ID   int64     `json:"id" xlsx:"ID 主键"`
	Name string    `json:"name" xlsx:"名称"`
	Age  string    `json:"age" xlsx:"年龄"`
	Date time.Time `json:"time" xlsx:"日期"`
}

func main() {
	list := []Data{
		{1, "a", "44", time.Now()},
		{2, "a", "45", time.Now()},
		{3, "a", "46", time.Now()},
		{4, "aaa", "47", time.Now()},
	}
	// 创建一个工作表
	f := WriteXlsx("Sheet1", list)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		println(err.Error())
	}
}
