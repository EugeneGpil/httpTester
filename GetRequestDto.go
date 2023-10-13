package tester

type GetRequestDto struct {
	Method string
	Url    string
	Body   interface{}
}

func (dto GetRequestDto) GetMethod() string {
	return dto.Method
}

func (dto GetRequestDto) GetUrl() string {
	return dto.Url
}

func (dto GetRequestDto) GetBody() interface{} {
	return dto.Body
}
