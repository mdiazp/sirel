package controllers

import (
	"errors"

	"github.com/astaxie/beego/context"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

/*
func AccessRolControl(this *beego.Controller, Rol string) {
	beego.Debug("llego a AccessController")
	u, ok := this.Ctx.Input.Data()["User"].(models.User)
	if !ok {
		this.Ctx.Output.SetStatus(http.StatusInternalServerError)
		beego.Error("User data not found in context")
		this.ServeJSON()
		this.StopRun()
	}
	if !u.HaveRol(Rol) {
		this.Ctx.Output.SetStatus(http.StatusUnauthorized)
		this.ServeJSON()
		this.StopRun()
	}
}
*/

func GetAuthorFromInputData(ctx *context.Context) (models.User, error) {
	x := ctx.Input.Data()["Author"]
	if auth, ok := x.(models.User); ok {
		return auth, nil
	}
	return models.User{}, errors.New("Not user founded in ctx.Input.Data[\"Author\"]")
}

type PagAndOrdOptions struct {
	Limit   int
	Offset  int
	OrderBy string
	Desc    bool
}