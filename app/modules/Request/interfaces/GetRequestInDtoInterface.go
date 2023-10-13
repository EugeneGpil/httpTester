package interfaces

type GetRequestDtoInterface interface{
	GetMethod() string
	GetUrl() string
	GetBody() interface{}
}
