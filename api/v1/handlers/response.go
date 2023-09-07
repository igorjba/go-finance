package handlers

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(msg string) *errorResponse {
	return &errorResponse{Message: msg}
}
