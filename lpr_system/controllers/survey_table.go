package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"lpr_system/dao/mysql"
	"lpr_system/logic"
	"lpr_system/models"
	"reflect"
	"strconv"
)

// SelectItemsByAge 添加孩子信息，并根据孩子年龄返回范围内的数据
func SelectItemsByAge(c *gin.Context) {
	// 1.获取参数参数校验
	p := new(models.ChildInfoTable)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，返回响应
		zap.L().Error("Input with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.业务处理
	// 2.1.检查children表，看是否需要更新
	err := logic.UpdateChildInfo(p)

	// 2.2 返回符合年龄的各表测试项目

	data, err := logic.GetALLSelectedSurveyItems(p.ID, p.Age, p.TestN)

	// 如果data为空，返回错误，不能提交
	if reflect.DeepEqual(data.AllItems, reflect.Zero(reflect.TypeOf(data.AllItems)).Interface()) {
		ResponseError(c, CodeEdit)
	} else {
		ResponseSucess(c, data)
	}

	if err != nil {
		zap.L().Error("logic.GetALLSelectedSurveyItems failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

}

//// ShowSurveyByIdHandler 显示某孩子某调查表的某次测试得分
//func ShowSurveyByIdHandler(c *gin.Context) {
//	// 1.获取参数参数校验
//
//	p := new(models.SurveyChildID)
//	// 1.从URL获取帖子ID
//	p.TestN = c.Query("test_n")
//
//	cid, err := strconv.Atoi(c.Query("child_id"))
//	sid, err := strconv.Atoi(c.Query("survey_id"))
//	p.ChildID = cid
//	p.SurveyID = sid
//	if err != nil {
//		zap.L().Error("get post detail invalid param", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	data, err := logic.GetItemsAllTimeById(p.ChildID, p.SurveyID, p.TestN)
//	if err != nil {
//		zap.L().Error("logic.GetItemsAllTimeById failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	if reflect.DeepEqual(data.AllTimesItems, reflect.Zero(reflect.TypeOf(data.AllTimesItems)).Interface()) {
//		ResponseError(c, CodeInvalidParam)
//	} else {
//		ResponseSucess(c, data)
//	}
//
//}

// ShowSurveyByIdHandler 显示某孩子某调查表的某次测试得分
func ShowSurveyByIdHandler(c *gin.Context) {
	// 1.获取参数参数校验

	cid, err := strconv.Atoi(c.Query("child_id"))
	test_n := c.Query("test_n")

	if err != nil {
		zap.L().Error("get post detail invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 1.获取参数参数校验
	p := new(models.ChildInfoTable)
	p.ID = cid
	p.TestN = test_n
	// 查找该次对应年龄
	p.Age, err = mysql.GetAgeByTime(p.ID, p.TestN)

	data, err := logic.GetALLSelectedSurveyItems2(p.ID, p.Age, p.TestN)

	// 如果data为空，返回错误，不能提交
	if reflect.DeepEqual(data.AllItems, reflect.Zero(reflect.TypeOf(data.AllItems)).Interface()) {
		ResponseError(c, CodeEdit)
	} else {
		ResponseSucess(c, data)
	}

	if err != nil {
		zap.L().Error("logic.GetALLSelectedSurveyItems failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
}

// InsertItemsGenSug 向info表插入数据，根据项目得分给出建议
func InsertItemsGenSug(c *gin.Context) {
	// 1.获取参数参数校验
	statusStr := c.Param("status")
	status, err := strconv.ParseInt(statusStr, 10, 64)
	p := new(models.SelectedSevenTableItems)

	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，返回响应
		zap.L().Error("Input with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	err = logic.InsertItemsGenSug(p, int(status))
	if err != nil {
		zap.L().Error("logic.InsertItemsGenSug failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSucess(c, "添加成功！")
}
