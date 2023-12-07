package should_return_request_with_correct_request_body

import (
	"testing"

	"github.com/EugeneGpil/httpTester/app/modules/GetRequest"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

func Test_should_return_request_with_correct_request_body(t *testing.T) {
	tester.SetTester(t)

	requestBody := struct {
		Test string
	}{
		Test: "test",
	}

	request := GetRequest.GetRequest(httpTester.GetRequestDto{
		Method: "test",
		Url:    "test",
		Body:   requestBody,
	})

	resultRequestBody := make([]byte, 15)

	request.Body.Read(resultRequestBody)

	tester.AssertSame("{\"Test\":\"test\"}", string(resultRequestBody))
}
