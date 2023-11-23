package api

import (
	"github.com/gofiber/fiber/v2"
	"go_pdf/config/yaml"
	"go_pdf/internal/api/routes"
	"go_pdf/internal/pkg/pdf"
	"go_pdf/internal/service"
)

func NewRouter(app *fiber.App, cfg *yaml.Config) *fiber.App {
	router := app

	//processor
	pdfProcessor := pdf.NewPdfProcessor()

	//service
	var (
		authService = service.NewProcessor(pdfProcessor, cfg)
	)

	//group
	api := router.Group("/api")

	routes.ProcessorRoutes(api, cfg, authService)

	//routes.HealthRouter(api)

	return router
}
