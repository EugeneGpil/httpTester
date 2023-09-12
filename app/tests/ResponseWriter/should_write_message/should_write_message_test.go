package should_write_message

import (
	"testing"

	httpTester "github.com/EugeneGpil/httpTester"
	"github.com/EugeneGpil/tester"
)

func Test_should_write_message(t *testing.T) {
	tester.SetTester(t)
	writer := httpTester.GetTestResponseWriter()

	writer.Write([]byte("Hello"))

	messages := writer.GetMessages()

	tester.AssertLen(messages, 1)

	tester.AssertSame(messages[0], []byte("Hello"))
}
