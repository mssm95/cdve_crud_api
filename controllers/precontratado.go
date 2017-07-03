package controllers

import (
  "cdve_crud_api/models"
  "github.com/astaxie/beego"
)

type PrecontratadoController struct {
	beego.Controller
}

func (c *PrecontratadoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOne", c.GetOne)
}


func (c *PrecontratadoController) GetAll() {
	idResolucion := c.Ctx.Input.Param(":idResolucion")
    listaPrecontratados := models.GetAllPrecontratado(idResolucion)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaPrecontratados
    c.ServeJSON()
}

func (c *PrecontratadoController) GetOne() {
	idResolucion := c.Ctx.Input.Param(":idResolucion")
	idPersona := c.Ctx.Input.Param(":idPersona")
    precontratado := models.GetOnePrecontratado(idResolucion, idPersona)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = precontratado
    c.ServeJSON()
}