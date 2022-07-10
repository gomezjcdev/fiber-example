package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
	Id        string
	Firstname string
	Lastname  string
}

// Get Example
func handleGetUser(ctx *fiber.Ctx) error {
	user := User{
		Firstname: "Juan",
		Lastname:  "Gomez",
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(ctx *fiber.Ctx) error {
	user := User{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	user.Id = uuid.NewString()

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userGroup := app.Group("/user")

	userGroup.Get("", handleGetUser)
	userGroup.Post("", handleCreateUser)

	app.Listen(":3000")
}
