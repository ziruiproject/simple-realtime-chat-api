package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simple-realtime-chat-api/models"
)

func main() {
	app := fiber.New()

	// CONFIG
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	// DATABASE
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.GetString("HOST"), config.GetString("DB_USERNAME"), config.GetString("DB_PASSWORD"),
		config.GetString("DB_NAME"), config.GetString("PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	userId := uuid.New()
	db.Create(&models.User{
		ID: userId,
		Email: "yudhalagicoy@gmail.com",
		Name: "yudha",
		Password: "halo",
		Profile: "hello.png",
		Messages: []models.Message{
			{
				ID: uuid.New(),
				UserId: userId,
				Content: "hello world",
			},

		},
	})

	// ROUTES
	app.Get("/", func(ctx *fiber.Ctx) error {
		user := models.User{}
		db.First(&user)
		fmt.Println(user)
		return ctx.SendStatus(200)
	})

	// APP START
	err = app.Listen("localhost:3000")

	if err != nil {
		panic(err)
	}


}