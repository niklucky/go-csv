package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Mapper struct {
	Path string
}

func (csvMapper *Mapper) Save(dir string, filename string, data [][]string, header []string) error {
	var records [][]string
	dest := csvMapper.Path + "/" + dir
	_, err := os.Stat(dest)
	if err != nil {
		if os.IsNotExist(err) {
			er := os.MkdirAll(dest, 0775)
			if er != nil {
				panic(er)
			}
		} else {
			fmt.Println("Create directory:", err)
		}
	}
	fname := dest + "/" + filename + ".csv"
	_, err = os.Stat(fname)
	if os.IsNotExist(err) {
		if len(header) > 0 {
			records = append(records, header)
		}
	}
	if err != nil {
		fmt.Println("Error stat:", err)
	}
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, v := range data {
		records = append(records, v)
	}
	e := writer.WriteAll(records)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("CSV	| ", filename, "	| ", len(data), "\n==================================")
	return nil
}
