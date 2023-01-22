package main

import (
	"fmt"
	"go-express/internal/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
func helloFromMiddleware(c *fiber.Ctx) error {
	fmt.Println("hello from middleware")
	return c.Next()
}
func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use("/api", helloFromMiddleware)
	app.Get("/ping", ping)
	app.Get("/api/products", handlers.GetAllProducts)
	app.Post("/api/products", handlers.CreateProduct)

	log.Fatal(app.Listen(":80"))
}

