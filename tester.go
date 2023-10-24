package tester

import (
	"net/http"

	"github.com/EugeneGpil/httpTester/app/modules/Response"
	"github.com/EugeneGpil/responseWriter"

	GetRequestModule "github.com/EugeneGpil/httpTester/app/modules/GetRequest"
	RequestModule "github.com/EugeneGpil/httpTester/app/modules/Request"
)

func GetTestResponseWriter() responseWriter.ResponseWriter {
	return responseWriter.New()
}

func GetRequest(dto GetRequestDto) http.Request {
	return GetRequestModule.GetRequest(dto)
}

func Request(dto GetRequestDto, mux *http.ServeMux) Response.Response {
	return RequestModule.Request(dto, mux)
}
