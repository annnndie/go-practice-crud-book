package main

import (
	"booooooook/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handler.BooksIndex)
	app.Get("/books", handler.BooksIndex)
	app.Get("/book/:id", handler.BookGet)
	app.Post("/books", handler.BookCreate)
	app.Patch("/book/:id", handler.BookUpdate)
	app.Delete("/book/:id", handler.BookDelete)

	app.Listen(":3000")
}
