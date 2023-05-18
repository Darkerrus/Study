package main

import (
	"test_RestApi/database"
	"test_RestApi/routes"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {

	database.Connect()
	engine := html.New("./views", ".html")
	//app := fiber.New()
	app := fiber.New(fiber.Config{
        Views: engine,
    })
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Static("/", "./public")
	routes.Setup(app)
	app.Listen(":3000")
}
