package utils

type successHttpResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type failHttpResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
	Error  T      `json:"error"`
}

func GenerateSuccessResponse[T any](data T) successHttpResponse[T] {
	return successHttpResponse[T]{
		Status: "OK",
		Data:   data,
	}
}

func GenerateFailedResponse[T any](data T) failHttpResponse[T] {
	return failHttpResponse[T]{
		Status: "KO",
		Error:  data,
	}
}
