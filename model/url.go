package model

import (
	"log"
)

type UrlShortner struct {
	Id      int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	LongURL string `json:"long_url" gorm:"not_null"`
}

func AddNewUrl(longURL string) (int64, error) {
	dbClient := GetDB()
	urlShortner := UrlShortner{
		LongURL: longURL,
	}
	result := dbClient.DB.Create(&urlShortner)
	if result.Error != nil {
		log.Fatalf("Some error occured %v", result.Error)
		return -1, result.Error
	}
	return urlShortner.Id, nil
}

func GetLongURL(id int64) (string, error) {
	dbClient := GetDB()
	var urlShortner UrlShortner
	result := dbClient.DB.First(&urlShortner, id)
	if result.Error != nil {
		return "", result.Error
	}
	return urlShortner.LongURL, nil
}
