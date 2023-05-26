package routes

import (
	"github.com/gin-gonic/gin"
	"lpr_system/controllers"
	logger "lpr_system/logger"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置为发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	{
		v1.GET("/user/getOpenId", controllers.GetOpenId)

		v1.GET("/user/login", controllers.Login)

		v1.GET("/children/showAll", controllers.GetAllChildrenHandler)

		v1.GET("/children/suggestion", controllers.GetSuggestionWithtime)

		v1.GET("/children/plot", controllers.PlotWithTime)

		v1.POST("/children/selectItemsByAge/", controllers.SelectItemsByAge)

		v1.GET("/survey/showSurveyById", controllers.ShowSurveyByIdHandler)

		v1.POST("/survey/insertItems/:status", controllers.InsertItemsGenSug)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})

	})

	return r
}
