package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MedicinaAntecedente struct {
	IdMedicinaAntecedente int    `orm:"column(id_antecedente);pk;auto"`
	id_tipo_antecedente   int    `orm:"column(id_tipo_antecedente);null"`
	observaciones         string `orm:"column(observaciones);null"`
	id_historia_clinica   int    `orm:"column(id_historia_clinica);null"`
}

func (t *MedicinaAntecedente) Schema() string {
	return "medicina"
}
func (t *MedicinaAntecedente) TableName() string {
	return "antecedente"
}
func init() {
	orm.RegisterModel(new(MedicinaAntecedente))
}

// AddAntecedente insert a new Usuario into database and returns
// Último registro insertado con éxito
func AddAntecendete(m *MedicinaAntecedente) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//GetAntecedenteById obtiene un antecedente por su id
//Id no existe
func GetAntecedenteById(id int) (v *MedicinaAntecedente, err error) {
	o := orm.NewOrm()
	v = &MedicinaAntecedente{IdMedicinaAntecedente: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetAllAntecedente obtiene todos los antecedentes
//No existen registros
func GetAllAntecedente(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MedicinaAntecedente))
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
	var l []MedicinaAntecedente
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
func UpdateAntecedente(m *MedicinaAntecedente) (err error) {
	o := orm.NewOrm()
	v := MedicinaAntecedente{IdMedicinaAntecedente: m.IdMedicinaAntecedente}
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
	v := MedicinaAntecedente{IdMedicinaAntecedente: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MedicinaAntecedente{IdMedicinaAntecedente: id}); err == nil {
			fmt.Println("Número de registros eliminados de la base de datos:", num)
		}
	}
	return
}
