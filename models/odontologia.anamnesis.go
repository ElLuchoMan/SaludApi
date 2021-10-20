package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type OdontologiaAnamnesis struct {
	IdOdontologiaAnamnesis  int       `orm:"column(id_anamnesis);pk;auto"`
	id_historia_clinica     int       `orm:"column(id_historia_clinica);null"`
	tratamiento             string    `orm:"column(tratamiento);null"`
	medicamentos            string    `orm:"column(medicamento);null"`
	alergias                string    `orm:"column(alergias);null"`
	hemorragias             string    `orm:"column(hemorragias);null"`
	irradiaciones           string    `orm:"column(irradiaciones);null"`
	sinusitis               string    `orm:"column(sinusitis);null"`
	enfermedad_respiratoria string    `orm:"column(enfermedad_respiratoria);null"`
	cardiopatias            string    `orm:"column(cardiopatias);null"`
	diabetes                string    `orm:"column(diabetes);null"`
	fiebre_reumatica        string    `orm:"column(fiebre_reumatica);null"`
	hepatitis               string    `orm:"column(hepatitis);null"`
	hipertension            string    `orm:"column(hipertension);null"`
	antecedente_familiar    string    `orm:"column(antecedente_familiar);null"`
	cepillado               string    `orm:"column(cepillado);null"`
	ceda                    string    `orm:"column(ceda);null"`
	enjuague                string    `orm:"column(enjuague);null"`
	dulces                  string    `orm:"column(dulces);null"`
	fuma                    string    `orm:"column(fuma);null"`
	chicle                  string    `orm:"column(chicle);null"`
	otras                   string    `orm:"column(otras);null"`
	ultima_visita           time.Time `orm:"column(ultima_visita);type(date);null"`
}

func (t *OdontologiaAnamnesis) Schema() string {
	return "odontologia"
}
func (t *OdontologiaAnamnesis) TableName() string {
	return "anamnesis"
}
func init() {
	orm.RegisterModel(new(OdontologiaAnamnesis))
}

// AddAnamnesis insert a new Usuario into database and returns
// Último registro insertado con éxito
func AddAnamnesis(m *OdontologiaAnamnesis) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//GetAnamnesisById consulta una anamnesis por su id
//Id no existe
func GetAnamnesisById(id int) (v *OdontologiaAnamnesis, err error) {
	o := orm.NewOrm()
	v = &OdontologiaAnamnesis{IdOdontologiaAnamnesis: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetAllAnamnesis consulta todas las anamnesis
//No existen registros
func GetAllAnamnesis(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OdontologiaAnamnesis))
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

//UpdateAnamnesis actualiza una anamnesis por su id
//El registro a actualizar no existe
func UpdateAnamnesisById(m *OdontologiaAnamnesis) (err error) {
	o := orm.NewOrm()
	v := OdontologiaAnamnesis{IdOdontologiaAnamnesis: m.IdOdontologiaAnamnesis}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Número de registros actualizados en la base de datos:", num)
		}
	}
	return
}

//DeleteAnamnesis elimina una anamnesis por su id
//El registro a eliminar no existe
func DeleteAnamnesis(id int) (err error) {
	o := orm.NewOrm()
	v := OdontologiaAnamnesis{IdOdontologiaAnamnesis: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OdontologiaAnamnesis{IdOdontologiaAnamnesis: id}); err == nil {
			fmt.Println("Número de registros eliminados en la base de datos:", num)
		}
	}
	return
}
