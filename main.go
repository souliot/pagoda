package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/souliot/pagoda/models"
	_ "github.com/souliot/pagoda/routers"
	sutil "github.com/souliot/siot-util"
)

func main() {
	dc := &models.DBConfig{
		Name:   "default",
		Driver: "sqlite3",
		Url:    "data.db?_loc=Local",
	}
	err := models.InitDB(dc)
	if err != nil {
		logs.Error(err)
		return
	}

	if beego.BConfig.RunMode == "dev" {
		orm.RunSyncdb(dc.Name, false, true)
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else {
		beego.BConfig.RecoverFunc = recoverPanic
	}

	err = models.InitData()
	if err != nil {
		logs.Error(err)
		return
	}

	beego.Run()
}

func recoverPanic(ctx *context.Context, cfg *beego.Config) {
	if err := recover(); err != nil {
		if ctx.Output.Status != 0 {
			ctx.ResponseWriter.WriteHeader(ctx.Output.Status)
		} else {
			ctx.Output.JSON(&sutil.Err500, false, false)
		}
	}
}
