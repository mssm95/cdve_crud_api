package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type ContratoGeneral struct {
	Actividades                  *string              `orm:"column(actividades);null"`
	LugarEjecucion               *LugarEjecucion     `orm:"column(lugar_ejecucion);rel(fk)"`
	TipoContrato                 *TipoContrato       `orm:"column(tipo_contrato);rel(fk)"`
	ResgistroPresupuestal        *int                 `orm:"column(resgistro_presupuestal);null"`
	Estado                       *bool                `orm:"column(estado);null"`
	NumeroConstancia             *int                 `orm:"column(numero_constancia);null"`
	Convenio                     *string              `orm:"column(convenio);null"`
	ClaseContratista             int                 `orm:"column(clase_contratista)"`
	Supervisor                   *SupervisorContrato `orm:"column(supervisor);rel(fk);null"`
	Observaciones                *string              `orm:"column(observaciones);null"`
	TipoControl                  *int                 `orm:"column(tipo_control);null"`
	ValorTasaCambio              *float64             `orm:"column(valor_tasa_cambio);null"`
	ValorContratoMe              *float64             `orm:"column(valor_contrato_me);null"`
	TipoMoneda                   int                 `orm:"column(tipo_moneda)"`
	OrigenRecursos               int                 `orm:"column(origen_recursos)"`
	OrigenPresupueso             int                 `orm:"column(origen_presupueso)"`
	TemaGastoInversion           int                 `orm:"column(tema_gasto_inversion)"`
	TipoGasto                    int                 `orm:"column(tipo_gasto)"`
	RegimenContratacion          int                 `orm:"column(regimen_contratacion)"`
	Procedimiento                int                 `orm:"column(procedimiento)"`
	ModalidadSeleccion           int                 `orm:"column(modalidad_seleccion)"`
	TipoCompromiso               int                 `orm:"column(tipo_compromiso)"`
	TipologiaContrato            int                 `orm:"column(tipologia_contrato)"`
	FechaRegistro                time.Time           `orm:"column(fecha_registro);type(date)"`
	UnidadEjecutora              *UnidadEjecutora    `orm:"column(unidad_ejecutora);rel(fk)"`
	Condiciones                  string              `orm:"column(condiciones)"`
	DescripcionFormaPago         string              `orm:"column(descripcion_forma_pago)"`
	Justificacion                string              `orm:"column(justificacion)"`
	ValorContrato                float64             `orm:"column(valor_contrato)"`
	UnidadEjecucion              *Parametros         `orm:"column(unidad_ejecucion);rel(fk)"`
	Contratista                  float64             `orm:"column(contratista)"`
	NumeroCdp                    int                 `orm:"column(numero_cdp)"`
	NumeroSolicitudNecesidad     int                 `orm:"column(numero_solicitud_necesidad)"`
	DependenciaSolicitante       string              `orm:"column(dependencia_solicitante);null"`
	SedeSolicitante              string              `orm:"column(sede_solicitante);null"`
	ClausulaRegistroPresupuestal *bool                `orm:"column(clausula_registro_presupuestal);null"`
	OrdenadorGasto               string              `orm:"column(ordenador_gasto);null"`
	FormaPago                    *Parametros         `orm:"column(forma_pago);rel(fk)"`
	PlazoEjecucion               int                 `orm:"column(plazo_ejecucion)"`
	ObjetoContrato               string              `orm:"column(objeto_contrato);null"`
	Vigencia                     int                 `orm:"column(vigencia)"`
	Id                           int                 `orm:"column(numero_contrato);pk;auto"`
}

type ContratoVinculacion struct {
	ContratoGeneral 			ContratoGeneral
	VinculacionDocente			VinculacionDocente
}

func (t *ContratoGeneral) TableName() string {
	return "contrato_general"
}

func init() {
	orm.RegisterModel(new(ContratoGeneral))
}

