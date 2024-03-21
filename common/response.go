package common

type Response struct {
	Timestamp int64  `json:"timestamp"`
	Success   bool   `json:"success"`
	Data      any    `json:"data"`
	Message   string `json:"message"`
}
