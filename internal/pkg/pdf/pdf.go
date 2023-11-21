package pdf

import (
	"bytes"
	pdfcpuapi "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"io"
)

type PdfProcessor interface {
	ImagetoPdf(img []byte) (output []byte, error error)
}

type pdfprocessor struct {
}

func NewPdfProcessor() PdfProcessor {
	return &pdfprocessor{}
}

func (p pdfprocessor) ImagetoPdf(img []byte) (output []byte, error error) {
	var (
		outBuf = &bytes.Buffer{}
		reader = bytes.NewReader(img) // convert byte slice to io.Reader
		imp    = pdfcpu.DefaultImportConfig()
		rr     = make([]io.Reader, 1)
	)

	rr[0] = reader

	err := pdfcpuapi.ImportImages(nil, outBuf, rr, imp, nil)
	if err != nil {
		return nil, err
	}

	return outBuf.Bytes(), err

}
