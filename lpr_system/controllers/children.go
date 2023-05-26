package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"lpr_system/logic"
	"lpr_system/models"
	"strconv"
)

// GetAllChildrenHandler 显示所有登记过的孩子
func GetAllChildrenHandler(c *gin.Context) {

	data, err := logic.GetALLChildren()
	if err != nil {
		zap.L().Error("logic.GetALLChildren failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSucess(c, data)

}

// GetSuggestionWithtime 显示孩子的针对某次测试结果给的康复训练意见
func GetSuggestionWithtime(c *gin.Context) {

	// 1.从URL获取孩子ID以及对应测试轮次
	p := new(models.IDTimeInfo1)
	// 1.从URL获取帖子ID
	p.Time = c.Query("test_n")

	cid, err := strconv.Atoi(c.Query("id"))
	p.ChildID = cid
	if err != nil {
		zap.L().Error("get post detail invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetSuggestionWithTime(p.ChildID, p.Time)
	if err != nil {
		zap.L().Error("logic.GetSuggestionWithtime", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSucess(c, data)

}

func PlotWithTime(c *gin.Context) {

	// 1.获取参数参数校验
	p := new(models.IDTimeInfo)
	cid, err := strconv.Atoi(c.Query("child_id"))

	p.ChildID = cid
	if err != nil {
		zap.L().Error("get invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//fmt.Println(p.ChildID)
	data, err := logic.PlotWithTime(p.ChildID)
	if err != nil {
		zap.L().Error("logic.PlotWithtime", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSucess(c, data)
}
