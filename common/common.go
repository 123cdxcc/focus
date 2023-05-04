package common

import "github.com/gin-gonic/gin"

func ErrorResp(c *gin.Context, err error) {
	c.JSON(200, Resp[interface{}]{
		Code:    400,
		Message: err.Error(),
	})
}

func SuccessResp(c *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		c.JSON(200, Resp[interface{}]{
			Code:    200,
			Message: "ok",
			Data:    nil,
		})
		return
	}
	c.JSON(200, Resp[interface{}]{
		Code:    200,
		Message: "ok",
		Data:    data[0],
	})
}
