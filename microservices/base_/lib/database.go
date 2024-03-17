package lib

import (
	"app/core/database"
	c "app/microservices/base_/config"
	"app/microservices/base_/log"
)

var DB database.Database = nil

func init() {
	DB, _ = database.New(database.MySQLDB, database.DBProps{
		Host:             c.Env.Micro.DB.Mysql.Host,
		Port:             c.Env.Micro.DB.Mysql.Port,
		Name:             c.Env.Micro.DB.Mysql.Name,
		User:             c.Env.Micro.DB.Mysql.User,
		Pass:             c.Env.Micro.DB.Mysql.Pass,
		MaxIdleConns:     c.Env.Micro.DB.Mysql.MaxIdleConns,
		MaxOpenConns:     c.Env.Micro.DB.Mysql.MaxOpenConns,
		MaxLifeTimeConns: c.Env.Micro.DB.Mysql.MaxLifeTimeConns,
		Log:              log.Log,
	})
}
