package ResponseWriter

import "net/http"

type ResponseWriter struct{}

var messages [][]byte
var status int

func NewResponseWriter() ResponseWriter {
	status = 200
	messages = [][]byte{}

	return ResponseWriter{}
}

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) Header() http.Header {
	return http.Header{}
}

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) Write(message []byte) (int, error) {
	messages = append(messages, message)

	return 0, nil
}

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) WriteHeader(statusCode int) {
	status = statusCode
}

func (writer *ResponseWriter) GetMessages() [][]byte {
	return messages
}

func (writer *ResponseWriter) GetStatus() int {
	return status
}
