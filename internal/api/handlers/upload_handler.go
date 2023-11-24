package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go_pdf/config/yaml"
	"go_pdf/internal/api/presenter"
	"go_pdf/internal/service"
	"io"
)

func UploadHandler(cfg *yaml.Config, service service.Processor) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.Context()
		)

		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("err 1 --> ", err)
			return c.JSON(presenter.ErrorResponse(errors.New("error blok")))
		}

		// Get Buffer from file
		buffer, err := file.Open()

		if err != nil {
			return c.JSON(presenter.ErrorResponse(errors.New("error blok")))
		}
		defer buffer.Close()

		bytesBuff := bytes.NewBuffer(nil)
		if _, err := io.Copy(bytesBuff, buffer); err != nil {
			return c.JSON(presenter.ErrorResponse(errors.New("error blok")))
		}

		serv := service.PdfProcessor(ctx, bytesBuff.Bytes())
		if serv.Code != 200 {
			c.Status(serv.Code)
			return c.JSON(presenter.ErrorResponse(serv.Errors))
		}

		filename := "result.pdf"

		c.Status(200)
		c.Set(fiber.HeaderContentType, fiber.MIMEOctetStream)
		c.Set(fiber.HeaderContentDisposition, fmt.Sprintf(`attachment; filename="%s"`, filename))

		return c.SendStream(serv.File, serv.File.Len())
	}
}
