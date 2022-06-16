package main

import (
	"yemeksepeti/database"
	"yemeksepeti/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	//database ile bağlantımızı oluşturuyoruz.
	database.Connect()

	//fiber app oluşturuyoruz.
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	//routes paketinden setup fonksiyoununda yolluyoruz.
	routes.Setup(app)

	//8080 portunu dinliyoruz.
	app.Listen(":8080")
}
