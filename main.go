package main

import (
	"github.com/fazt/go-fiber-crud/bookmark"
	"github.com/fazt/go-fiber-crud/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	dbErr := database.InitDB()

	if dbErr != nil {
		panic(dbErr)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/bookmarks", bookmark.GetAllBookmarks)
	app.Get("/api/bookmarks/:id", bookmark.GetBookmark)
	app.Post("/api/bookmarks", bookmark.NewBookmark)
	app.Patch("/api/bookmarks/:id", bookmark.UpdateBookmark)
	app.Delete("/api/bookmarks/:id", bookmark.DeleteBookmark)

	app.Listen(":3000")
}
