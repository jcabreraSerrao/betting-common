package dto

type CustomResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	TexShow      string      `json:"texShow"`
	ValidateCode string      `json:"validateCode"`
	Data         interface{} `json:"data"`
}

type CustomError struct {
	Message      string `json:"message"`
	Code         int    `json:"code"`
	ValidateCode string `json:"validateCode"`
	TexShow      string ``
}

func (e *CustomError) Error() string {
	return e.Message
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

type Response struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}
