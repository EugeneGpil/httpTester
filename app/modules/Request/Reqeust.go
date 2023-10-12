package Request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/EugeneGpil/httpTester/app/modules/Request/dto"
)

func GetRequest(dto dto.GetRequestDto) http.Request {
	urlObj, _ := url.Parse(dto.Url)

	bodyBytes, _ := json.Marshal(dto.Body)

	return http.Request{
		Method: dto.Method,
		URL:    urlObj,
		Body:   io.NopCloser(bytes.NewReader(bodyBytes)),
	}
}
