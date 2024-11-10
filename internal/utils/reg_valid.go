package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidFullUrl(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	re := regexp.MustCompile(`https?://[^\s/$.?#].[^\s]*\.[a-z]{2,}$`)

	return re.MatchString(url)
}

func ValidShortenUrl(fl validator.FieldLevel) bool {
	shortUrl := fl.Field().String()
	re := regexp.MustCompile(

		// TODO: Кол-во символов перенести в конфиг
		fmt.Sprintf("^[a-zA-Z0-9]{%d}$", 5),
	)

	return re.MatchString(shortUrl)
}
