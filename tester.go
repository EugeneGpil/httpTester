package tester

import (
	"github.com/EugeneGpil/httpTester/app"
	"github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"
)

func GetTestResponseWriter() ResponseWriter.ResponseWriter {
	return app.GetTestResponseWriter()
}
