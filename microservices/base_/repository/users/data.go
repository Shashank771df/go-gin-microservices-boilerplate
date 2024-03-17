package users

import (
	"app/core/database"
	"app/core/date"
	"app/core/logger"
	"app/microservices/base_/lib"
	"app/microservices/base_/log"
	"fmt"
	"strings"
)

const alias = "pl"
const table = "Users"
const colPk = "userId"

type Data struct {
}

func (o Data) FiltersAllowed() map[string]database.FilterAllowed {
	return map[string]database.FilterAllowed{
		"origin": {
			Column:  fmt.Sprintf("%v.%v", alias, "origin"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "origin",
				"label":     "Origin",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el Origin",
			},
		},
		"email": {
			Column:  fmt.Sprintf("%v.%v", alias, "email"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "email",
				"label":     "Email",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el Email",
			},
		},
		"icCode": {
			Column:  fmt.Sprintf("%v.%v", alias, "icCode"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "icCode",
				"label":     "icCode",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el icCode",
			},
		},
		"phone": {
			Column:  fmt.Sprintf("%v.%v", alias, "phone"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "phone",
				"label":     "phone",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el phone",
			},
		},
		"ip": {
			Column:  fmt.Sprintf("%v.%v", alias, "ip"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "ip",
				"label":     "ip",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el ip",
			},
		},
		"birthDate": {
			Column:  fmt.Sprintf("%v.%v", alias, "birthDate"),
			Order:   "desc",
			Pattern: database.BetweenPattern,
			Public: map[string]interface{}{
				"key":       "birthDate",
				"label":     "birthDate",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el birthDate",
			},
		},
		"birthYear": {
			Column:  fmt.Sprintf("%v.%v", alias, "birthYear"),
			Order:   "desc",
			Pattern: database.BetweenPattern,
			Public: map[string]interface{}{
				"key":       "birthYear",
				"label":     "birthYear",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el birthYear",
			},
		},
		"birthMonth": {
			Column:  fmt.Sprintf("%v.%v", alias, "birthMonth"),
			Order:   "desc",
			Pattern: database.BetweenPattern,
			Public: map[string]interface{}{
				"key":       "birthMonth",
				"label":     "birthMonth",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el birthMonth",
			},
		},
		"countryId": {
			Column:  fmt.Sprintf("%v.%v", alias, "countryId"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "countryId",
				"label":     "countryId",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el countryId",
			},
		},
		"completed": {
			Column:  fmt.Sprintf("%v.%v", alias, "completed"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "completed",
				"label":     "completed",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el completed",
			},
		},
		"active": {
			Column:  fmt.Sprintf("%v.%v", alias, "active"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "active",
				"label":     "active",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el active",
			},
		},
		"tester": {
			Column:  fmt.Sprintf("%v.%v", alias, "active"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "active",
				"label":     "active",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el active",
			},
		},
		"regUserId": {
			Column:  fmt.Sprintf("%v.%v", alias, "regUserId"),
			Order:   "desc",
			Pattern: database.EqualPattern,
			Public: map[string]interface{}{
				"key":       "regUserId",
				"label":     "regUserId",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el regUserId",
			},
		},
		"regDate": {
			Column:  fmt.Sprintf("%v.%v", alias, "regDate"),
			Order:   "desc",
			Pattern: database.BetweenPattern,
			Public: map[string]interface{}{
				"key":       "regDate",
				"label":     "regDate",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el regDate",
			},
		},
		"regTimestamp": {
			Column:  fmt.Sprintf("%v.%v", alias, "regTimestamp"),
			Order:   "desc",
			Pattern: database.BetweenPattern,
			Public: map[string]interface{}{
				"key":       "regTimestamp",
				"label":     "regTimestamp",
				"hasOption": false,
				"options":   []string{},
				"hint":      "Ingresa el regTimestamp",
			},
		},
	}
}

