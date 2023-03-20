package routers

import (
	"gin_vue_project/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	apiRouterGroup := router.Group("api")
	//系统配置api
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingsRouter()
	//SettingsRouter(router)
	return router
}
