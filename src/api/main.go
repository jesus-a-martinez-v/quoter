package main

import (
	"quoter/src/api/config/constants"
	"quoter/src/api/config/db"
	"quoter/src/api/config/loggers"
	"quoter/src/api/config/server"
	"quoter/src/api/model"
)

func main() {
	loggers.Init()
	db.ConnectAndSetDatabase()
	db.Init()
	model.PopulateDb(constants.AllQuotesCsvFilePath)
	server.InitRouter()
}
