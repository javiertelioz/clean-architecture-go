package response

// Response
// @Description response information.
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data" example:"JSON information" swaggerignore:"true"`
}
