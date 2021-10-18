package models

import (
	"github.com/astaxie/beego/orm"
)

type Prueba struct {
	Id     int    `orm:"column(id);pk;auto"`
	Nombre string `orm:"column(nombre)"`
}

func (t *Prueba) TableName() string {
	return "prueba"
}
func init() {
	orm.RegisterModel(new(Prueba))
}

// AddPrueba insert a new Prueba into database and returns
// last inserted Id on success.
func AddPrueba(m *Prueba) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAllRol retrieves all Rol matches certain condition. Returns empty list if
func GetAllPrueba() (objects []Prueba, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Prueba)).All(&objects)
	return
}

//GetA returns a single Prueba
func GetPruebaByNombre(nombre string) (v *Prueba, err error) {
	o := orm.NewOrm()
	v = &Prueba{Nombre: nombre}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
