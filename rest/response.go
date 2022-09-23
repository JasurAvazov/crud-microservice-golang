package rest

type JSONResponse struct {
	ErrorNote string      `json:"error_note"`
	ErrorCode int         `json:"error_code"`
	Status    string      `json:"status"`
	Data      interface{} `json:"data"`
}
