package migrate

import (
	"Learn/initializers"
	"Learn/models"
	"fmt"
)

func MigrateDB() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		fmt.Println("Error in connecting to Database")
	}
}
