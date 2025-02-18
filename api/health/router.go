package health

import (
	"github.com/gofiber/fiber/v2"
)

func RouterHealth(app *fiber.App, txID string) {
	h := handlerHealth{txID: txID}
	v1 := app.Group("/health")
	v1.Get("/", h.Health)

}
