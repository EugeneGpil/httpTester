package should_request

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var mux = http.NewServeMux()
var method = http.MethodGet
var url = "/test"
var message = "test message"

func Test_should_request(t *testing.T) {
	tester.SetTester(t)

	router.AddRoute(method, url, func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(message))
	}, "test")

	router.DefineRoutes(mux)

	response := httpTester.Request(httpTester.GetRequestDto{
		Method: method,
		Url:    url,
	}, mux)

	tester.AssertSame([]byte(message), response.GetBody())
}
