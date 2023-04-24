package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujen77300/curd-practice/handlers"
	"github.com/yujen77300/curd-practice/middlewares"
)

func SetupRoutes(app *fiber.App) {

	// private routes, authentication is required
	// the middleware is added

	var privateRoutes fiber.Router = app.Group("/api/v1", middlewares.CreateMiddleware())

	privateRoutes.Post("/items", handlers.CreateItem)
	privateRoutes.Put("/items/:id", handlers.UpdateItem)
	privateRoutes.Delete("/items/:id", handlers.DeleteItem)

	var publicRoutes fiber.Router = app.Group("/api/v1")
	publicRoutes.Get("/items", handlers.GetAllItems)
	publicRoutes.Get("/items/:id", handlers.GetItemByID)
	publicRoutes.Post("/signup", handlers.Signup)
	publicRoutes.Post("/login", handlers.Login)
}
