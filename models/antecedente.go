package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)
type Antecedente struct {
	id_antecedente int `orm:"column(id_antecedente);pk"`
	id_tipo_antecedente int `orm:"column(id_tipo_antecedente);null"`
	observaciones	string `orm:"column(observaciones);null"`
	id_historia_clinica int `orm:"column(id_historia_clinica);null"`
}
func (t *Antecedente) Schema() string {
	return "medicina"
}
func (t *Antecedente) TableName() string {
	return "antecedente"
}
func init() {
	orm.RegisterModel(new(Antecedente))
}
// AddAntecedente insert a new Usuario into database and returns
// Último registro insertado con éxito
func AddAntecendete(m *Antecedente) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
//GetAntecedenteById obtiene un antecedente por su id
//Id no existe
func GetAntecedenteById(id int) (v *Antecedente, err error) {
	o := orm.NewOrm()
	v = &Antecedente{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
//GetAllAntecedente obtiene todos los antecedentes
//No existen registros
func GetAllAntecedente() (query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
		o := orm.NewOrm()
		qs := o.QueryTable(new(Antecedente))
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
		var l []Antecedente
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
//UpdateAntecedente actualiza un antecedente
//El regustro a actualizar no existe
func UpdateAntecedente (m *Antecedente) (err error) {
	o := orm.NewOrm()
	v := Antecedente{IdAntecedente: m.IdAntecedente}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Registros actualizados:", num)
		}
	}
	return
}
//DeleteAntecedente elimina un antecedente
//El registro a eliminar no existe
func DeleteAntecedente(id int) (err error) {
	o := orm.NewOrm()
	v := Antecedente{IdAntecedente: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Antecedente{IdAntecedente: id}); err == nil {
			fmt.Println("Número de registros eliminados de la base de datos:", num)
		}
	}
	return
}
