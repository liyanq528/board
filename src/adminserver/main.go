package main

import (
	"git/inspursoft/board/src/adminserver/dao"
	"git/inspursoft/board/src/adminserver/models"
	"git/inspursoft/board/src/adminserver/service"
	_ "git/inspursoft/board/src/adminserver/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//beego framework configuring and booting.
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	dao.InitDatabase()
	models.RegisterModels()
	dao.InitGlobalCache()
	if err := models.InitInstallationStatus(); err != nil {
		logs.Error(err)
	}
	service.RegisterDBWhenBooting()
	beego.Run()
}
