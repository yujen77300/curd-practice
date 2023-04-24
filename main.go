package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yujen77300/curd-practice/database"
	"github.com/yujen77300/curd-practice/routes"
	"github.com/yujen77300/curd-practice/utils"
)

func main() {
	var app *fiber.App = fiber.New()
	port := ":3000"

	// register the routes
	routes.SetupRoutes(app)

	// connect to the database
	database.InitDatabase(utils.GetValue("DB_NAME"))

	app.Listen(port)
}
