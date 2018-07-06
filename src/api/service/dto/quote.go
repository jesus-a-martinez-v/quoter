package dto

type QuoteDto struct {
	Id     int64  `json:"id"`
	Author string `json:"author" binding:"required"`
	Genre  string `json:"genre" binding:"required"`
	Quote  string `json:"quote" binding:"required"`
}
