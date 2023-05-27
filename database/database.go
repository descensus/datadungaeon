package database

import (
	"datadungaeon/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

const appName = "datadungaeon"

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		log.Fatalf("(%s) <-> Unable to connect to DB :(\n", appName)
	}
	log.Printf("(%s) <-> Connected to DB!\n", appName)
}
func Migrate() {
	Instance.AutoMigrate(&models.AqaraPlug{})
	Instance.AutoMigrate(&models.AqaraTemperature{})
	Instance.AutoMigrate(&models.AqaraMagnet{})
	Instance.AutoMigrate(&models.Pocsag{})

	log.Printf("(%s) <-> DB Migrations Done.\n", appName)
}
