package errorResponse

type ErrorResponse struct {
	Error ErrorMsg `json:"error"`
}

type ErrorMsg struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

var ErrorStatusInternalServerError = ErrorMsg{StatusCode: 500, Message: "Что-то пошло не так, попробуйте еще раз"}

func NewErrorMsg(statusCode int, message string) ErrorMsg {
	return ErrorMsg{
		StatusCode: statusCode,
		Message:    message,
	}
}
