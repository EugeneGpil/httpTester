package tester

import (
	"net/http"

	"github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"

	GetRequestModule "github.com/EugeneGpil/httpTester/app/modules/GetRequest"
)

func GetTestResponseWriter() ResponseWriter.ResponseWriter {
	return ResponseWriter.NewResponseWriter()
}

func GetRequest(dto GetRequestDto) http.Request {
	return GetRequestModule.GetRequest(dto)
}
