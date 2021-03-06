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
	random := context.Query("random")

	if len(random) > 0 && (random != "true" && random != "false") {
		context.String(http.StatusBadRequest, "Random must be either true or false if specified.")
		return
	}

	quotes := service.GetQuotes(author, genre, random)
	context.JSON(http.StatusOK, quotes)
}

func GetQuoteById(context *gin.Context) {
	quoteId, _ := strconv.Atoi(context.Param("id"))
	quote := service.GetQuoteById(int64(quoteId))

	if quote == nil {
		context.String(http.StatusNotFound, "Couldn't find quote with id %d", quoteId)
		return
	}

	context.JSON(http.StatusOK, quote)
}

func DeleteQuote(context *gin.Context) {
	quoteId, _ := strconv.Atoi(context.Param("id"))
	service.DeleteQuoteById(int64(quoteId))

	context.String(http.StatusNoContent, "Deleted")

}

func SaveQuote(context *gin.Context) {
	var newQuote dto.QuoteDto
	err := context.ShouldBindJSON(&newQuote)

	if err != nil {
		context.String(http.StatusBadRequest, "Bad request", err)
	}

	createdQuoteId, err := service.SaveQuote(newQuote)

	if err != nil {
		context.String(http.StatusInternalServerError, "Couldn't save quote", err)
		return
	}

	context.JSON(http.StatusCreated, createdQuoteId)
}
