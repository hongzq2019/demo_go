package config

type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// SuccessMsg 操作成功
func SuccessMsg(msg string) Result {
	return Result{
		Success: true,
		Code:    200,
		Data:    nil,
		Message: msg,
	}
}

// SimpleSuccess 操作成功
func SimpleSuccess() Result {
	return Result{
		Success: true,
		Code:    200,
		Data:    nil,
		Message: "操作成功!",
	}
}

// SuccessData 操作成功,包括返回数据
func SuccessData(msg string, data interface{}) Result {
	return Result{
		Success: true,
		Code:    200,
		Data:    data,
		Message: msg,
	}
}

// Fail 操作失败
func Fail(msg string, code int) Result {
	return Result{
		Success: false,
		Code:    code,
		Data:    nil,
		Message: msg,
	}
}

// SimpleFail 操作失败
func SimpleFail() Result {
	return Result{
		Success: false,
		Code:    500,
		Data:    nil,
		Message: "操作失败!",
	}
}
