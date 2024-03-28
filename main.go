package main

import (
	"log"
	_"fmt"
	controller "tmp/database/controller"
	routes "tmp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func AllRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Post("/createpost", routes.AddPost)
	app.Get("/getpost/:title", routes.GetPostByID)
	app.Get("/delpost/:title", routes.DeletePostByTitle)
}

func main() {

	controller.ConnectToDB()

	app := fiber.New()
	AllRoutes(app)
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))

	


}
