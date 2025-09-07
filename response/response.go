package response

type FailedStruct struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SuccessStruct struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}
