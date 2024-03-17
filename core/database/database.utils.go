package database

import (
	"app/core/utils"
	"app/core/validator"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type UtilsDB struct {
}

func (mysql UtilsDB) FiltersBuilder(i interface{}) string {
	filters := []map[string]interface{}{}

	refVal := reflect.ValueOf(i)
	refType := reflect.TypeOf(i)
	for x := 0; x < refVal.NumField(); x++ {
		if refVal.Field(x).Kind() == reflect.Ptr && !refVal.Field(x).IsNil() {
			field := refVal.Field(x)
			val := fmt.Sprintf("%v", field.Elem())

			key := refType.Field(x).Tag.Get("key")

			filters = append(filters, map[string]interface{}{
				"key": key,
				"val": val,
			})
		}
	}

	if len(filters) == 0 {
		return ""
	}

	return utils.UtilsMap{}.MapToString(filters)
}

// UpdateScript .
func (mysql UtilsDB) UpdateScript(table string, i interface{}, where string) string {
	var values []string

	refVal := reflect.ValueOf(i)
	refType := reflect.TypeOf(i)

	for x := 0; x < refVal.NumField(); x++ {
		if refVal.Field(x).Kind() == reflect.Ptr && !refVal.Field(x).IsNil() {
			field := refVal.Field(x)
			type1 := field.Type().String()
			val := fmt.Sprintf("%v", field.Elem())

			if type1 == "*string" || type1 == "string" {
				forceEmptyUpdate := refType.Field(x).Tag.Get("forceEmptyUpdate")
				if val == "" && forceEmptyUpdate != "true" {
					continue
				}

				val = fmt.Sprintf("'%v'", field.Elem())
			}
			col := refType.Field(x).Tag.Get("db")
			values = append(values, fmt.Sprintf("%v = %v", col, val))
		}
	}

	updScript := strings.Join(values, ", ")
	script := fmt.Sprintf(`UPDATE %v SET %v WHERE %v`, table, updScript, where)

	return script
}

func (mysql UtilsDB) DeleteScript(table string, where string) string {
	return fmt.Sprintf(`DELETE FROM %v WHERE %v`, table, where)
}

func (mysql UtilsDB) AffectedRowsScript(table string, where string) string {
	if where == "" {
		return fmt.Sprintf(`SELECT count(*) AS itemsCounter FROM %v`, table)
	} else {
		return fmt.Sprintf(`SELECT count(*) AS itemsCounter FROM %v WHERE %v`, table, where)
	}
}

// InsertScript .
func (mysql UtilsDB) InsertScript(table string, i interface{}) string {
	var values []string
	var cols []string

	refVal := reflect.ValueOf(i)
	refType := reflect.TypeOf(i)

	for x := 0; x < refVal.NumField(); x++ {
		if refVal.Field(x).Kind() == reflect.Ptr && !refVal.Field(x).IsNil() {
			field := refVal.Field(x)
			type1 := field.Type().String()
			val := fmt.Sprintf("%v", field.Elem())

			if type1 == "*string" || type1 == "string" {
				forceEmptyInsert := refType.Field(x).Tag.Get("forceEmptyInsert")
				if val == "" && forceEmptyInsert != "true" {
					continue
				}

				val = fmt.Sprintf("'%v'", field.Elem())
			}

			col := refType.Field(x).Tag.Get("db")
			values = append(values, val)
			cols = append(cols, col)
		}
	}

	insScript := fmt.Sprintf(`(%s) VALUES (%s)`, strings.Join(cols, ", "), strings.Join(values, ", "))
	script := fmt.Sprintf(`INSERT INTO %s %s`, table, insScript)

	return script
}

// InsertIntoScript .
func (mysql UtilsDB) InsertIntoScript(table string, i interface{}) string {
	var cols []string

	refType := reflect.TypeOf(i)

	for x := 0; x < refType.NumField(); x++ {
		col := refType.Field(x).Tag.Get("db")
		cols = append(cols, col)
	}

	return fmt.Sprintf(`INSERT INTO %s(%s) VALUES`, table, strings.Join(cols, ", "))
}

// ValuesScript .
func (mysql UtilsDB) ValuesScript(i interface{}) string {
	var values []string

	refVal := reflect.ValueOf(i)

	for x := 0; x < refVal.NumField(); x++ {
		if refVal.Field(x).Kind() == reflect.Ptr && !refVal.Field(x).IsNil() {
			field := refVal.Field(x)
			type1 := field.Type().String()
			val := fmt.Sprintf("%v", field.Elem())

			if type1 == "*string" || type1 == "string" {
				val = fmt.Sprintf("'%v'", field.Elem())
			}

			values = append(values, val)
		}
	}

	return fmt.Sprintf(`(%s)`, strings.Join(values, ", "))
}

// WhereScript .
func (mysql UtilsDB) WhereScript(i interface{}, joinClause string, alias string) string {
	var whereItems []string

	refVal := reflect.ValueOf(i)
	refType := reflect.TypeOf(i)

	for x := 0; x < refVal.NumField(); x++ {
		if refVal.Field(x).Kind() == reflect.Ptr && !refVal.Field(x).IsNil() {
			pattern := refType.Field(x).Tag.Get("pattern")
			field := refVal.Field(x)
			type1 := field.Type().String()

			val := fmt.Sprintf("%v", field.Elem())
			if type1 == "*string" || type1 == "string" {
				if val == "" {
					continue
				}

				val = fmt.Sprintf("'%v'", field.Elem())
			}

			col := refType.Field(x).Tag.Get("db")

			where := fmt.Sprintf("%v %v %v", col, pattern, val)
			if alias != "" {
				where = fmt.Sprintf("%v.%v %v %v", alias, col, pattern, val)
			}
			whereItems = append(whereItems, where)
		}
	}

	return strings.Join(whereItems, joinClause)
}

// ColsScript .
func (mysql UtilsDB) ColsScript(i interface{}, alias string) string {
	var cols []string

	refVal := reflect.ValueOf(i)
	refType := reflect.TypeOf(i)

	for x := 0; x < refType.NumField(); x++ {
		col := refType.Field(x).Tag.Get("db")
		enableDefault := refType.Field(x).Tag.Get("enableDefault")

		if enableDefault == "true" {
			type1 := refVal.Field(x).Type().String()

			if type1 == "*string" || type1 == "string" {
				cols = append(cols, fmt.Sprintf("IfNull(%v.%v,'') AS %v", alias, col, col))
			} else {
				cols = append(cols, fmt.Sprintf("IfNull(%v.%v,0) AS %v", alias, col, col))
			}

		} else {
			cols = append(cols, fmt.Sprintf("%v.%v AS %v", alias, col, col))
		}

	}
	return strings.Join(cols, ", ")
}

func (obj UtilsDB) BuildFilter(filtersAllowed map[string]FilterAllowed, filtersVals string) string {
	// If there is no filters allowed
	if len(filtersAllowed) == 0 {
		return ""
	}
	// If there is no filters
	if len(filtersVals) == 0 {
		return ""
	}

	filterObj := []validator.FilterJoi{}
	err := json.Unmarshal([]byte(filtersVals), &filterObj)
	if err != nil {
		return ""
	}

	filters := []string{}

	for _, item := range filterObj {
		value, ok := filtersAllowed[item.Key]

		if !ok {
			return ""
		}

		pattern := ""

		if value.Pattern == BetweenPattern {
			pattern = fmt.Sprintf(value.Pattern, value.Column, item.Val, item.Val1)
		} else {
			pattern = fmt.Sprintf(value.Pattern, value.Column, item.Val)
		}

		filters = append(filters, pattern)
	}

	return strings.Join(filters, " AND ")
}

func (obj UtilsDB) BuildOrder(filtersAllowed map[string]FilterAllowed, filtersVals string,
	ordersAllowed map[string]OrdersAllowed, orderVals string) string {
	order := []string{}

	filterObj := []validator.FilterJoi{}
	json.Unmarshal([]byte(filtersVals), &filterObj)

	orderObj := []validator.OrderJoi{}
	json.Unmarshal([]byte(orderVals), &orderObj)

	if len(filterObj) == 0 && len(orderObj) == 0 {
		return ""
	}

	// If there is orders allowed
	if len(ordersAllowed) > 0 && len(orderObj) > 0 {
		for _, item := range orderObj {
			value, ok := ordersAllowed[item.Key]

			//Enviar un error
			if !ok {
				return ""
			}
			orderType := item.Order
			if orderType == "" {
				orderType = value.Order
			}

			order = append(order, fmt.Sprintf("%v %v", value.Column, orderType))
		}
	}

	// If there are filters allowed
	if len(filtersAllowed) > 0 && len(filterObj) > 0 {
		for _, item := range filterObj {
			value, ok := filtersAllowed[item.Key]

			if !ok {
				return ""
			}

			orderType := item.Order
			if orderType == "" {
				orderType = value.Order
			}

			order = append(order, fmt.Sprintf("%v %v", value.Column, orderType))
		}
	}

	return strings.Join(order, " , ")
}

func (obj UtilsDB) BuildPaging(data ItemsIn) string {
	var limiting string = ""
	var offset uint32 = 0

	if !data.EnablePaging {
		if data.EnableDefaultLimit {
			return fmt.Sprintf(" LIMIT %d", data.DefaultLimit)
		}
		return limiting
	}

	if data.PagingIndex == 0 {
		data.PagingIndex = 1
	}

	offset = (data.PagingIndex - 1) * data.PagingSize
	limiting = fmt.Sprintf(" LIMIT %d OFFSET %d", data.PagingSize, offset)

	return limiting
}
