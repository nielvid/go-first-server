package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nielvid/go-first-server.git/router"
)

type Product struct {
	Code  string `json: "code"`
	Price uint   `json: "price`
}

type Books struct {
	Id       int               `json: "id:`
	Title    string            `json: "title"`
	Author   string            `json: "autthor"`
	Read     bool              `json: "read"`
	Array    []string          `json: "Array"`
	Comments []byte				`json: "comments"`
	// Comments map[string]string `json: "comments"`
}

func main() {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello go server")
	})

	app.Get("/books", func(c *fiber.Ctx) error {
		books := Books{1, "The Guide to golang", "godson", true, []string{"books"}, []byte(`{"shade": "the best book i've read this year"}`)}
		return c.JSON(books)
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		p1 := Product{"fmm233", 30}

		return c.JSON(p1)
	})

	router.MyRoutes(app)
	app.Listen(":3000")

}
