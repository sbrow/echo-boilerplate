package app

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName,
	LastName string `validate:"alpha|required"`
}
