package datastore

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Db struct {
	Data [][]string
	path string
}

// Init(path string): Initialize and load the data of the csv database file
func (db *Db) Init(path string) error {
	db.path = path
	csvFile, err := os.Open(db.path)
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

// WriteLine(line []string): Append a new line in the end of the csv database file
func (db *Db) WriteLine(line []string) error {
	db.Data = append(db.Data, line)

	f, err := os.Create(db.path)
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
