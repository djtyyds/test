package api

import (
	"github.com/gin-gonic/gin"
	"test/service"
	"test/tool"
)

func QueryRecord(c *gin.Context) {
	recordTxt := c.Param("txt")
	record, err := service.GetRecordByTxt(recordTxt)
	if err != nil {
		tool.RespInternalError(c)
	}
	tool.RespSuccessfulWithData(c, record)
}
