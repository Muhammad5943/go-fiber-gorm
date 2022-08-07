package migration

import (
	"fmt"
	"log"

	"github.com/Muhammad5943/go-fiber-gorm/database"
	"github.com/Muhammad5943/go-fiber-gorm/models/entity"
)

func RunMigration() {
	/* To drop table */
	// err := database.DB.Migrator().DropTable(&entity.User{})
	/* To Migrate database */
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Migrated")
}
