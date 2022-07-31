package common

// ErrResp holds the structure for error response
type ErrResp struct {
	Error string `json:"error"`
}

// Resp holds the response structure of success response
type Resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
