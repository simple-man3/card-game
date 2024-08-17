package responses

type (
	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	ErrorResponse struct {
		FailedField string `json:"field"`
		Tag         string `json:"validation"`
		Value       any    `json:"value"`
	}
)
