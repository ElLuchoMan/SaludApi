package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type MedicinaTipoAntecedente struct {
	IdTipoAntecedente int       `orm:"column(id_tipo_antecedente);pk;auto"`
	Nombre            string    `orm:"column(nombre);null"`
	Descripcion       string    `orm:"column(descripcion);null"`
	Activo            bool      `orm:"column(activo);null"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(date);null"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(date);null"`
}

func (p *MedicinaTipoAntecedente) TableName() string {
	return "tipoantecedente"
}
func init() {
	orm.RegisterModel(new(PsicologiaDiagnostico))
}

//AddMedicinaTipoAntecedente inserta un registro en la tabla diagnostico_psicologia
//Último registro insertado con éxito
func AddMedicinaTipoAntecedente(m *MedicinaTipoAntecedente) (err error) {
	o := orm.NewOrm()
	m.IdTipoAntecedente = 0
	if _, err = o.Insert(m); err == nil {
		return nil
	}
	return err
}

//GetMedicinaTipoAntecedenteById trae un registro de la tabla diagnostico_psicologia
//Id no existe
func GetMedicinaTipoAntecedenteById(id int) (v *MedicinaTipoAntecedente, err error) {
	o := orm.NewOrm()
	v = &MedicinaTipoAntecedente{IdTipoAntecedente: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetAllMedicinaTipoAntecedenteo trae todos los registros de la tabla diagnostico_psicologia
//No existen registros
func GetAllMedicinaTipoAntecedente(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MedicinaTipoAntecedente))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: los tamaños de 'sortby', 'order' no coinciden o el tamaño de 'order' no es 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: campos de 'order' no utilizados")
		}
	}
	var l []MedicinaTipoAntecedente
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
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

//UpdateMedicinaTipoAntecedente actualiza un registro en la tabla diagnostico_psicologia
//El registro a actualizar no existe
func UpdateMedicinaTipoAntecedente(m *MedicinaTipoAntecedente) (err error) {
	o := orm.NewOrm()
	v := MedicinaTipoAntecedente{IdTipoAntecedente: m.IdTipoAntecedente}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

//DeleteMedicinaTipoAntecedente elimina un registro de la tabla diagnostico_psicologia
//El registro a eliminar no existe
func DeleteMedicinaTipoAntecedente(id int) (err error) {
	o := orm.NewOrm()
	v := MedicinaTipoAntecedente{IdTipoAntecedente: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MedicinaTipoAntecedente{IdTipoAntecedente: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
