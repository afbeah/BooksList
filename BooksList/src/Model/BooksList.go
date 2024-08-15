package model

type Book struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Publisher string `json:"publisher"`
	Date string `json:"date"`
	BookNote string `json:"booknote"`
	Review string `json:"review"`
}