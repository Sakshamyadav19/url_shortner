package model

import "time"

type Url struct{
	Id string	`json:"id"`
	OriginalUrl string `json:"original_url"`
	ShortenUrl string	`json:"shorten_url"`
	Date time.Time		`json:"date"`
}