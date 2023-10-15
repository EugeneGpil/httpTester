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

	writer := httpTester.Request(httpTester.GetRequestDto{
		Method: method,
		Url:    url,
	}, mux)

	tester.AssertLen(writer.GetMessages(), 1)
	tester.AssertSame(writer.GetMessages()[0], []byte(message))
}
