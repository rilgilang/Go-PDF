package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_pdf/config/yaml"
	"go_pdf/internal/api/handlers"
	"go_pdf/internal/service"
)

func ProcessorRoutes(app fiber.Router, cfg *yaml.Config, service service.Processor) {
	app.Post("/upload", handlers.UploadHandler(cfg, service))
}
