package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type OrdenadorGasto struct {
	Id              int     `orm:"column(id);pk"`
	Cargo           string  `orm:"column(cargo)"`
	DependenciaId   int     `orm:"column(dependencia_id)"`
}

func (t *OrdenadorGasto) TableName() string {
	return "ordenador_gasto"
}

func init() {
	orm.RegisterModel(new(OrdenadorGasto))
}

// AddOrdenadorGasto insert a new OrdenadorGasto into database and returns
// last inserted Id on success.
func AddOrdenadorGasto(m *OrdenadorGasto) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrdenadorGastoById retrieves OrdenadorGasto by Id. Returns error if
// Id doesn't exist
func GetOrdenadorGastoById(idFacultad string) (v *OrdenadorGasto, err error) {
	o := orm.NewOrm()
	var temp []*OrdenadorGasto
	_, err1 := o.Raw("SELECT * FROM core.ordenador_gasto WHERE core.ordenador_gasto.dependencia_id="+idFacultad+";").QueryRows(&temp)
	if err1 == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp[0], err1
}

// GetAllOrdenadorGasto retrieves all OrdenadorGasto matches certain condition. Returns empty list if
// no records exist
func GetAllOrdenadorGasto()(arregloIDs []OrdenadorGasto){
	o := orm.NewOrm()
	var temp []OrdenadorGasto
	_, err := o.Raw("SELECT * FROM core.ordenador_gasto;").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}

// UpdateOrdenadorGasto updates OrdenadorGasto by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdenadorGastoById(m *OrdenadorGasto) (err error) {
	o := orm.NewOrm()
	v := OrdenadorGasto{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrdenadorGasto deletes OrdenadorGasto by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrdenadorGasto(id int) (err error) {
	o := orm.NewOrm()
	v := OrdenadorGasto{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrdenadorGasto{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
