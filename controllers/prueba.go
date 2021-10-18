package controllers

import (
	"encoding/json"

	"github.com/ElLuchoMan/SaludApi/models"

	"github.com/astaxie/beego"
)

//PruebaController ...
type PruebaController struct {
	beego.Controller
}

// URLMapping ...
func (c *PruebaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @Description create Prueba
// @Param	body		body 	models.Prueba	true		"body for Prueba content"
// @Success 201 {int} models.Prueba
// @Failure 403 body is empty
// @router / [post]
func (c *PruebaController) Post() {
	var v models.Prueba
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddPrueba(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
}

// GetAll ...
// @Title Get All
// @Description get Prueba
// @Success 200 {object} models.Prueba
// @Failure 403
// @router / [get]
func (c *PruebaController) GetAll() {

	l, err := models.GetAllPrueba()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

//GetA ...
//@Title GetA
//@Description GetA
// @Param	nombre		path 	string	true		"The key for staticblock"
//@Success 200 {object} models.Prueba
//@Failure 403
//@router /GetA [get]
func (c *PruebaController) GetA() {
	nombre := c.Ctx.Input.Param(":nombre")
	v, err := models.GetPruebaByNombre(nombre)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
