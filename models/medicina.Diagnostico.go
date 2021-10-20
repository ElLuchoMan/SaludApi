package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"time"

	"github.com/astaxie/beego/orm"
)

type DiagnosticoMedicina struct {
	IdDiagnosticoMedicina int       `orm:"column(id_diagnostico_medicina);pk;auto"`
	Nombre                string    `orm:"column(nombre);null"`
	Descripcion           string    `orm:"column(descripcion);null"`
	Numero                int       `orm:"column(numero);null"`
	Activo                bool      `orm:"column(activo);null"`
	Fecha_creacion        time.Time `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	Fecha_modificacion    time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	Plan_manejo           string    `orm:"column(plan_manejo);null"`
	Analisis              string    `orm:"column(analisis);null"`
}

func (t *DiagnosticoMedicina) TableName() string {
	return "diagnosticomedicina"
}
func init() {
	orm.RegisterModel(new(DiagnosticoMedicina))
}

//AddDiagnosticoMedicina inserta un registro en la tabla Diagnostico
//Último registro insertado con éxito
func AddDiagnosticoMedicina(m *DiagnosticoMedicina) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//GetDiagnosticoMedicinaById obtiene un registro de la tabla Diagnostico mediante su id
//Id no existe
func GetDiagnosticoMedicinaById(id int) (v *DiagnosticoMedicina, err error) {
	o := orm.NewOrm()
	v = &DiagnosticoMedicina{IdDiagnosticoMedicina: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetAllDiagnosticoMedicina obtiene todos los registros de la tabla Diagnostico
//No existen registros
func GetAllDiagnosticoMedicina(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DiagnosticoMedicina))
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
	var l []DiagnosticoMedicina
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

//UpdateDiagnosticoMedicina actualiza un registro en la tabla Diagnostico
//El registro a actualizar no existe
func UpdateDiagnosticoMedicina(m *DiagnosticoMedicina) (err error) {
	o := orm.NewOrm()
	v := DiagnosticoMedicina{IdDiagnosticoMedicina: m.IdDiagnosticoMedicina}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

//DeleteDiagnosticoMedicina elimina un registro de la tabla Diagnostico
//El registro a eliminar no existe
func DeleteDiagnosticoMedicina(id int) (err error) {
	o := orm.NewOrm()
	v := DiagnosticoMedicina{IdDiagnosticoMedicina: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DiagnosticoMedicina{IdDiagnosticoMedicina: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
