package model

// ListOperationsByTypeResponseModel ...

type ResponseModel struct {
	Message string      `json:"message" `
	Error   bool        `json:"error"`
	Data    interface{} `json:"data" `
}
