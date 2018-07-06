package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quoter/src/api/service"
	"quoter/src/api/service/dto"
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
		return
	}

	ctx.JSON(http.StatusOK, quote)
}

func SaveQuote(ctx *gin.Context) {

	var newQuote dto.QuoteDto

	err := ctx.ShouldBindJSON(&newQuote)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request", err)
	}

	createdQuoteId, err := service.SaveQuote(newQuote)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Couldn't save quote", err)
		return
	}

	ctx.JSON(http.StatusCreated, createdQuoteId)
}
