package tester

type GetRequestDto struct {
	Method string
	Url    string
	Body   interface{}
	Query  map[string]string
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

func (dto GetRequestDto) GetQuery() map[string]string {
	return dto.Query
}
