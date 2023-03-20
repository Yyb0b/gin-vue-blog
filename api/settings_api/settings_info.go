package settings_api

import (
	"gin_vue_project/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.Ok(map[string]string{}, "test", c)
	res.OkWithData(map[string]string{
		"id": "11",
	}, c)
	res.FailWithCode(res.SettingsError, c)
}
