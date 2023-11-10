package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
//
//	firstUserId := uuid.New()
//	secondUserId := uuid.New()
//
//	db.Create(&models.User{
//		ID: firstUserId,
//		Email: "yudha@gmail.com",
//		Name: "yudha",
//		Password: "rahasia",
//		Profile: "hello.png",
//	})
//
//	db.Create(&models.User{
//		ID: secondUserId,
//		Email: "harto@gmail.com",
//		Name: "sugiharto",
//		Password: "rahasia",
//		Profile: "hello.png",
//	})
//
//	firstUser := models.User{}
//	db.First(&firstUser, "id = ?", firstUserId)
//
//	secondUser := models.User{}
//	db.First(&secondUser, "id = ?", secondUserId)
//
//	db.Create(&models.Message{
//		ID: uuid.New(),
//		Content: "Halo",
//		SenderID: firstUserId,
//		ReceiverID: secondUserId,
//	})
//
//	fmt.Println(firstUser)
//	fmt.Println(secondUser)

	// MIDDLEWARE

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// ROUTES

	app.Get("/ws/:user", websocket.New(func(c *websocket.Conn) {

		senderId, err := uuid.Parse(c.Params("sender"))
		if err != nil{
			panic("gagal")
		}

		var (
			mt  int
			msg []byte
		)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("Read:", err)
				break
			}
			log.Printf("Diterima: %s", msg)

			db.Create(&models.Message{
				ID: uuid.New(),
				Content: string(msg),
				SenderID: senderId,
				ReceiverID: receiverId,
			})

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("Write:", err)
				break
			}
		}
	}))

	// RETRIVE ALL USERS
	app.Get("/", func(ctx *fiber.Ctx) error {
		var users []models.User
		db.Find(&users)

		jsonUser, err := json.Marshal(users)
		if err != nil {
			return ctx.SendStatus(500)
		}

		return ctx.Send(jsonUser)
	})

	// RETRIVE USER BY ID
	app.Get("/users/:id", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("id")

		user := models.User{}
		db.First(&user, "id = ?", userId)

		userJson, err := json.Marshal(user)

		if err != nil {
			return ctx.SendStatus(500)
		}

		return ctx.Send(userJson)
	})

	//	RETRIVE MESSAGE
	app.Get("/users/:id/:receiver", func(ctx *fiber.Ctx) error {
		senderId := ctx.Params("id")
		receiverId := ctx.Params("receiver")

		var messages []models.Message

		db.Preload("Sender").Preload("Receiver").Where("sender_id = ? and receiver_id = ?", senderId, receiverId).Find(&messages)

		jsonMessage, err := json.Marshal(messages)

		if err != nil {
			return ctx.SendStatus(500)
		}

		return ctx.Send(jsonMessage)
	})

	// APP START
	err = app.Listen("localhost:3000")

	if err != nil {
		panic(err)
	}


}