package dto

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

func ResponseError(message string) Response[string] {
	return Response[string]{
		Code:    99,
		Message: message,
		Data:    nil,
	}
}

func ResponseErrorData(message string, data *map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Code:    99,
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
