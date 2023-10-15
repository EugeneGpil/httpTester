package Request

import (
	"net/http"

	"github.com/EugeneGpil/httpTester/app/modules/GetRequest"
	"github.com/EugeneGpil/httpTester/app/modules/GetRequest/interfaces"
	"github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"
)

func Request(dto interfaces.GetRequestDtoInterface, mux *http.ServeMux) ResponseWriter.ResponseWriter {
	request := GetRequest.GetRequest(dto)

	handler, _ := mux.Handler(&request)

	writer := ResponseWriter.NewResponseWriter()

	handler.ServeHTTP(writer, &request)

	return writer
}
