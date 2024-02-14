package main

import (
	"belajar-gorm/handler"
	"belajar-gorm/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Issekai99!!@tcp(127.0.0.1:3306)/bwa-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	router := gin.Default()

	apiUser := router.Group("/api/user")
	{
		userRepository := user.NewRepository(db)

		userService := user.NewService(userRepository)

		userHandler := handler.NewUserHandler(userService)

		apiUser.POST("/register", userHandler.RegisterUser)

		apiUser.POST("/login", userHandler.LoginUser)
	}

	router.Run()
}
