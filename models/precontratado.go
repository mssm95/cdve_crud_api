package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Precontratado struct {
    Id int `orm:"column(id);pk"`
    PrimerNombre string `orm:"column(primer_nombre)"`
    SegundoNombre string `orm:"column(segundo_nombre)"`
    PrimerApellido string `orm:"column(primer_apellido)"`
    SegundoApellido string `orm:"column(segundo_apellido)"`
    Documento int `orm:"column(documento)"`
    Expedicion string `orm:"column(expedicion)"`
    Categoria string `orm:"column(categoria)"`
    Dedicacion string `orm:"column(dedicacion)"`
    HorasSemanales int `orm:"column(horas_semanales)"`
    Semanas int `orm:"column(semanas)"`
    ProyectoCurricular int `orm:"column(proyecto_curricular)"`
}

func init() {
	orm.RegisterModel(new(Precontratado))
}

func GetAllPrecontratado(idResolucion string) (arregloIDs []Precontratado) {
	o := orm.NewOrm()
	var temp []Precontratado
	_, err := o.Raw("SELECT administrativa.vinculacion_docente.id id, agora.informacion_persona_natural.primer_nombre primer_nombre, agora.informacion_persona_natural.segundo_nombre segundo_nombre, agora.informacion_persona_natural.primer_apellido primer_apellido, agora.informacion_persona_natural.segundo_apellido segundo_apellido, agora.informacion_persona_natural.num_documento_persona documento, core.ciudad.nombre expedicion, administrativa.escalafon.nombre_escalafon categoria, administrativa.dedicacion.nombre_dedicacion dedicacion, administrativa.vinculacion_docente.numero_horas_semanales horas_semanales, administrativa.vinculacion_docente.numero_semanas semanas, administrativa.vinculacion_docente.id_proyecto_curricular proyecto_curricular FROM agora.informacion_persona_natural, core.ciudad, administrativa.escalafon_persona, administrativa.escalafon, administrativa.dedicacion, administrativa.vinculacion_docente WHERE agora.informacion_persona_natural.id_ciudad_expedicion_documento=core.ciudad.id_ciudad AND agora.informacion_persona_natural.num_documento_persona=administrativa.vinculacion_docente.id_persona AND administrativa.escalafon_persona.id_persona_natural=agora.informacion_persona_natural.num_documento_persona AND administrativa.vinculacion_docente.id_dedicacion=administrativa.dedicacion.id_dedicacion AND administrativa.escalafon_persona.id_escalafon=administrativa.escalafon.id_escalafon AND administrativa.vinculacion_docente.estado=TRUE AND administrativa.vinculacion_docente.id_resolucion = "+idResolucion+";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}

func GetOnePrecontratado(idResolucion string, idPersona string) (precontratado *Precontratado) {
    o := orm.NewOrm()
    var temp []*Precontratado
    _, err := o.Raw("SELECT administrativa.vinculacion_docente.id id, agora.informacion_persona_natural.primer_nombre primer_nombre, agora.informacion_persona_natural.segundo_nombre segundo_nombre, agora.informacion_persona_natural.primer_apellido primer_apellido, agora.informacion_persona_natural.segundo_apellido segundo_apellido, agora.informacion_persona_natural.num_documento_persona documento, core.ciudad.nombre expedicion, administrativa.escalafon.nombre_escalafon categoria, administrativa.dedicacion.nombre_dedicacion dedicacion, administrativa.vinculacion_docente.numero_horas_semanales horas_semanales, administrativa.vinculacion_docente.numero_semanas semanas, administrativa.vinculacion_docente.id_proyecto_curricular proyecto_curricular FROM agora.informacion_persona_natural, core.ciudad, administrativa.escalafon_persona, administrativa.escalafon, administrativa.dedicacion, administrativa.vinculacion_docente WHERE agora.informacion_persona_natural.id_ciudad_expedicion_documento=core.ciudad.id_ciudad AND agora.informacion_persona_natural.num_documento_persona=administrativa.vinculacion_docente.id_persona AND administrativa.escalafon_persona.id_persona_natural=agora.informacion_persona_natural.num_documento_persona AND administrativa.vinculacion_docente.id_dedicacion=administrativa.dedicacion.id_dedicacion AND administrativa.escalafon_persona.id_escalafon=administrativa.escalafon.id_escalafon AND administrativa.vinculacion_docente.estado=TRUE AND administrativa.vinculacion_docente.id_resolucion = "+idResolucion+" AND agora.informacion_persona_natural.num_documento_persona="+idPersona+";").QueryRows(&temp)
    if err == nil {
        fmt.Println("Consulta exitosa")
    }

    if(temp!=nil && len(temp)>0){
        return temp[0]
    }else{
        return nil
    }
}