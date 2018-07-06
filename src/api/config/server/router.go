package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quoter/src/api/controller"
)

const port = ":8000"

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	Router.GET("/quotes", controller.GetQuotes)
	Router.GET("/quotes/:id", controller.GetQuoteById)
	Router.Run(port)
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	Router.ServeHTTP(w, req)
}
