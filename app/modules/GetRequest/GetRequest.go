package GetRequest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/EugeneGpil/httpTester/app/modules/GetRequest/interfaces"
)

func GetRequest(dto interfaces.GetRequestDtoInterface) http.Request {
	query := getQuery(dto.GetQuery())

	urlObj, _ := url.Parse(dto.GetUrl() + "?" + query)

	bodyBytes := getBody(dto.GetBody())

	return http.Request{
		Method: dto.GetMethod(),
		URL:    urlObj,
		Body:   io.NopCloser(bytes.NewReader(bodyBytes)),
	}
}

func getBody(dtoBody interface{}) []byte {
	if (dtoBody == nil) {
		return []byte("{}")
	}

	bodyBytes, _ := json.Marshal(dtoBody)

	return bodyBytes
}

func getQuery(params map[string]string) string {
	keyValue := make([]string, 0)
	for key, value := range params {
		keyValue = append(keyValue, key + "=" + value)
	}

	return strings.Join(keyValue, "&")
}
