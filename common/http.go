package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetQueryInt(c *gin.Context, key string) (int, error) {
	keyStr, ok := c.GetQuery(key)
	if !ok {
		return 0, errors.New("没有" + key + "参数")
	}
	keyInt, err := strconv.Atoi(keyStr)
	if err != nil {
		return 0, errors.New(key + "不是int")
	}
	return keyInt, nil
}
