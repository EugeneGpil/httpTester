package Request

import (
	"net/http"

	"github.com/EugeneGpil/httpTester/app/modules/GetRequest"
	"github.com/EugeneGpil/httpTester/app/modules/GetRequest/interfaces"
	"github.com/EugeneGpil/httpTester/app/modules/Response"
	"github.com/EugeneGpil/responseWriter"
)

func Request(dto interfaces.GetRequestDtoInterface, mux *http.ServeMux) Response.Response {
	request := GetRequest.GetRequest(dto)

	handler, _ := mux.Handler(&request)

	writer := responseWriter.New()

	handler.ServeHTTP(writer, &request)

	return Response.New(Response.NewResponseDto{
		Writer: writer,
	})
}
