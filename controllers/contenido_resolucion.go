package controllers

import (
  	"cdve_crud_api/models"
  	"encoding/json"
	"strconv"
	"fmt"

  	"github.com/astaxie/beego"
)

type ResolucionCompletaController struct {
	beego.Controller
}

func (c *ResolucionCompletaController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)

}


func (c *ResolucionCompletaController) GetOne() {
	idResolucion := c.Ctx.Input.Param(":idResolucion")
    resolucion := models.GetOneResolucionCompleta(idResolucion)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = resolucion
    c.ServeJSON()
}

func (c *ResolucionCompletaController) Put() {
	idResolucionStr := c.Ctx.Input.Param(":idResolucion")
	idResolucion, _ := strconv.Atoi(idResolucionStr)
	v := models.ResolucionCompleta{Id: idResolucion}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateResolucionCompletaById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err.Error())
		c.Data["json"] = err.Error()
	}
    c.ServeJSON()
}