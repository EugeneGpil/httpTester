package should_return_request_with_empty_body

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

func Test_should_return_request_with_empty_body(t *testing.T) {
	tester.SetTester(t)

	request := httpTester.GetRequest(httpTester.GetRequestDto{
		Url: "/test",
		Method: http.MethodPost,
	})

	resultRequestBody := make([]byte, 2)

	request.Body.Read(resultRequestBody)

	tester.AssertSame("{}", string(resultRequestBody))
}
