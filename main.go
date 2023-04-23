package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujen77300/curd-practice/routes"
)

func main() {
	var app *fiber.App = fiber.New()
	port := ":3000"

	// register the routes
	routes.SetupRoutes(app)

	app.Listen(port)
}
