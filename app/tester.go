package app

import (
	"github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"
)

func GetTestResponseWriter() ResponseWriter.ResponseWriter {
	return ResponseWriter.NewResponseWriter()
}
