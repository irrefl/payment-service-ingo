package fiberConfig

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func get() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if _, ok := err.(*fiber.Error); ok {
				return errors.New("error in fiber")
			}

			return errors.New("Managed error")
		},
	})

	app.Use(logger.New())

	//app.routes(app)

	app.Listen(":3000")
}
