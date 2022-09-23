package views

// R - ...
type R struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorNote string      `json:"error_note"`
	Data      interface{} `json:"data"`
}

// Languages ...
type Languages struct {
	Uz string `json:"uz"`
	Ru string `json:"ru"`
	En string `json:"en"`
}

// FileUploadResponse ...
type FileUploadResponse struct {
	Filename string `json:"filename"`
}
