package model

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"quoter/src/api/model/domain"
	"quoter/src/api/model/repository"
	"quoter/src/api/config/loggers"
)

func PopulateDb(filePath string) {
	f, err := os.Open(filePath)

	if err != nil {
		loggers.Warning.Println("Could not populate the database", err)
		return
	}

	csvReader := csv.NewReader(bufio.NewReader(f))
	csvReader.Comma = ';'

	recordRead := 0
	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			loggers.Info.Println("Reached end of file.")
			break
		}

		entity := rowToEntity(record)

		repository.InsertQuote(&entity)
		recordRead += 1
	}
}

func rowToEntity(row []string) domain.QuoteEntity {
	return domain.QuoteEntity{
		Quote: row[0],
		Author: row[1],
		Genre: row[2],
	}
}
