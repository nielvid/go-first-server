package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nielvid/go-first-server.git/handler"
)

func MyRoutes(app *fiber.App) {

	app.Get("/url", handler.GetBooks) 

	app.Get("contact-us", func(c *fiber.Ctx) error {
		type Person struct {
			name string
			age  int
			sex  string
		}
		// john := Person{"john", 23, "male"}
		// toJs, _ := json.Marshal(john)
		var staff = map[string]string{"name": "Paul", "sex": "male", "state": "lagos"}
		return c.JSON(staff)
	})
}
