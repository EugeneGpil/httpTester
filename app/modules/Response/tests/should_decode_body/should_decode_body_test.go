package should_decode_body

import (
	"testing"

	"github.com/EugeneGpil/httpTester/app/modules/Response"
	"github.com/EugeneGpil/tester"

	responseWriterPackage "github.com/EugeneGpil/responseWriter"
)

func Test_should_decode_body(t *testing.T) {
	tester.SetTester(t)

	responseWriter := responseWriterPackage.New()

	responseWriter.Write([]byte("{\"test\":\"test\"}"))

	response := Response.New(Response.NewResponseDto{
		Writer: responseWriter,
	})

	var decodedBody = struct{
		Test string
	} {}

	err := response.DecodeBody(&decodedBody)
	tester.AssertNil(err)

	tester.AssertSame(decodedBody.Test, "test")
}