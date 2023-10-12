package should_return_request_with_correct_url_and_method

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/httpTester/app/modules/Request"
	"github.com/EugeneGpil/httpTester/app/modules/Request/dto"
	"github.com/EugeneGpil/tester"
)

var method string = http.MethodPost
var url string = "/hello/world"

func Test_should_return_request_with_correct_url_and_method(t *testing.T) {
	tester.SetTester(t)

	request := Request.GetRequest(dto.GetRequestDto{
		Method: method,
		Url:    url,
	})

	tester.AssertSame(request.Method, method)
	tester.AssertSame(request.URL.Path, url)
}
