package api

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/uuid"
	"portfolio-api/api/contact"
	"portfolio-api/api/health"
)

// Se cargan los loggerHttp, y los allowedOrigins (registrosHttp) (permisos de origen)
func routes() *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("Juancx web Service")
	prometheus.RegisterAt(app, "/metrics")

	app.Get("/doc/*", swagger.New(swagger.Config{
		URL:         "/doc/doc.json",
		DeepLinking: false,
	}))

	app.Use(recover.New())
	app.Use(prometheus.Middleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization, signature",
		AllowMethods: "GET,POST",
	}))

	app.Use(logger.New())

	TxID := uuid.New().String()

	loadRoutes(app, TxID)

	return app
}

// Aqui se cargan las direcciones o las ubicaciones de las funciones Handler
func loadRoutes(app *fiber.App, TxID string) {
	contact.RouterContact(app, TxID)
	health.RouterHealth(app, TxID)
}
