package presenter

import "bytes"

type Response struct {
	Code   int           `json:"code"`
	Data   interface{}   `json:"data,omitempty"`
	Errors error         `json:"errors,omitempty"`
	File   *bytes.Buffer `json:"file,omitempty"`
}

func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) WithFile(file *bytes.Buffer) *Response {
	r.File = file
	return r
}

func (r *Response) WithError(err error) *Response {
	r.Errors = err
	return r
}
