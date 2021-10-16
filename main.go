package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ezzycreative1/majoo-test/database"
	"github.com/ezzycreative1/majoo-test/database/migration"
	routes "github.com/ezzycreative1/majoo-test/routes"

	UserRepo "github.com/ezzycreative1/majoo-test/app/user/repository"
	UserUsecase "github.com/ezzycreative1/majoo-test/app/user/usecase"

	//"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	err error
)

func main() {

	err = godotenv.Load()
	db := database.PostsqlConn()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	dbEvent := os.Getenv("DBEVENT")
	if dbEvent == "rollback_migrate" || dbEvent == "rollback" {
		migration.RunRollback()
	}
	if dbEvent == "migrate_only" {
		migration.RunMigration()
	}
	if dbEvent == "migrate" || dbEvent == "rollback_migrate" {
		migration.RunMigration()
		// Type Seeders
		//seeder.StatusTicketSeeder()
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// CORS
	router.Use(CORSMiddleware())
	//router.Use(cors.Default())

	//Logging
	//router.Use(middle.SetUp())

	// Size Images
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Middleware Token
	//router.Use(middle.AuthMiddleware())

	//USer
	userRepo := UserRepo.NewUserRepository(db)
	userUsecase := UserUsecase.NewUserUsecase(userRepo)

	// Health check
	routes.HealthCheckHTTPHandler(router)
	routes.NewUserHandler(router, userUsecase)

	// User
	//routes.NewUserHandler(router, userUsecase)

	// Server
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))); err != nil {
		log.Fatal(err)
	}
}

// CORSMiddleware ..
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			//c.Next()
			return
		}

		c.Next()
	}
}
