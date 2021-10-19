package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AccesoHistoriaClinica struct {
id_acceso int `orm:"column(id_acceso);pk;auto"`
fecha_acceso time.Time `orm:"column(fecha_acceso);type(date)";not null`
profesional_id int `orm:"column(profesional_id)"`
id_historia_clinica int `orm:"column(id_historia_clinica)"`
}
func (t ¨*Anamnesis) Schema() string {
	return "acceso_historia"
}
func (t *AccesoHistoriaClinica) TableName() string {
	return "AccesoHistoriaClinica"
}

func init() {
	orm.RegisterModel(new(AccesoHistoriaClinica))
}
//AddAccesoHistoriaClinica agrega un acceso a la historia clinica
//Último registro insertado con exito
func AddAccesoHistoriaClinica(m *AccesoHistoriaClinica) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
//GetAccesoHistoriaClinicaById consulta el acceso a la historia clinica por su id
//Id no existe
func GetAccesoHistoriaClinicaById(id int) (v *AccesoHistoriaClinica, err error) {
	o := orm.NewOrm()
	v = &AccesoHistoriaClinica{IdAcceso: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAccesoHistoriaClinica consulta todos los accesos a la historia clinica
//No existen registros
func GetAllAccesoHistoriaClinica(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AccesoHistoriaClinica))
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
	var l []AccesoHistoriaClinica
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
//UpdateAccesoHistoriaClinica actualiza un acceso a la historia clinica
//El regisro a actualizar no existe
func UpdateAccesoHistoriaClinica(m *AccesoHistoriaClinica) (err error) {
	o := orm.NewOrm()
	v := AccesoHistoriaClinica{IdAcceso: m.IdAcceso}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Número de registros actualizados de la base de datos:", num)
		}
	}
	return
}

//DeleteAccesoHistoriaClinica elimina un acceso a la historia clinica
//El registro a eliminar no existe
func DeleteAccesoHistoriaClinica(id int) (err error) {
	o := orm.NewOrm()
	v := AccesoHistoriaClinica{IdAcceso: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AccesoHistoriaClinica{IdAcceso: id}); err == nil {
			fmt.Println("Número de registros eliminados en la base de datos:", num)
		}
	}
	return
}