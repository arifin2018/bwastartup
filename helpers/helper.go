package helpers

type response struct {
	Meta meta
	Data interface{}
}

type meta struct {
	Message string
	Code    int
	Status  string
}

func ApiResponse(message string, code int, status string, data interface{}) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
