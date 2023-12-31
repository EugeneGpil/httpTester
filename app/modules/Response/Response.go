package Response

import (
	"encoding/json"

	"github.com/EugeneGpil/httpTester/app/modules/Response/methods/GetBody"
	"github.com/EugeneGpil/responseWriter"
)

type Response struct {
	writer responseWriter.ResponseWriter
}

type NewResponseDto struct {
	Writer responseWriter.ResponseWriter
}

func New(dto NewResponseDto) Response {
	return Response{
		writer: dto.Writer,
	}
}

func (response Response) GetBody() []byte {
	return GetBody.GetBody(response.writer)
}

func (response Response) GetStatus() int {
	return response.writer.GetStatus()
}

func (response Response) DecodeBody(body interface{}) error {
	bodyRaw := response.GetBody()

	return json.Unmarshal(bodyRaw, body)
}
