package models

type Review struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Author   string `json:"author"`
	Url      string `json:"url"`
	Stars    int    `json:"stars"`
	Message  string `json:"message"`
	Verified bool   `json:"verified"`
}
