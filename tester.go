package tester

import (
	"net/http"

	"github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"

	GetRequestModule "github.com/EugeneGpil/httpTester/app/modules/GetRequest"
	RequestModule "github.com/EugeneGpil/httpTester/app/modules/Request"
)

func GetTestResponseWriter() ResponseWriter.ResponseWriter {
	return ResponseWriter.NewResponseWriter()
}

func GetRequest(dto GetRequestDto) http.Request {
	return GetRequestModule.GetRequest(dto)
}

func Request(dto GetRequestDto, mux *http.ServeMux) ResponseWriter.ResponseWriter {
	return RequestModule.Request(dto, mux)
}
