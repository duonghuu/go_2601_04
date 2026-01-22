package response

import "github.com/gin-gonic/gin"

func Error(
	c *gin.Context,
	httpStatus int,
	code string,
	message string,
	details any,
) {
	c.JSON(httpStatus, ErrorResponse{
		Success: false,
		Error: ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}
