package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

type ResolucionVinculacion struct {
    Id int `orm:"column(id);pk"`
    Estado bool `orm:"column(estado)"`
    Numero string `orm:"column(numero)"`
    Vigencia int `orm:"column(vigencia)"`
    Facultad string `orm:"column(facultad)"`
    NivelAcademico string `orm:"column(nivel_academico)"`
    Dedicacion string `orm:"column(dedicacion)"`
    FechaExpedicion time.Time `orm:"column(fecha_expedicion);type(date);null"`
}

func init() {
	orm.RegisterModel(new(ResolucionVinculacion))
}

func GetAllResolucionVinculacion() (arregloIDs []ResolucionVinculacion) {
	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	_, err := o.Raw("SELECT r.id_resolucion id, r.estado estado, r.numero_resolucion numero, r.vigencia vigencia, d.nombre facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, oikos.dependencia d WHERE rv.id_facultad=d.id AND r.id_resolucion=rv.id_resolucion ORDER BY id desc;").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}