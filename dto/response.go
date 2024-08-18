package dto

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

func ResponseError(message string, code ...int) Response[string] {
	defaultCode := 99
	if len(code) > 0 {
		defaultCode = code[0]
	}
	return Response[string]{
		Code:    defaultCode,
		Message: message,
		Data:    nil,
	}
}

func ResponseErrorData(message string, data *map[string]string, code ...int) Response[map[string]string] {
	defaultCode := 99
	if len(code) > 0 {
		defaultCode = code[0]
	}

	return Response[map[string]string]{
		Code:    defaultCode,
		Message: message,
		Data:    data,
	}
}

func ResponseSuccess[T any](message string, data *T, code ...int) Response[T] {
	defaultCode := 200

	if len(code) > 0 {
		defaultCode = code[0]
	}
	return Response[T]{
		Code:    defaultCode,
		Message: message,
		Data:    data,
	}
}
