package tester

import (
	"github.com/EugeneGpil/httpTester/app/types"
)

func GetTestResponseWriter() types.ResponseWriter {
	return types.NewResponseWriter()
}
