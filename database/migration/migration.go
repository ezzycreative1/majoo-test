package migration

import (
	"fmt"
	"log"

	"github.com/ezzycreative1/majoo-test/models"

	"github.com/ezzycreative1/majoo-test/database"
)

// RunRollback ..
func RunRollback() {

	db := database.PostsqlConn()

	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}

	if exist := db.HasTable("users"); exist {
		fmt.Println("drop table users")
		err := db.DropTable("users")
		if err == nil {
			fmt.Println("success drop table users")
		}
	}
}

// RunMigration ..
func RunMigration() {
	db := database.PostsqlConn()

	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}

	if exist := db.HasTable("users"); !exist {
		fmt.Println("migrate table users")
		err := db.CreateTable(&models.User{})
		if err == nil {
			fmt.Println("success migrate table users")
		}
	}
}
