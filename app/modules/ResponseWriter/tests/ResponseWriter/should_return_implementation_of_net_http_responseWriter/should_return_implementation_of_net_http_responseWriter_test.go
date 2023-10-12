package should_return_implementation_of_net_http_responseWriter

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/httpTester"
)

func Test_should_return_implementation_of_net_http_responseWriter(t *testing.T) {
	getResponseWriter()
}

func getResponseWriter() http.ResponseWriter {
	return tester.GetTestResponseWriter()
}
