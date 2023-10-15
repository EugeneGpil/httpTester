package Response

import (
	"github.com/EugeneGpil/httpTester/app/modules/Response/methods/GetBody"
	"github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"
)

type Response struct {
	writer ResponseWriter.ResponseWriter
}

type NewResponseDto struct {
	Writer ResponseWriter.ResponseWriter
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
