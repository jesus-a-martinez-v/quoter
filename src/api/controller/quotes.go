package controller

import (
	"net/http"
	"quoter/src/api/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"quoter/src/api/config/loggers"
)

func GetQuotes(context *gin.Context) {
	author := context.Query("author")
	genre := context.Query("genre")

	loggers.Info.Println("QUERY PARAM", context.Query("author"))
	quotes := service.GetQuotes(author, genre)
	context.JSON(http.StatusOK, quotes)
}

func GetQuoteById(ctx *gin.Context) {
	quoteId, _ := strconv.Atoi(ctx.Param("id"))
	quote := service.GetQuoteById(int64(quoteId))

	if quote == nil {
		ctx.String(http.StatusNotFound, "Couldn't find quote with id %d", quoteId)
	}

	ctx.JSON(http.StatusOK, quote)
}