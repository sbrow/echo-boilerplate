package app

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *Database = &Database{}

func init() {
	DB.Connect()
	DB.Migrate()
}

func (d *Database) Connect() {
	db, err := gorm.Open("sqlite3", "test.sqlite")
	if err != nil {
		panic("failed to connect database")
	}
	DB.DB = db
}

func (d *Database) Migrate() {
	// Migrate the schema
	d.AutoMigrate(User{})
	log.Println("Migrated")
}
