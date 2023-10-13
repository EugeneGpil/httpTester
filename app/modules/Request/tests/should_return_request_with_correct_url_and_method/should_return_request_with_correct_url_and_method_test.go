package should_return_request_with_correct_url_and_method

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var method string = http.MethodPost
var url string = "/hello/world"

func Test_should_return_request_with_correct_url_and_method(t *testing.T) {
	tester.SetTester(t)

	request := httpTester.GetRequest(httpTester.GetRequestDto{
		Method: method,
		Url:    url,
	})

	tester.AssertSame(request.Method, method)
	tester.AssertSame(request.URL.Path, url)
}