func AddContratosVinculcionEspecial(m []ContratoVinculacion)(err error){
	o := orm.NewOrm()
	o.Begin()
	for _, vinculacion := range m {
		v:=vinculacion.VinculacionDocente
		if err = o.Read(&v); err == nil {
			if(v.NumeroContrato==nil && v.Vigencia==nil){
				contrato := vinculacion.ContratoGeneral
				aux1:=181
				contrato.Vigencia, _, _=time.Now().Date()
			    contrato.FormaPago=&Parametros{Id:240}
			    contrato.DescripcionFormaPago="Abono a Cuenta Mensual de acuerdo a puntas y hotras laboradas"
			    contrato.Justificacion="Docente de Vinculacion Especial"
			    contrato.UnidadEjecucion=&Parametros{Id:205}
			    contrato.LugarEjecucion=&LugarEjecucion{Id:2}
			    contrato.TipoControl=&aux1
			    contrato.ClaseContratista=33
			    contrato.TipoMoneda=137
			    contrato.OrigenRecursos=149
			    contrato.OrigenPresupueso=156
			    contrato.TemaGastoInversion=166
			    contrato.TipoGasto=146
			    contrato.RegimenContratacion=136
			    contrato.Procedimiento=132
			    contrato.ModalidadSeleccion=123
			    contrato.TipoCompromiso=35
			    contrato.TipologiaContrato=46
			    contrato.FechaRegistro=time.Now()
			    contrato.UnidadEjecutora=&UnidadEjecutora{Id:1}
			    contrato.Condiciones="Sin condiciones"
			    _, err = o.Insert(&contrato)
			    if (err == nil){
			    	aux1 := strconv.Itoa(contrato.Id)
			    	aux2 := contrato.Vigencia
			    	e:=ContratoEstado{}
			    	e.NumeroContrato = aux1
			    	e.Vigencia = aux2
			    	e.FechaRegistro=time.Now()
			    	e.Estado=1
			    	_, err = o.Insert(&e)
			    	if (err == nil){
				    	a := vinculacion.VinculacionDocente
						if err = o.Read(&a); err == nil {
							v := a
							v.NumeroContrato = &aux1
							v.Vigencia = &aux2
							_, err = o.Update(&v)
							if(err != nil){
								o.Rollback()
								return
							}
						}else{
							o.Rollback()
							return
						}
					}else{
						o.Rollback()
						return
					}
			    }else{
			    	fmt.Println(err.Error())
			    	o.Rollback()
			    	return
			    }
			}else{
				aux1 := *v.NumeroContrato
				aux2 := *v.Vigencia
				e:=ContratoEstado{}
				e.NumeroContrato = aux1
				e.Vigencia = aux2
				e.FechaRegistro=time.Now()
				e.Estado=1
				_, err = o.Insert(&e)
				if (err != nil){
					o.Rollback()
					return
				}
			}
		}else{
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

// AddContratoGeneral insert a new ContratoGeneral into database and returns
// last inserted Id on success.
func AddContratoGeneral(m *ContratoGeneral) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetContratoGeneralById retrieves ContratoGeneral by Id. Returns error if
// Id doesn't exist
func GetContratoGeneralById(id int) (v *ContratoGeneral, err error) {
	o := orm.NewOrm()
	v = &ContratoGeneral{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContratoGeneral retrieves all ContratoGeneral matches certain condition. Returns empty list if
// no records exist
func GetAllContratoGeneral(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ContratoGeneral))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ContratoGeneral
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateContratoGeneral updates ContratoGeneral by Id and returns error if
// the record to be updated doesn't exist
func UpdateContratoGeneralById(m *ContratoGeneral) (err error) {
	o := orm.NewOrm()
	v := ContratoGeneral{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContratoGeneral deletes ContratoGeneral by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContratoGeneral(id int) (err error) {
	o := orm.NewOrm()
	v := ContratoGeneral{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ContratoGeneral{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
