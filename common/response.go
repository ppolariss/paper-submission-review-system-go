package common

import "time"

type Response struct {
	Timestamp int64  `json:"timestamp"`
	Success   bool   `json:"success"`
	Data      any    `json:"data"`
	Message   string `json:"message"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NoDataSuccessfulResponse(message string) *BaseResponse {
	return &BaseResponse{
		Success: true,
		Message: message,
	}
}

func SuccessfulResponse(message string, data any) *Response {
	return &Response{
		Timestamp: time.Now().Unix(),
		Success:   true,
		Data:      data,
		Message:   message,
	}
}
