package helper

func SetResponse(code int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
}
