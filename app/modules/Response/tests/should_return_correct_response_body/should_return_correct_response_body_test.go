package should_return_correct_response_body

import (
	"testing"

	"github.com/EugeneGpil/httpTester/app/modules/Response"
	"github.com/EugeneGpil/tester"

	responseWriterPackage "github.com/EugeneGpil/responseWriter"
)

var message1 = "message1"
var message2 = "message2"

func Test_should_return_correct_response(t *testing.T) {
	tester.SetTester(t)

	responseWriter := responseWriterPackage.New()

	responseWriter.Write([]byte(message1))
	responseWriter.Write([]byte(message2))

	response := Response.New(Response.NewResponseDto{
		Writer: responseWriter,
	})

	expectedBody := []byte(message1 + "\n" + message2)

	tester.AssertSame(expectedBody, response.GetBody())
}
