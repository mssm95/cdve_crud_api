package controllers

import (
	"cdve_crud_api/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

// OrdenadorGastoController oprations for OrdenadorGasto
type OrdenadorGastoController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrdenadorGastoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create OrdenadorGasto
// @Param	body		body 	models.OrdenadorGasto	true		"body for OrdenadorGasto content"
// @Success 201 {int} models.OrdenadorGasto
// @Failure 403 body is empty
// @router / [post]
func (c *OrdenadorGastoController) Post() {
	var v models.OrdenadorGasto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddOrdenadorGasto(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get OrdenadorGasto by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.OrdenadorGasto
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrdenadorGastoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	v, err := models.GetOrdenadorGastoById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get OrdenadorGasto
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.OrdenadorGasto
// @Failure 403
// @router / [get]
func (c *OrdenadorGastoController) GetAll() {
    listaOrdenadoresGasto := models.GetAllOrdenadorGasto()
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaOrdenadoresGasto
    c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the OrdenadorGasto
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.OrdenadorGasto	true		"body for OrdenadorGasto content"
// @Success 200 {object} models.OrdenadorGasto
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrdenadorGastoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.OrdenadorGasto{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateOrdenadorGastoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the OrdenadorGasto
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrdenadorGastoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOrdenadorGasto(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