func (o Data) OrdersAllowed() map[string]database.OrdersAllowed {
	return map[string]database.OrdersAllowed{
		"origin": {
			Column: fmt.Sprintf("%v.%v", alias, "origin"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "origin",
				"label": "Origin",
			},
		},
		"email": {
			Column: fmt.Sprintf("%v.%v", alias, "email"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "email",
				"label": "Email",
			},
		},
		"icCode": {
			Column: fmt.Sprintf("%v.%v", alias, "icCode"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "icCode",
				"label": "icCode",
			},
		},
		"phone": {
			Column: fmt.Sprintf("%v.%v", alias, "phone"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "phone",
				"label": "phone",
			},
		},
		"ip": {
			Column: fmt.Sprintf("%v.%v", alias, "ip"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "ip",
				"label": "ip",
			},
		},
		"birthDate": {
			Column: fmt.Sprintf("%v.%v", alias, "birthDate"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "birthDate",
				"label": "birthDate",
			},
		},
		"birthYear": {
			Column: fmt.Sprintf("%v.%v", alias, "birthYear"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "birthYear",
				"label": "birthYear",
			},
		},
		"birthMonth": {
			Column: fmt.Sprintf("%v.%v", alias, "birthMonth"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "birthMonth",
				"label": "birthMonth",
			},
		},
		"countryId": {
			Column: fmt.Sprintf("%v.%v", alias, "countryId"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "countryId",
				"label": "countryId",
			},
		},
		"completed": {
			Column: fmt.Sprintf("%v.%v", alias, "completed"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "completed",
				"label": "completed",
			},
		},
		"active": {
			Column: fmt.Sprintf("%v.%v", alias, "active"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "active",
				"label": "active",
			},
		},
		"tester": {
			Column: fmt.Sprintf("%v.%v", alias, "active"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "active",
				"label": "active",
			},
		},
		"regUserId": {
			Column: fmt.Sprintf("%v.%v", alias, "regUserId"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "regUserId",
				"label": "regUserId",
			},
		},
		"regDate": {
			Column: fmt.Sprintf("%v.%v", alias, "regDate"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "regDate",
				"label": "regDate",
			},
		},
		"regTimestamp": {
			Column: fmt.Sprintf("%v.%v", alias, "regTimestamp"),
			Order:  "desc",
			Public: map[string]interface{}{
				"key":   "regTimestamp",
				"label": "regTimestamp",
			},
		},
	}
}

func (o Data) BuildWhere(data Where) string {
	var utildb database.UtilsDB
	return utildb.WhereScript(data, " AND ", data.Alias)
}

func (o Data) ItemDB(data ItemDB) database.ItemOut {
	var out Users
	var utildb database.UtilsDB

	where := o.BuildWhere(Where{
		UserId:    data.UserId,
		Email:     data.Email,
		IcCode:    data.IcCode,
		Phone:     data.Phone,
		DocIdType: data.DocIdType,
		DocId:     data.DocId,
	})

	columnsScript := utildb.ColsScript(Users{}, alias)
	mainScript := fmt.Sprintf("SELECT %s FROM %s %s WHERE %s LIMIT 1", columnsScript, table, alias, where)

	return lib.DB.ItemDB(database.ItemIn{
		Script: mainScript,
		Trace:  data.Trace,
		Dest:   &out,
	})
}

func (o Data) ItemsDB(data ItemsDB) database.ItemsOut {
	var out []Users
	var utildb database.UtilsDB

	if data.Order == "" {
		data.Order = "desc"
	}

	where := o.BuildWhere(Where{
		IcCode:       data.IcCode,
		Phone:        data.Phone,
		Ip:           data.Ip,
		BirthDate:    data.BirthDate,
		BirthYear:    data.BirthYear,
		BirthMonth:   data.BirthMonth,
		CountryId:    data.CountryId,
		DocIdType:    data.DocIdType,
		Completed:    data.Completed,
		Active:       data.Active,
		Tester:       data.Tester,
		RegUserId:    data.RegUserId,
		RegDate:      data.RegDate,
		RegTimestamp: data.RegTimestamp,
	})

	columnsScript := utildb.ColsScript(Users{}, alias)
	mainScript := fmt.Sprintf("SELECT %s FROM %s %s", columnsScript, table, alias)
	affectedRowsScript := utildb.AffectedRowsScript(fmt.Sprintf("%s %s", table, alias), where)
	defaultOrferBy := fmt.Sprintf(" ORDER BY %s.%s %s", alias, colPk, data.Order)

	return lib.DB.ItemsDB(database.ItemsIn{
		SelectScript:       mainScript,
		WhereScript:        where,
		FilterVals:         data.FilterVals,
		OrdersVals:         data.OrderVals,
		CounterScript:      affectedRowsScript,
		DefaultOrderBy:     defaultOrferBy,
		FiltersAllowed:     o.FiltersAllowed(),
		OrdersAllowed:      o.OrdersAllowed(),
		EnableDefaultLimit: false,
		DefaultLimit:       1,
		EnablePaging:       data.EnablePaging,
		PagingSize:         data.PagingSize,
		PagingIndex:        data.PagingIndex,
		Label:              data.Label,
		Trace:              data.Trace,
		Dest:               &out,
	})
}

