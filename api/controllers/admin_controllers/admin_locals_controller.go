package admin_controllers

import (
	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type AdminLocalsController struct {
	controllers.BaseLocalsController
}

// @Title Retrieve Local Info
// @Description Get local info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	int	true		"Local id"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [get]
func (this *AdminLocalsController) Get() {
	id, e := this.GetInt("id")
	this.WE(e, 400)

	o := models.Local{}

	this.BaseLocalsController.Show(id, &o, nil)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Create new local
// @Description Create new local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local		body	models.Local	true		"New Local"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [post]
func (this *AdminLocalsController) Post() {
	o := models.Local{}

	this.BaseLocalsController.Create(&o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Update Local
// @Description Edit local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	int	true		"Local id"
// @Param	local		body	models.Local	true		"Edited Local"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [put]
func (this *AdminLocalsController) Put() {
	id, e := this.GetInt("id")
	this.WE(e, 400)

	o := models.Local{}

	this.BaseLocalsController.Update(id, &o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Delete Local
// @Description Remove local by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	string	true		"Local id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [delete]
func (this *AdminLocalsController) Remove() {
	id, e := this.GetInt("id")
	this.WE(e, 400)
	this.BaseLocalsController.Remove(id)
	this.ServeJSON()
}

// @Title Get Locals List
// @Description Get locals list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	false		"Order Desc"
// @Param	enable_to_reserve		query	string	false		"Local Property (true o false)"
// @Param	area_id		query	int	false		"Local Property"
// @Param	fname		query	string	false		"Search in name"
// @Success 200 {object} []models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /locals [get]
func (this *AdminLocalsController) List() {
	var l []models.Local
	this.BaseLocalsController.List(&l)
	this.Data["json"] = l
	this.ServeJSON()
}