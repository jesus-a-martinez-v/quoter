package repository

import (
	"database/sql"
	"fmt"
	"quoter/src/api/config/db"
	"quoter/src/api/config/loggers"
	"quoter/src/api/model/domain"
)

func GetQuotes(author string, genre string) []domain.QuoteEntity {
	statement := `
		SELECT id, quote, author, genre 
		FROM quotes
		WHERE (($1 = '') OR author = $1)
		  AND (($2 = '') OR genre = $2)`
	results, err := db.Database.Query(statement, author, genre)

	if err != nil {
		// TODO Handle in a better way
		panic(err)
	}

	defer results.Close()

	return toEntities(results)
}

func InsertQuote(entity *domain.QuoteEntity) (int64, error) {
	const insertStatement = `INSERT INTO quotes(quote, author, genre) VALUES($1, $2, $3) RETURNING id`

	insertedRowId := int64(0)
	err := db.Database.QueryRow(insertStatement, entity.Quote, entity.Author, entity.Genre).Scan(&insertedRowId)

	if err != nil {
		loggers.Error.Println("Couldn't save quote. REASON:", err)
		return -1, err
	}

	return insertedRowId, nil
}

func GetQuoteById(id int64) *domain.QuoteEntity {
	const statement = `
		SELECT id, quote, author, genre 
		FROM quotes 
		WHERE id = $1`

	result := db.Database.QueryRow(statement, id)

	if result == nil {
		return nil
	}

	return toEntity(result)
}

func toEntities(results *sql.Rows) []domain.QuoteEntity {
	var quotes []domain.QuoteEntity

	for results.Next() {
		var quote domain.QuoteEntity
		err := results.Scan(
			&quote.Id,
			&quote.Quote,
			&quote.Author,
			&quote.Genre)

		if err != nil {
			fmt.Println("Couldn't process row.")
			continue
		}

		quotes = append(quotes, quote)
	}

	return quotes
}

func toEntity(result *sql.Row) *domain.QuoteEntity {
	var quote domain.QuoteEntity
	err := result.Scan(
		&quote.Id,
		&quote.Quote,
		&quote.Author,
		&quote.Genre)

	if err != nil {
		loggers.Info.Println("Couldn't process row.", err)
		return nil
	}

	return &quote
}
