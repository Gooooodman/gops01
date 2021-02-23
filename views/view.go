package views

import "github.com/gin-gonic/gin"

// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /go [get]
func GoTest(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message":"test go ",
	})
}