func (o Data) NewItemDB(data NewItemDB) database.ItemOut {
	var utildb database.UtilsDB

	current := date.UtilDate{}.CurrentTimeUTC()

	parameters := Users{
		Email:        data.Email,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		IcCode:       data.IcCode,
		Phone:        data.Phone,
		Ip:           &data.Ip,
		BirthDate:    data.BirthDate,
		BirthYear:    data.BirthYear,
		BirthMonth:   data.BirthMonth,
		CountryId:    data.CountryId,
		Address:      data.Address,
		DocIdType:    data.DocIdType,
		DocId:        data.DocId,
		Completed:    data.Completed,
		Active:       data.Active,
		Tester:       data.Tester,
		RegUserId:    data.RegUserId,
		RegDate:      &current.Date,
		RegDatetime:  &current.DateTime,
		RegTimestamp: &current.TimeStamp,
	}

	script := utildb.InsertScript(table, parameters)

	ret := lib.DB.NewItemDB(database.NewItemIn{
		Script: script,
		Trace:  data.Trace,
		Label:  data.Label,
	})

	if !ret.Success {
		return database.ItemOut{Message: ret.Message}
	}

	userId := uint32(ret.InsertId)

	return o.ItemDB(ItemDB{
		UserId: &userId,
		Trace:  data.Trace,
		Label:  data.Label,
	})
}

func (o Data) UpdateItemDB(data UpdateItemDB) database.ItemOut {
	var utildb database.UtilsDB
	parameters := Users{
		Email:      data.Email,
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		IcCode:     data.IcCode,
		Phone:      data.Phone,
		BirthDate:  data.BirthDate,
		BirthYear:  data.BirthYear,
		BirthMonth: data.BirthMonth,
		CountryId:  data.CountryId,
		Address:    data.Address,
		DocIdType:  data.DocIdType,
		DocId:      data.DocId,
		Completed:  data.Completed,
		Active:     data.Active,
	}

	where := o.BuildWhere(Where{
		UserId: data.UserId,
	})

	script := utildb.UpdateScript(table, parameters, where)
	affectedRowsScript := utildb.AffectedRowsScript(table, where)

	ret := lib.DB.UpdateItemDB(database.UpdateItemIn{
		Script:             script,
		AffectedRowsScript: affectedRowsScript,
		Trace:              data.Trace,
		Label:              data.Label,
	})

	if !ret.Success {
		return database.ItemOut{Message: ret.Message}
	}

	current := date.UtilDate{}.CurrentTimeUTC()
	log.Log.Info(logger.LogInfo{
		Key:   "users-upd-user-id",
		Value: "Update item by user id",
		Extra: map[string]interface{}{
			"data":      data,
			"date":      current.Date,
			"dateTime":  current.DateTime,
			"timestamp": current.TimeStamp,
		},
	})

	return o.ItemDB(ItemDB{
		UserId: data.UserId,
		Trace:  data.Trace,
		Label:  data.Label,
	})
}

func (o Data) RemoveItemDB(data RemoveItemDB) database.RemoveItemOut {
	var utildb database.UtilsDB

	where := o.BuildWhere(Where{
		UserId: data.UserId,
	})

	script := utildb.DeleteScript(table, where)
	affectedRowsScript := utildb.AffectedRowsScript(table, where)

	ret := lib.DB.RemoveItemDB(database.RemoveItemIn{
		Script:             script,
		Label:              data.Label,
		Trace:              data.Trace,
		AffectedRowsScript: affectedRowsScript,
	})

	if ret.Success {
		current := date.UtilDate{}.CurrentTimeUTC()
		log.Log.Info(logger.LogInfo{
			Key:   "users-del-user-id",
			Value: "Update item by user id",
			Extra: map[string]interface{}{
				"data":      data,
				"date":      current.Date,
				"dateTime":  current.DateTime,
				"timestamp": current.TimeStamp,
			},
		})
	}

	return ret
}

