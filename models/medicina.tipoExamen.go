package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MedicinaTipoExamen struct {
	IdTipoExamen int    `orm:"column(id_tipo_examen);pk;auto"`
	Nombre       string `orm:"column(nombre);null"`
}

func (p *MedicinaTipoExamen) TableName() string {
	return "tipoexamen"
}
func init() {
	orm.RegisterModel(new(MedicinaTipoExamen))
}

//AddMedicinaTipoExameninserta un registro en la tabla diagnostico_psicologia
//Último registro insertado con éxito
func AddMedicinaTipoExamen(m *MedicinaTipoExamen) (err error) {
	o := orm.NewOrm()
	m.IdTipoExamen = 0
	if _, err = o.Insert(m); err == nil {
		return nil
	}
	return err
}

//GetMedicinaTipoExamenById trae un registro de la tabla diagnostico_psicologia
//Id no existe
func GetMedicinaTipoExamenById(id int) (v *MedicinaTipoExamen, err error) {
	o := orm.NewOrm()
	v = &MedicinaTipoExamen{IdTipoExamen: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetAllMedicinaTipoExamen trae todos los registros de la tabla diagnostico_psicologia
//No existen registros
func GetAllMedicinaTipoExamen(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MedicinaTipoExamen))
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
	var l []MedicinaTipoExamen
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

//UpdateMedicinaTipoExamen actualiza un registro en la tabla diagnostico_psicologia
//El registro a actualizar no existe
func UpdateMedicinaTipoExamen(m *MedicinaTipoExamen) (err error) {
	o := orm.NewOrm()
	v := MedicinaTipoExamen{IdTipoExamen: m.IdTipoExamen}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

//DeleteMedicinaTipoExamen elimina un registro de la tabla diagnostico_psicologia
//El registro a eliminar no existe
func DeleteMedicinaTipoExamen(id int) (err error) {
	o := orm.NewOrm()
	v := MedicinaTipoExamen{IdTipoExamen: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MedicinaTipoExamen{IdTipoExamen: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}