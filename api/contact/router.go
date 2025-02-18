package contact

import (
	"github.com/gofiber/fiber/v2"
)

func RouterContact(app *fiber.App, txID string) {
	h := handlerHealth{txID: txID}
	v1 := app.Group("/v1")
	contact := v1.Group("/contact")
	contact.Post("/", h.ContactHandler)
}
