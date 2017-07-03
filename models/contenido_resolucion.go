package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Paragrafo struct {
	Id						int
	Numero 					int        
	Texto					string          
}

type Articulo struct {
	Id						int   
	Numero                  int       
	Texto					string 
	Paragrafos				[]Paragrafo         
}

type ResolucionCompleta struct {
	Vinculacion				ResolucionVinculacionDocente
	Consideracion 			string          
	Preambulo 			    string          
	Vigencia                int             
	Numero 			        string             
	Id                      int             
	Articulos				[]Articulo
}

func GetOneResolucionCompleta(idResolucion string) (resolucion ResolucionCompleta) {
	o := orm.NewOrm()
	var temp []Resolucion
	_, err := o.Raw("SELECT * FROM administrativa.resolucion WHERE administrativa.resolucion.id_resolucion="+idResolucion+";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	resolucionCompleta := ResolucionCompleta{Id: temp[0].Id, Consideracion: temp[0].ConsideracionResolucion, Preambulo: temp[0].PreambuloResolucion, Vigencia: temp[0].Vigencia, Numero: temp[0].NumeroResolucion};

	var arts []ComponenteResolucion
	_, err2 := o.Raw("SELECT * FROM administrativa.componente_resolucion WHERE resolucion_id="+idResolucion+" AND tipo_componente like 'Articulo' ORDER BY numero asc;").QueryRows(&arts)
	if err2 == nil {
		fmt.Println("Consulta exitosa")
	}
    
    var articulos []Articulo

	for _, art := range arts {
    	articulo := Articulo{Id: art.Id, Numero: art.Numero, Texto: art.Texto}

    	var pars []ComponenteResolucion
		_, err3 := o.Raw("SELECT * FROM administrativa.componente_resolucion WHERE resolucion_id="+idResolucion+" AND tipo_componente like 'Paragrafo' AND componente_padre="+strconv.Itoa(articulo.Id)+" ORDER BY numero asc;").QueryRows(&pars)
		if err3 == nil {
			fmt.Println("Consulta exitosa")
		}
	    
	    var paragrafos []Paragrafo

	    for _, par := range pars {
	    	paragrafo := Paragrafo{Id: par.Id, Numero: par.Numero, Texto: par.Texto}
	    	paragrafos = append(paragrafos,paragrafo)
	    }

	    articulo.Paragrafos = paragrafos

    	articulos = append(articulos,articulo)
	}
	resolucionCompleta.Articulos=articulos
	return resolucionCompleta
}

func UpdateResolucionCompletaById(m *ResolucionCompleta) (err error) {
	o := orm.NewOrm()
	v := Resolucion{Id: m.Id}
	idResolucionStr:=strconv.Itoa(m.Id)
	r := m.Vinculacion
	fmt.Println(r.Id)
	a := ResolucionVinculacionDocente{Id: r.Id}
	if err = o.Read(&a); err == nil {
		_, err = o.Update(&r)
	}else{
		return
	}
	if err = o.Read(&v); err == nil {
		v.ConsideracionResolucion=m.Consideracion
		v.PreambuloResolucion=m.Preambulo
		v.NumeroResolucion=m.Numero
		fmt.Println(v)
		if err := UpdateResolucionById(&v); err != nil {
		}

		resolucionCompleta:=GetOneResolucionCompleta(idResolucionStr)

		for _, articulo := range resolucionCompleta.Articulos {
			if(articulo.Paragrafos!=nil){
				for _, paragrafo := range articulo.Paragrafos {
					if err := DeleteComponenteResolucion(paragrafo.Id); err != nil {
					}
				}
			}
	    	if err := DeleteComponenteResolucion(articulo.Id); err != nil {
			}
	    }

		for indexArticulo, articulo := range m.Articulos {
			componenteArticulo:=ComponenteResolucion{ResolucionId: &Resolucion{Id: m.Id}, Texto: articulo.Texto, Numero: indexArticulo+1, TipoComponente: "Articulo"}
			if _, err := AddComponenteResolucion(&componenteArticulo); err == nil {
				if(articulo.Paragrafos!=nil){
					for indexParagrafo, paragrafo := range articulo.Paragrafos {
						componenteParagrafo:=ComponenteResolucion{ResolucionId: &Resolucion{Id: m.Id}, Texto: paragrafo.Texto, Numero: indexParagrafo+1, TipoComponente: "Paragrafo", ComponentePadre: &ComponenteResolucion{Id: componenteArticulo.Id}}
						if _, err := AddComponenteResolucion(&componenteParagrafo); err == nil {

						}
					}
				}
			}
	    }
	}
	return
}