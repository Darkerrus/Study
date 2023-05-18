package routes

import (
	"test_RestApi/controllers"

	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/", controllers.Index)
	app.Get("/create", controllers.Create)
	app.Post("/save_article", controllers.Save)
	app.Get("/animals", controllers.Animals)
	app.Get("/animal/:id", controllers.Animal)
	app.Get("/auth", controllers.Auth)

}
