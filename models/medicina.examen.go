package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type MecicinaExamen struct {
	IdExamen       int                   `orm:"column(id_examen);pk;auto"`
	IdHojaHistoria *MedicinaHojaHistoria `orm:"column(id_hoja_historia);rel(fk);null"`
	Nombre         string                `orm:"column(nombre);null"`
	Observacion    string                `orm:"column(observacion);null"`
	IdTipoExamen   *MedicinaTipoExamen   `orm:"column(id_tipo_examen);rel(fk);null"`
	FechaExamen    time.Time             `orm:"column(fecha_examen);type(date);null"`
}

func (t *MecicinaExamen) TableName() string {
	return "examen"
}
func init() {
	orm.RegisterModel(new(MecicinaExamen))
}

// AddMecicinaExamen inserta un registro en la tabla examen
// Último registro insertado con éxito
func AddMecicinaExamen(m *MecicinaExamen) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	return id, err
}

// GetMecicinaExamenById obtiene un registro de la tabla examen por su id
// Id no existe
func GetMecicinaExamenById(id int) (MecicinaExamen, error) {
	o := orm.NewOrm()
	m := MecicinaExamen{IdExamen: id}
	err := o.Read(&m)
	return m, err
}

// GetAllMecicinaExamen obtiene todos los registros de la tabla examen
// No existen registros
func GetAllMecicinaExamen(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MecicinaExamen))
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
					return nil, errors.New("error: Orden inválido, debe ser del tipo [asc|desc]")
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
					return nil, errors.New("error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: los tamaños de 'sortby', 'order' no coinciden o el tamaño de 'order' no es 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: campos de 'order' no utilizados")
		}
	}
	var l []MecicinaExamen
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

// UpdateMecicinaExamen actualiza un registro de la tabla examen
// El registro a actualizar no existe
func UpdateMecicinaExamen(m *MecicinaExamen) (err error) {
	o := orm.NewOrm()
	v := MecicinaExamen{IdExamen: m.IdExamen}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteMecicinaExamen elimina un registro de la tabla examen
// El registro a eliminar no existe
func DeleteMecicinaExamen(id int) (err error) {
	o := orm.NewOrm()
	v := MecicinaExamen{IdExamen: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MecicinaExamen{IdExamen: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
