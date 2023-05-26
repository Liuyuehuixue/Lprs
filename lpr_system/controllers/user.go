package controllers

import (
	"github.com/gin-gonic/gin"
	"lpr_system/logic"
	"lpr_system/models"
 "fmt"
)

// Login 登录权限校验
func Login(c *gin.Context) {
	p := new(models.Login)
	p.WechatId = c.Query("wechat_id")

	//if err != nil {
	//	zap.L().Error("get post detail invalid param", zap.Error(err))
	//	ResponseError(c, CodeInvalidParam)
	//	return
	//}
	//if err := c.ShouldBindJSON(p); err != nil {
	//	// 请求参数有误，返回响应
	//	zap.L().Error("invalid param", zap.Error(err))
	//	// 判断err是不是validator.ValidationErrors类型
	//	errs, ok := err.(validator.ValidationErrors)
	//	if !ok {
	//		ResponseError(c, CodeInvalidParam)
	//		return
	//	}
	//	ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
	//	return
	//}
	count, err := logic.Login(p.WechatId)
	if err != nil || count == 0 {
  fmt.Println("err:",err)
		ResponseError(c, CodeLogin)
		return
	} else {
		ResponseSucess(c, "")
	}

}
