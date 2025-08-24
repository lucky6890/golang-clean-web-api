package helpers

import "github.com/lucky6890/golang-clean-web-api/api/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	ResultCode       int                            `json:"resultCode"`
	Success          bool                           `json:"success"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		ResultCode: resultCode,
		Success:    success,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		ResultCode: resultCode,
		Success:    success,
		Error:      err.Error(),
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		ResultCode:       resultCode,
		Success:          success,
		ValidationErrors: validations.GetValidationErrors(err),
	}
}
