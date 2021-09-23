package main

import (
	"errors"
	_ "errors"
	"payment-service/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type AccountType int

const (
	Normal AccountType = iota
	Premium
	Gold
)

func GetAccountType(a AccountType) (string, error) {
	if a == 0 {
		return "Normal", nil
	}
	return "data", errors.New("empty string")
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if _, ok := err.(*fiber.Error); ok {
				return errors.New("error in fiber")
			}

			return errors.New("Managed error")
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())

	api.routes(app)

	app.Listen(":3000")
}
