package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=restapigo password=lintang12345 sslmode=disable")
	if err != nil {
		return nil, err
	}
	return DB, nil
}
