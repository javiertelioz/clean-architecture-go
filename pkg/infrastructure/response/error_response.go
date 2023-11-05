package response

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, status int, message string) {
	res := Response{
		Message: message,
		Data:    nil,
	}

	c.JSON(status, res)
}
