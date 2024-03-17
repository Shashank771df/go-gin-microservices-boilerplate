package database

import (
	"app/core/logger"
	"database/sql"
	"fmt"
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQLDB      Driver = "MYSQL"
	MongoDB      Driver = "MONGO"
	SQLiteDB     Driver = "SQLITE"
	PostgreSQLDB Driver = "POSTGRES"
)

type DBProps struct {
	Host             string
	Port             uint16
	Name             string
	User             string
	Pass             string
	MaxOpenConns     uint32
	MaxIdleConns     uint32
	MaxLifeTimeConns uint32
	Log              logger.Logger
}

// ClientDB interfaz cliente de Base de datos
type Database interface {
	initialize(mc DBProps)
	SetLogger(logger.Logger)
	GetConnection() *sql.DB
	CheckDBHealth(in HealthIn) HealthOut
	NewItemDB(in NewItemIn) NewItemOut
	NewItemsDB(in NewItemsIn) NewItemsOut
	UpdateItemsDB(in UpdateItemIn) UpdateItemOut
	UpdateItemDB(in UpdateItemIn) UpdateItemOut
	RemoveItemsDB(in RemoveItemIn) RemoveItemOut
	RemoveItemDB(in RemoveItemIn) RemoveItemOut
	ItemsCounterDB(in ItemsCounterIn) ItemsCounterOut
	ItemDB(in ItemIn) ItemOut
	ItemsDB(in ItemsIn) ItemsOut
	RunScript(in RunScriptIn) RunScriptOut
	SelectScript(in SelectScriptIn) SelectScriptOut
}

// InstanceClientDB create the connection with db
func New(d Driver, props DBProps) (Database, error) {
	var item Database

	switch d {
	case MySQLDB:
		item = &MySQL{}
	default:
		return nil, fmt.Errorf("driver not implemented")
	}

	item.initialize(props)

	return item, nil
}
