package contact

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"portfolio-api/internal/env"
	"portfolio-api/internal/helper"
	"portfolio-api/internal/logger"
	"portfolio-api/internal/models"
	"portfolio-api/services/smtp"
)

type handlerHealth struct {
	db   *sqlx.DB
	txID string
}

// Health godoc
// @Summary Verificar estado del servicio
// @Description Endpoint para verificar el estado y conectividad del servicio
// @Tags Health
// @Accept json
// @Produce json
// @Success 201 {object} models.Response "Sistema conectado"
// @Failure 500 {object} models.Response "Error interno del servidor"
// @Router /health [GET]
func (h *handlerHealth) Health(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := RequestEmailMessage{}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Msg = 1, "Cuerpo de la petición no válida"
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	isValid, err := req.Valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Msg = 2, "Datos enviados no cumplen con los requisitos"
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Msg = 2, "Datos enviados no cumplen con los requisitos"
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	e := env.NewConfiguration()

	html, err := helper.GenerateEmailHTML(req.From)
	if err != nil {
		logger.Error.Printf("couldn't get header image, error: %v", err)
		res.Code, res.Msg = 11, "Error al enviar el mensaje"
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	params := smtp.MapperSimpleEmailMessage(e.SmtpConfig.To, "onboarding@resend.dev", html, "Envío de correo", []string{
		"SMTP_Resend",
	})

	srv := smtp.NewService(e.SmtpConfig.ApiKey)
	response, err := srv.Send(params)
	if err != nil {
		logger.Error.Printf("couldn't login, error: %v", err)
		res.Code, res.Msg = 11, "Error al enviar el mensaje"
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	fmt.Println(response)

	res.Error = false
	res.Code, res.Msg = 210, "Correo enviado correctamente"
	return c.Status(fiber.StatusOK).JSON(res)
}
