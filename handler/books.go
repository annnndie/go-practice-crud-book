package handler

import (
	"booooooook/database"
	"booooooook/model"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "bui",
	})
}

func BooksIndex(c *fiber.Ctx) error {
	db := database.DatabaseInit()

	books := &[]model.Book{}
	result := db.Find(books)

	if result.Error != nil {
		panic("Create failed")
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func BookGet(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DatabaseInit()

	book := &model.Book{}

	db.First(book, id)

	if book.ID == 0 {
		return c.SendString("Record not found")
	}

	return c.JSON(book)
}

func BookCreate(c *fiber.Ctx) error {
	db := database.DatabaseInit()

	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		panic(err)
	}
	result := db.Create(&book)

	if result.Error != nil {
		panic("Create failed")
	}

	if result.RowsAffected != 1 {
		panic("Create failed")
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func BookUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DatabaseInit()

	book := &model.Book{}
	db.First(book, id)

	if book.ID == 0 {
		return c.SendString("Record not found")
	}

	if err := c.BodyParser(book); err != nil {
		db.Save(book)
	}

	return c.JSON(book)
}

func BookDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DatabaseInit()

	book := &model.Book{}

	db.First(book, id)

	if book.ID == 0 {
		return c.SendString("Record not found")
	}

	db.Delete(book)

	return c.SendString("the book deleted")
}
