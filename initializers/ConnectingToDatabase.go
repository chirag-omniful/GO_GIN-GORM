package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=arjuna.db.elephantsql.com user=wdgruxui password=Z8Mc_XgrFdagUz-F_dwl4B8RPcXK88RZ dbname=wdgruxui port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error in connecting with DB")
	}
}