func (o Data) ValidateNewItemDB(data ValidateNewItemDB) database.ItemsOut {
	var out []Users
	var utildb database.UtilsDB

	columnsScript := utildb.ColsScript(Users{}, alias)

	var whereOR []string

	if data.DocId != nil && *data.DocId != "" {
		whereOR = append(whereOR, fmt.Sprintf("(%s.docId = '%v' AND %s.docIdType = '%v')", alias, *data.DocId, alias, *data.DocIdType))
	}

	if data.Phone != nil && *data.Phone != "" {
		whereOR = append(whereOR, fmt.Sprintf("(%s.phone = '%v' AND %s.icCode = '%v')", alias, *data.Phone, alias, *data.IcCode))
	}

	if data.Email != nil && *data.Email != "" {
		whereOR = append(whereOR, fmt.Sprintf(" OR %s.email = '%v'", alias, *data.Email))
	}

	where := strings.Join(whereOR, " OR ")

	mainScript := fmt.Sprintf("SELECT %s FROM %s %s", columnsScript, table, alias)
	affectedRowsScript := utildb.AffectedRowsScript(fmt.Sprintf("%s %s", table, alias), "")
	defaultOrferBy := fmt.Sprintf(" ORDER BY %s.%s desc", alias, colPk)

	return lib.DB.ItemsDB(database.ItemsIn{
		SelectScript:       mainScript,
		WhereScript:        where,
		CounterScript:      affectedRowsScript,
		DefaultOrderBy:     defaultOrferBy,
		FiltersAllowed:     o.FiltersAllowed(),
		OrdersAllowed:      o.OrdersAllowed(),
		EnableDefaultLimit: false,
		DefaultLimit:       1,
		Label:              data.Label,
		Trace:              data.Trace,
		Dest:               &out,
	})
}

func (o Data) ValidateUpdateItemDB(data ValidateUpdateItemDB) database.ItemsOut {
	var out []Users
	var utildb database.UtilsDB

	columnsScript := utildb.ColsScript(Users{}, alias)
	where := ""

	var whereOR []string

	if data.DocId != nil && *data.DocId != "" {
		whereOR = append(whereOR, fmt.Sprintf("(%s.docId = '%v' AND %s.docIdType = '%v')", alias, *data.DocId, alias, *data.DocIdType))
	}

	if data.Phone != nil && *data.Phone != "" {
		whereOR = append(whereOR, fmt.Sprintf("(%s.phone = '%v' AND %s.icCode = '%v')", alias, *data.Phone, alias, *data.IcCode))
	}

	if data.Email != nil && *data.Email != "" {
		whereOR = append(whereOR, fmt.Sprintf(" OR %s.email = '%v'", alias, *data.Email))
	}

	if len(whereOR) == 0 {
		where = fmt.Sprintf("%s.userId != '%v'", alias, *data.UserId)
	} else {
		where += fmt.Sprintf("(%v) AND %s.userId != '%v'", strings.Join(whereOR, " OR "), alias, *data.UserId)
	}

	mainScript := fmt.Sprintf("SELECT %s FROM %s %s", columnsScript, table, alias)
	affectedRowsScript := utildb.AffectedRowsScript(fmt.Sprintf("%v %v", table, alias), "")
	defaultOrferBy := fmt.Sprintf(" ORDER BY %s.%s desc", alias, colPk)

	return lib.DB.ItemsDB(database.ItemsIn{
		SelectScript:       mainScript,
		WhereScript:        where,
		CounterScript:      affectedRowsScript,
		DefaultOrderBy:     defaultOrferBy,
		FiltersAllowed:     o.FiltersAllowed(),
		OrdersAllowed:      o.OrdersAllowed(),
		EnableDefaultLimit: false,
		DefaultLimit:       1,
		Label:              data.Label,
		Trace:              data.Trace,
		Dest:               &out,
	})
}

func (o Data) ItemsGroupByMonthDB() database.ItemsOut {
	var out []ItemsGroupByMonthDB

	mainScript := fmt.Sprintf(`Select birthMonth,count(*) As users From %s Group By birthMonth`, table)

	return lib.DB.ItemsDB(database.ItemsIn{
		SelectScript:       mainScript,
		EnableDefaultLimit: false,
		DefaultLimit:       1,
		Dest:               &out,
	})
}
