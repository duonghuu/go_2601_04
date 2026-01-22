package response

import "github.com/gin-gonic/gin"

func OK(c *gin.Context, data any) {
	c.JSON(200, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func OKWithMeta(c *gin.Context, data any, meta any) {
	c.JSON(200, SuccessResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func Created(c *gin.Context, data any) {
	c.JSON(201, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func NoContent(c *gin.Context) {
	c.Status(204)
}
