package dto

type GetRequestDto struct {
	Method string
	Url    string
	Body   interface{}
}
