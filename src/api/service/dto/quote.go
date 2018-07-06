package dto

type QuoteDto struct {
	Id     int64  `json:"id"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Quote  string `json:"quote"`
}
