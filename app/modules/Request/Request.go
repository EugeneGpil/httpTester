package Request

import (
	"net/http"

	"github.com/EugeneGpil/httpTester/app/modules/GetRequest"
	"github.com/EugeneGpil/httpTester/app/modules/GetRequest/interfaces"
	"github.com/EugeneGpil/httpTester/app/modules/Response"
	"github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"
)

func Request(dto interfaces.GetRequestDtoInterface, mux *http.ServeMux) Response.Response {
	request := GetRequest.GetRequest(dto)

	handler, _ := mux.Handler(&request)

	writer := ResponseWriter.NewResponseWriter()

	handler.ServeHTTP(writer, &request)

	return Response.Response{
		ResponseWriter: writer,
	}
}
