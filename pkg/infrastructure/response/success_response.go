package response

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context, status int, payload interface{}) {
	res := Response{
		Data:    payload,
		Message: "Operation was successful",
	}

	c.JSON(status, res)
}
