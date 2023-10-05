package response

import "time"

type Error struct {
	ErrorCode string      `json:"errorCode"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details"`
	RequestId string      `json:"request_id"`
	Timestamp time.Time   `json:"timestamp"`
}
