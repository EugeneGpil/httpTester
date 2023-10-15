package Response

import "github.com/EugeneGpil/httpTester/app/modules/ResponseWriter"

type Response struct {
	ResponseWriter ResponseWriter.ResponseWriter
}

func (response Response) GetBody() []byte {
	messages := response.ResponseWriter.GetMessages()

	res := make([]byte, 0)

	countOfMessages := len(messages)

	for i := 0; i < countOfMessages; i++ {
		res = addMessageToGetBody(res, messages[i])

		res = addMessagesSeparatorToGetBody(res, i+1, countOfMessages)
	}

	return res
}

func addMessageToGetBody(bodyToReturn []byte, message []byte) []byte {
	for _, symbol := range message {
		bodyToReturn = append(bodyToReturn, symbol)
	}

	return bodyToReturn
}

func addMessagesSeparatorToGetBody(bodyToReturn []byte, messagesCountInBody int, messagesCountInResponseWriter int) []byte {
	if messagesCountInBody == messagesCountInResponseWriter {
		return bodyToReturn
	}

	return append(bodyToReturn, '\n')
}
