package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AsignacionEvaluadorItem struct {
	Id                       int                     `orm:"column(id);pk;auto"`
	AsignacionEvaluadorId    *AsignacionEvaluador    `orm:"column(asignacion_evaluador_id);rel(fk)"`
	ItemId                   *Item                   `orm:"column(item_id);rel(fk)"`
	Activo                   bool                    `orm:"column(activo);default(true)"`
	RolAsignacionEvaluadorId *RolAsignacionEvaluador `orm:"column(rol_asignacion_evaluador_id);rel(fk)"`
	FechaCreacion            time.Time               `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion        time.Time               `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone);null"`
}

func (t *AsignacionEvaluadorItem) TableName() string {
	return "asignacion_evaluador_item"
}

func init() {
	orm.RegisterModel(new(AsignacionEvaluadorItem))
}

// AddAsignacionEvaluadorItem insert a new AsignacionEvaluadorItem into database and returns
// last inserted Id on success.
// func AddAsignacionEvaluadorItem(m *AsignacionEvaluadorItem) (id int64, err error) {
// 	o := orm.NewOrm()
// 	m.Activo = true
// 	id, err = o.Insert(m)
// 	return
// }

// GetAsignacionEvaluadorItemById retrieves AsignacionEvaluadorItem by Id. Returns error if
// Id doesn't exist
func GetAsignacionEvaluadorItemById(id int) (v *AsignacionEvaluadorItem, err error) {
	o := orm.NewOrm()
	v = &AsignacionEvaluadorItem{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAsignacionEvaluadorItem retrieves all AsignacionEvaluadorItem matches certain condition. Returns empty list if
// no records exist
func GetAllAsignacionEvaluadorItem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AsignacionEvaluadorItem)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.HasSuffix(k, "in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []AsignacionEvaluadorItem
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
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

// UpdateAsignacionEvaluadorItem updates AsignacionEvaluadorItem by Id and returns error if
// the record to be updated doesn't exist
func UpdateAsignacionEvaluadorItemById(m *AsignacionEvaluadorItem) (err error) {
	o := orm.NewOrm()
	v := AsignacionEvaluadorItem{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAsignacionEvaluadorItem deletes AsignacionEvaluadorItem by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAsignacionEvaluadorItem(id int) (err error) {
	o := orm.NewOrm()
	v := AsignacionEvaluadorItem{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		v.Activo = false
		if num, err = o.Update(&v); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// ReadOrUpdateAsignacionEvaluadorItem reads an AsignacionEvaluador by personaId and evaluacionId,
// updates it if exists, or creates it if not.
func CreateOrUpdateAsignacionEvaluadorItem(m *AsignacionEvaluadorItem) (created bool, err error) {
	o := orm.NewOrm()

	var existing AsignacionEvaluadorItem
	err = o.QueryTable(new(AsignacionEvaluadorItem)).
		Filter("AsignacionEvaluadorId__Id", m.AsignacionEvaluadorId.Id).
		Filter("ItemId__Id", m.ItemId.Id).
		Filter("Activo", true).
		One(&existing)

	if err == orm.ErrNoRows {
		m.Activo = true
		if _, err = o.Insert(m); err == nil {
			created = true
			return created, nil
		}
	} else if err == nil {
		m.Id = existing.Id
		m.FechaCreacion = existing.FechaCreacion
		m.Activo = existing.Activo
		if _, err = o.Update(m); err == nil {
			created = false
			return created, nil
		}
	}

	return false, err
}
