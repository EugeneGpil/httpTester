package should_return_request_with_correct_query_params

import (
	"testing"

	"github.com/EugeneGpil/httpTester/app/modules/GetRequest"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var key1 = "test_key_1"
var expectedValue1 = "test_value_1"
var key2 = "test_key_2"
var expectedValue2 = "test_value_2"

func Test_should_return_request_with_correct_query_params(t *testing.T) {
	tester.SetTester(t)

	queryParams := map[string]string{
		key1: expectedValue1,
		key2: expectedValue2,
	}

	request := GetRequest.GetRequest(httpTester.GetRequestDto{
		Method: "test",
		Url:    "test",
		Query:  queryParams,
	})

	resultQuery := request.URL.Query()

	resultValue1 := resultQuery.Get(key1)
	resultValue2 := resultQuery.Get(key2)

	tester.AssertSame(expectedValue1, resultValue1)
	tester.AssertSame(expectedValue2, resultValue2)

	tester.AssertLen(resultQuery, 2)
}
