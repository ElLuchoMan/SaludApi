package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type OdontologiaControlPlaca struct {
	IdcontrolPlaca      int       `orm:"column(idcontrol_placa);pk;auto"`
	Indice_anterior     int       `orm:"column(indice_anterior);null"`
	Indice_actual       int       `orm:"column(indice_actual);null"`
	Fecha               time.Time `orm:"column(fecha);type(date);null"`
	Id_hoja_historia    int       `orm:"column(id_hoja_historia);null"`
	Vestibulares        string    `orm:"column(vestibulares);null"`
	Observaciones       string    `orm:"column(observaciones);null"`
	Id_tipo_odontograma int       `orm:"column(id_tipo_odontograma);null"`
}

func (t *OdontologiaControlPlaca) TableName() string {
	return "controlplaca"
}
func init() {
	orm.RegisterModel(new(OdontologiaControlPlaca))
}

//AddControl plata inserta un registro en la tabla ControlPlacca
//Último registro insertado con éxito
func AddControl(m *OdontologiaControlPlaca) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//GetControlPlacaById consulta un registro en la tabla ControlPlaca por su id
//Id no existe
func GetControlPlacaById(id int) (v *OdontologiaControlPlaca, err error) {
	o := orm.NewOrm()
	v = &OdontologiaControlPlaca{IdcontrolPlaca: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetAllControlPlaca consulta todos los registros en la tabla ControlPlaca
//No existen registros
func GetAllControlPlaca(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OdontologiaControlPlaca))
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
	var l []OdontologiaControlPlaca
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

//UpdateControlPlaca actualiza un registro en la tabla ControlPlaca
//El registro a actualizar no existe
func UpdateControlPlaca(m *OdontologiaControlPlaca) (err error) {
	o := orm.NewOrm()
	v := OdontologiaControlPlaca{IdcontrolPlaca: m.IdcontrolPlaca}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Número de registros actualizados de la base de datos:", num)
		}
	}
	return
}

//DeleteControlPlaca elimina un registro en la tabla ControlPlaca
//El registro a eliminar no existe
func DeleteControlPlaca(id int) (err error) {
	o := orm.NewOrm()
	v := OdontologiaControlPlaca{IdcontrolPlaca: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OdontologiaControlPlaca{IdcontrolPlaca: id}); err == nil {
			fmt.Println("Número de registros eliminados de la base de datos:", num)
		}
	}
	return
}
