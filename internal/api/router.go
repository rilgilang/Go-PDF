package api

import (
	"github.com/gofiber/fiber/v2"
	"go_pdf/config/yaml"
	"go_pdf/internal/api/routes"
	"go_pdf/internal/pkg/pdf"
	"go_pdf/internal/service"
)

func NewRouter(cfg *yaml.Config) *fiber.App {
	router := fiber.New()

	//processor
	pdfProcessor := pdf.NewPdfProcessor()

	//service
	var (
		authService = service.NewProcessor(pdfProcessor, cfg)
	)

	//group
	api := router.Group("/api")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("uwow"))
	})

	routes.ProcessorRoutes(api, cfg, authService)

	//routes.HealthRouter(api)

	return router
}
