package db

const createTableStatement = `
	CREATE TABLE IF NOT EXISTS quotes(
		id SERIAL PRIMARY KEY, 
		quote TEXT NOT NULL,
		author VARCHAR(250) NOT NULL,
		genre  VARCHAR(100) NOT NULL,
		UNIQUE(quote, author, genre)
	)
	`

var tables = []string{createTableStatement}

var indices = []string{
	"CREATE INDEX IF NOT EXISTS quotes_author ON quotes(author)",
	"CREATE INDEX IF NOT EXISTS quotes_genre ON quotes(genre)",
}

func CreateTables() {
	for _, table := range tables {
		Database.Exec(table)
	}
}

func CreateIndices() {
	for _, index := range indices {
		Database.Exec(index)
	}
}

func Init() {
	CreateTables()
	CreateIndices()
}
