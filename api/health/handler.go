package health

import (
	"github.com/gofiber/fiber/v2"
	"portfolio-api/internal/models"
)

type handlerHealth struct {
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
	res := models.Response{Error: false, Data: "Sistema conectado"}

	res.Code, res.Msg = 210, "Operación exitosa"
	return c.Status(fiber.StatusOK).JSON(res)
}
