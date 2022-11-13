package model

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Id      int64  `json:"id,omitempty"`
}
