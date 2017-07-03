package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Contratado struct {
    Id string `orm:"column(id);pk"`
    Nombre string `orm:"column(nombre)"`
    Documento string `orm:"column(documento)"`
    Expedicion string `orm:"column(expedicion)"`
    Categoria string `orm:"column(categoria)"`
    Dedicacion string `orm:"column(dedicacion)"`
    HorasSemanales int `orm:"column(horas_semanales)"`
    PeriodoVinculacion int `orm:"column(periodo_vinculacion)"`
    ValorContrato float64 `orm:"column(valor_contrato)"`
    IdProyectoCurricular int `orm:"column(id_proyecto_curricular)"`
}

func init() {
	orm.RegisterModel(new(Contratado))
}

func GetAllContratado(tipoDedicacion string) (arregloIDs []Contratado) {
	o := orm.NewOrm()
	var temp []Contratado
	_, err := o.Raw("SELECT agora.informacion_persona_natural.num_documento_persona id, (agora.informacion_persona_natural.primer_nombre || ' ' || agora.informacion_persona_natural.segundo_nombre || ' ' || agora.informacion_persona_natural.primer_apellido || ' ' || agora.informacion_persona_natural.segundo_apellido) nombre, agora.informacion_persona_natural.num_documento_persona documento, core.ciudad.nombre expedicion, administrativa.escalafon.nombre_escalafon categoria, administrativa.dedicacion.nombre_dedicacion dedicacion, administrativa.contrato_docente.numero_horas_semanales horas_semanales, administrativa.contrato_docente.numero_semanas periodo_vinculacion, argo.contrato_general.valor_contrato valor_contrato, administrativa.contrato_docente.id_proyecto_curricular id_proyecto_curricular FROM agora.informacion_persona_natural, core.ciudad, administrativa.escalafon_persona, administrativa.escalafon, administrativa.dedicacion, administrativa.contrato_docente, argo.contrato_general WHERE agora.informacion_persona_natural.id_ciudad_expedicion_documento=core.ciudad.id_ciudad AND agora.informacion_persona_natural.num_documento_persona=administrativa.contrato_docente.id_persona AND administrativa.contrato_docente.numero_contrato=argo.contrato_general.numero_contrato AND administrativa.contrato_docente.vigencia=argo.contrato_general.vigencia AND administrativa.escalafon_persona.id_persona_natural=agora.informacion_persona_natural.num_documento_persona AND administrativa.contrato_docente.id_dedicacion=administrativa.dedicacion.id_dedicacion AND administrativa.escalafon_persona.id_escalafon=administrativa.escalafon.id_escalafon AND administrativa.dedicacion.nombre_dedicacion like '"+tipoDedicacion+"';").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}