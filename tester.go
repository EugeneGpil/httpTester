package tester

import (
	"net/http"

	"github.com/EugeneGpil/httpTester/app/modules/Request"
	"github.com/EugeneGpil/httpTester/app/modules/Request/dto"
	"github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"
)

func GetTestResponseWriter() ResponseWriter.ResponseWriter {
	return ResponseWriter.NewResponseWriter()
}

func GetRequest(dto dto.GetRequestDto) http.Request {
	return Request.GetRequest(dto)
}
