package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quoter/src/api/service"
	"strconv"
)

func GetQuotes(context *gin.Context) {
	author := context.Query("author")
	genre := context.Query("genre")

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
