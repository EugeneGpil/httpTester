package should_write_response_status

import (
	"net/http"
	"testing"

	httpTester "github.com/EugeneGpil/httpTester"
	"github.com/EugeneGpil/tester"
)

func Test_should_write_response_status(t *testing.T) {
	tester.SetTester(t)

	writer := httpTester.GetTestResponseWriter()

	writer.WriteHeader(http.StatusOK)

	tester.AssertSame(http.StatusOK, writer.GetStatus())
}
