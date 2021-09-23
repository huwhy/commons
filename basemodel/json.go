package basemodel

type Json struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonOk(message string) *Json {
	return &Json{
		Code:    200,
		Message: message,
		Data:    nil,
	}
}

func JsonData(data interface{}) *Json {
	return &Json{
		Code: 200,
		Data: data,
	}
}

func JsonFail(message string) *Json {
	return &Json{
		Code:    500,
		Message: message,
	}
}

func Json302() *Json {
	return &Json{Code: 302}
}
