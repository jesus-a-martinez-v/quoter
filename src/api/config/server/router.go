package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quoter/src/api/config/constants"
	"quoter/src/api/controller"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()

	Router.GET("/quotes", controller.GetQuotes)
	Router.GET("/quotes/:id", controller.GetQuoteById)
	Router.POST("/quotes", controller.SaveQuote)

	Router.Run(constants.ServerPort)
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	Router.ServeHTTP(w, req)
}
