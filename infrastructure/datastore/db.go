package datastore

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Db struct {
	Data [][]string
}

func (db *Db) Init() error {
	csvFile, err := os.Open("db.csv")
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err
	}
	db.Data = csvLines
	return nil
}
