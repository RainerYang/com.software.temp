package api

//Code 组装Code
func Code(code int, message string) map[string]interface{} {
	m := make(map[string]interface{})
	m["code"] = code
	m["message"] = message
	return m
}

var (
	errAPINotDone = Code(999, "接口未完成")
	errOK         = Code(0, "OK")
)
