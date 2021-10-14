package datastore

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Db struct {
	Data [][]string
	url  string
}

func (db *Db) Init(url string) error {
	db.url = url
	csvFile, err := os.Open(db.url)
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

func (db *Db) WriteLine(line []string) error {
	db.Data = append(db.Data, line)

	f, err := os.Create(db.url)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range db.Data {
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}
