package GetRequest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/EugeneGpil/httpTester/app/modules/GetRequest/interfaces"
)

func GetRequest(dto interfaces.GetRequestDtoInterface) http.Request {
	urlObj, _ := url.Parse(dto.GetUrl())

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
