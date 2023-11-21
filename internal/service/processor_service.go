package service

import (
	"bytes"
	"context"
	"fmt"
	"go_pdf/config/yaml"
	"go_pdf/internal/api/presenter"
	"go_pdf/internal/pkg/pdf"
)

type Processor interface {
	PdfProcessor(ctx context.Context, imgByte []byte) *presenter.Response
}

type processor struct {
	cfg *yaml.Config
	pdf pdf.PdfProcessor
}

func NewProcessor(pdf pdf.PdfProcessor, cfg *yaml.Config) Processor {
	return &processor{
		cfg: cfg,
		pdf: pdf,
	}
}

func (s *processor) PdfProcessor(ctx context.Context, imgByte []byte) *presenter.Response {
	var (
		response = presenter.Response{}
		//log      = logger.NewLog("pdf_processor", s.cfg.Logger.Enable)
	)

	imagetoPdf, err := s.pdf.ImagetoPdf(imgByte)
	if err != nil {
		response.WithCode(500).WithData("icikiwir")
	}

	fmt.Println("image to pdf --> ", imagetoPdf)

	return response.WithCode(200).WithFile(bytes.NewBuffer(imagetoPdf))
}
