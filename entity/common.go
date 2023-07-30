package entity

type CommonResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
