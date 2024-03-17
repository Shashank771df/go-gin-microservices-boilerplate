package database

import (
	"app/core/errors"
	"app/core/lang"
	"app/core/logger"
	"app/core/scan"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type (
	MySQL struct {
		instance *sql.DB
		log      logger.Logger
	}
)

func (mysql *MySQL) GetConnection() *sql.DB {
	return mysql.instance
}

func (o *MySQL) initialize(mc DBProps) {
	// dns := fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%d)/%s?parseTime=true",
	// 	mc.User, mc.Pass, mc.Host, mc.Port, mc.Name,
	// )
	dns := fmt.Sprintf(
		"%s@tcp(%s:%d)/%s?parseTime=true",
		mc.User, mc.Host, mc.Port, mc.Name,
	)

	sqlDB, err := sql.Open("mysql", dns)
	if err != nil {
		mc.Log.Warn(logger.LogInfo{
			Key:   "DB-UTILS",
			Value: fmt.Sprintf("Can't open db: %v", err),
		})
	}

	if err = sqlDB.Ping(); err != nil {
		mc.Log.Warn(logger.LogInfo{
			Key:   "DB-UTILS",
			Value: fmt.Sprintf("Can't do ping: %v", err),
		})
	}

	sqlDB.SetConnMaxLifetime(time.Duration(mc.MaxLifeTimeConns) * time.Second)
	sqlDB.SetMaxIdleConns(int(mc.MaxIdleConns))
	sqlDB.SetMaxOpenConns(int(mc.MaxOpenConns))

	o.instance = sqlDB
	o.log = mc.Log
}

func (o *MySQL) SetLogger(log logger.Logger) {
	o.log = log
}

func (obj MySQL) RunScript(in RunScriptIn) RunScriptOut {
	db := obj.instance
	result, err := db.Exec(in.Script)

	if err != nil {
		message := fmt.Sprintf("RunScript: Script returned %v error. Script: %s", err, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	val, err := result.RowsAffected()

	if err != nil {
		message := fmt.Sprintf("RunScript: Invalid rows affected %v error. Script: %s", err, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	lastid, _ := result.LastInsertId()

	return RunScriptOut{
		LastInsertId: lastid,
		AffectedRows: int32(val),
		ChangedRows:  int32(val),
	}
}

func (obj MySQL) SelectScript(in SelectScriptIn) SelectScriptOut {
	db := obj.instance
	ret := SelectScriptOut{}

	rows, err := db.Query(in.Script)

	if err != nil {
		message := fmt.Sprintf("SelectScript: Script returned %v error. Script: %s", err, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	defer rows.Close()

	_, err = rows.Columns()

	if err != nil {
		message := fmt.Sprintf("SelectScript: Invalid columns %v error. Script: %s", err, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	ret.DataLen, _ = scan.Rows(in.Dest, rows)

	return ret
}

func (obj MySQL) ItemDB(in ItemIn) ItemOut {
	if in.Dest == nil {
		return ItemOut{Success: false, Message: fmt.Sprintf("Dst model must exists in ItemIn but is %v", in.Dest)}
	}

	obj.log.Debug(logger.LogInfo{Key: "ItemDB", Value: fmt.Sprintf("%s %s", "SELECT SCRIPT - ", in.Script)})

	result := obj.SelectScript(SelectScriptIn{Script: in.Script, Trace: in.Trace, Dest: in.Dest})

	if result.Error != nil {
		message := fmt.Sprintf("ItemDB: Script have error %v. Script: %v", result.Error.Error(), in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	if in.CheckDuplicated && result.DataLen > 1 {
		message := fmt.Sprintf("ItemDB: Script must return 0 or 1 item. ItemsFound %v. Script: %v", result.DataLen, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	if result.DataLen == 0 {
		in.Dest = nil
	}

	return ItemOut{
		Success:   true,
		ItemFound: result.DataLen == 1,
		Item:      in.Dest,
	}
}

func (obj MySQL) ItemsDB(in ItemsIn) ItemsOut {
	var dbUtil UtilsDB

	if in.Dest == nil {
		message := fmt.Sprintf("ItemsDB: Dst model must exists in ItemsIn but is %v", in.Dest)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	destValue := reflect.ValueOf(in.Dest)

	if destValue.Elem().Type().Kind() != reflect.Slice {
		message := "ItemsDB: Dst model must be a slice"
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	filterScript := dbUtil.BuildFilter(in.FiltersAllowed, in.FilterVals)
	orderFilterScript := dbUtil.BuildOrder(in.FiltersAllowed, in.FilterVals, in.OrdersAllowed, in.OrdersVals)
	pagingScript := dbUtil.BuildPaging(in)
	whereClause := ""
	orderScript := in.DefaultOrderBy

	if in.WhereScript != "" || filterScript != "" {
		whereClause = " WHERE "
	}
	if orderFilterScript != "" {
		orderScript = " ORDER BY " + orderFilterScript
	}

	//Building main script
	mainScript := in.SelectScript + whereClause + in.WhereScript
	if whereClause != "" && in.WhereScript != "" && filterScript != "" {
		mainScript += " And " + filterScript
	} else {
		mainScript += filterScript
	}
	mainScript += orderScript + pagingScript

	//Building counter script
	counterFilterScript := in.CounterScript + whereClause + in.WhereScript
	if whereClause != "" && in.WhereScript != "" && filterScript != "" {
		counterFilterScript += " And " + filterScript
	} else {
		counterFilterScript += filterScript
	}

	obj.log.Debug(logger.LogInfo{Key: "ItemsDB", Value: fmt.Sprintf("%s %s", "SELECT SCRIPT - ", in.SelectScript)})
	obj.log.Debug(logger.LogInfo{Key: "ItemsDB", Value: fmt.Sprintf("%s %s", "WHERE SCRIPT - ", in.WhereScript)})
	obj.log.Debug(logger.LogInfo{Key: "ItemsDB", Value: fmt.Sprintf("%s %s", "FILTER SCRIPT - ", filterScript)})
	obj.log.Debug(logger.LogInfo{Key: "ItemsDB", Value: fmt.Sprintf("%s %s", "ORDER SCRIPT - ", orderScript)})
	obj.log.Debug(logger.LogInfo{Key: "ItemsDB", Value: fmt.Sprintf("%s %s", "PAGINGS SCRIPT - ", pagingScript)})
	obj.log.Debug(logger.LogInfo{Key: "ItemsDB", Value: fmt.Sprintf("%s %s", "MAIN SCRIPT - ", mainScript)})
	obj.log.Debug(logger.LogInfo{Key: "ItemsDB", Value: fmt.Sprintf("%s %s", "COUNTER SCRIPT - ", counterFilterScript)})

	var itemCounter uint32 = 0

	if in.EnablePaging || in.EnableDefaultLimit {
		countResult := obj.ItemsCounterDB(ItemsCounterIn{Script: counterFilterScript, Trace: in.Trace, Label: in.Label})

		if !countResult.Success {
			message := fmt.Sprintf("ItemsDB: CountResult not success. Script: %v", counterFilterScript)
			errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
				Message: lang.Message{
					Message: message,
				},
			})
		}

		itemCounter = *countResult.ItemsCounter
	}

	result := obj.SelectScript(SelectScriptIn{Script: mainScript, Trace: in.Trace, Dest: in.Dest})

	if result.Error != nil {
		message := fmt.Sprintf("ItemsDB: SelectScript not success %v. Script: %v", result.Error.Error(), mainScript)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	if itemCounter == 0 {
		itemCounter = result.DataLen
	}

	if itemCounter == 0 || destValue.Elem().Len() == 0 {
		in.Dest = []interface{}{}
	}

	FiltersAllowed := []interface{}{}
	OrdersAllowed := []interface{}{}

	for _, value := range in.FiltersAllowed {
		if value.Public != nil {
			FiltersAllowed = append(FiltersAllowed, value.Public)
		}
	}

	for _, value := range in.OrdersAllowed {
		if value.Public != nil {
			OrdersAllowed = append(OrdersAllowed, value.Public)
		}
	}

	return ItemsOut{
		Success: true,
		Applied: map[string]interface{}{
			"filtering": in.FilterVals,
			"ordering":  in.OrdersVals,
			"paging": map[string]interface{}{
				"size":  in.PagingSize,
				"index": in.PagingIndex,
			},
		},
		Items:          in.Dest,
		HasFilter:      filterScript != "",
		HasPaging:      in.EnablePaging,
		ItemsCounter:   itemCounter,
		PagingSize:     in.PagingSize,
		DefaultLimit:   in.DefaultLimit,
		FiltersAllowed: FiltersAllowed,
		OrdersAllowed:  OrdersAllowed,
	}
}

func (obj MySQL) ItemsCounterDB(in ItemsCounterIn) ItemsCounterOut {
	obj.log.Debug(logger.LogInfo{Key: "ItemsCounterDB", Value: fmt.Sprintf("%s %s", "COUNTER SCRIPT - ", in.Script)})

	if !strings.Contains(in.Script, " count(*) AS itemsCounter") {
		message := fmt.Sprintf("ItemsCounterDB: Invalid count script %v", in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	counter := ItemsCounterOut{}
	result := obj.SelectScript(SelectScriptIn{Script: in.Script, Trace: in.Trace, Dest: &counter})

	if result.Error != nil {
		message := fmt.Sprintf("ItemsCounterDB: Error in count script %v. Script %v", result.Error.Error(), in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	counter.Success = true
	return counter
}

func (obj MySQL) NewItemDB(in NewItemIn) NewItemOut {
	obj.log.Debug(logger.LogInfo{Key: "NewItemDB", Value: fmt.Sprintf("%s %s", "INSERT SCRIPT - ", in.Script)})

	result := obj.RunScript(RunScriptIn{Script: in.Script, Trace: in.Trace})

	if result.Error != nil {
		message := fmt.Sprintf("NewItemDB: Script returned %v error. Script: %s", result.Error, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	if result.AffectedRows != 1 {
		message := fmt.Sprintf("NewItemDB: Script returned %v affected rows. Script: %s", result.AffectedRows, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	return NewItemOut{Success: true, InsertId: result.LastInsertId}
}

func (obj MySQL) NewItemsDB(in NewItemsIn) NewItemsOut {
	obj.log.Debug(logger.LogInfo{Key: "NewItemsDB", Value: fmt.Sprintf("%s %s", "INSERT SCRIPT - ", in.Script)})

	result := obj.RunScript(RunScriptIn{Script: in.Script, Trace: in.Trace})

	if result.Error != nil {
		message := fmt.Sprintf("NewItemsDB: Script returned %v error. Script: %s", result.Error, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	if result.AffectedRows < 1 {
		message := fmt.Sprintf("NewItemsDB: Script returned %v affected rows. Script: %s", result.AffectedRows, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	return NewItemsOut{Success: true, AffectedRows: result.AffectedRows}
}

func (obj MySQL) UpdateItemsDB(in UpdateItemIn) UpdateItemOut {
	obj.log.Debug(logger.LogInfo{Key: "UpdateItemsDB", Value: fmt.Sprintf("%s %s", "UPDATE SCRIPT - ", in.Script)})

	affectedRows := obj.ItemsCounterDB(ItemsCounterIn{Script: in.AffectedRowsScript})

	if !affectedRows.Success {
		message := fmt.Sprintf("UpdateItemsDB: AffectedRows invalid. Script: %s", in.AffectedRowsScript)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	if in.EnableAffectedRows && in.ExpectedAffectedRows != *affectedRows.ItemsCounter {
		message := fmt.Sprintf("UpdateItemsDB: Affected rows must be %v and items found are %v. Script: %v", in.ExpectedAffectedRows, *affectedRows.ItemsCounter, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	result := obj.RunScript(RunScriptIn{Script: in.Script})

	if result.Error != nil {
		message := fmt.Sprintf("UpdateItemsDB: Update script error %v. Script: %v", result.Error.Error(), in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	return UpdateItemOut{Success: true, AffectedRows: result.AffectedRows}
}

func (obj MySQL) UpdateItemDB(in UpdateItemIn) UpdateItemOut {
	return obj.UpdateItemsDB(UpdateItemIn{
		Script:               in.Script,
		AffectedRowsScript:   in.AffectedRowsScript,
		Trace:                in.Trace,
		CheckDuplicated:      in.CheckDuplicated,
		EnableAffectedRows:   true,
		ExpectedAffectedRows: 1,
	})
}

func (obj MySQL) RemoveItemsDB(in RemoveItemIn) RemoveItemOut {
	obj.log.Debug(logger.LogInfo{Key: "RemoveItemsDB", Value: fmt.Sprintf("%s %s", "REMOVE SCRIPT - ", in.Script)})

	affectedRows := obj.ItemsCounterDB(ItemsCounterIn{Script: in.AffectedRowsScript})

	if !affectedRows.Success {
		message := fmt.Sprintf("RemoveItemsDB: AffectedRows invalid. Script: %s", in.AffectedRowsScript)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	if in.EnableAffectedRows && in.ExpectedAffectedRows != *affectedRows.ItemsCounter {
		message := fmt.Sprintf("RemoveItemsDB: Affected rows must be %v and items found are %v. Script: %v", in.ExpectedAffectedRows, *affectedRows.ItemsCounter, in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	result := obj.RunScript(RunScriptIn{Script: in.Script})

	if result.Error != nil {
		message := fmt.Sprintf("RemoveItemsDB: Update script error %v. Script: %v", result.Error.Error(), in.Script)
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: message,
			},
		})
	}

	return RemoveItemOut{Success: true}
}

func (obj MySQL) RemoveItemDB(in RemoveItemIn) RemoveItemOut {
	return obj.RemoveItemsDB(RemoveItemIn{
		Script:               in.Script,
		AffectedRowsScript:   in.AffectedRowsScript,
		Trace:                in.Trace,
		CheckDuplicated:      in.CheckDuplicated,
		EnableAffectedRows:   true,
		ExpectedAffectedRows: 1,
	})
}

func (obj *MySQL) CheckDBHealth(in HealthIn) HealthOut {
	var health HealthData

	ret := obj.SelectScript(SelectScriptIn{Script: "Select now() As now;", Trace: in.Trace, Dest: &health})

	if ret.Error != nil {
		return HealthOut{Success: false, Message: ret.Error.Error()}
	}

	return HealthOut{
		Success: ret.Error == nil,
		Data:    health,
		Micro:   in.Micro,
	}
}
