package main

import (
	"belajargo/controls"
	"belajargo/repos"
	"belajargo/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// ROUTE
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading godotenv")
	}

	// APP
	appPort := os.Getenv("APP_PORT")
	// DB
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// DATABASE
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Koneksi database error")
	}

	// REPOSITORY
	prizeRepos := repos.PrizeRepository(db)

	// SERVICE
	prizeService := services.PrizeService(prizeRepos)

	// HANDLER
	prizeHandle := controls.PrizeController(prizeService)

	//   ROUTES
	v1 := router.Group("/v1")
	v1.GET("/", prizeHandle.PrizeIndex)
	v1.POST("/prize", prizeHandle.PostPrize)
	v1.GET("/prize/:id", prizeHandle.PrizeById)
	v1.PUT("/prize/:id", prizeHandle.UpdatePrize)
	v1.DELETE("/prize/:id", prizeHandle.DeleteById)

	//
	router.Run(appPort)
}
