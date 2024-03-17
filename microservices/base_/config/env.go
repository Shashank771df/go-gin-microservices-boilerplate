package config

import (
	"app/core/dotenv"
	"app/core/logger"
)

var de dotenv.DotEnv

// Env variables de entorno de la aplicacion
var Env = env{}

func init() {
	de.Load()

	Env = env{
		Micro: Micro{
			APIServices: APIServices{
				Base: APIParams{
					Host: de.GetString("BASE_API_HOST_URL", "http://127.0.0.1:3000", ""),
				},
				SecurityAdmin: APIParams{
					Host: de.GetString("SECURITY_ADMIN_API_HOST_URL", "http://127.0.0.1:3000", ""),
				},
				SecurityCustomer: APIParams{
					Host: de.GetString("SECURITY_CUSTOMER_API_HOST_URL", "http://127.0.0.1:3000", ""),
				},
			},
			Alert: Alert{
				Enable: de.GetBool("ALERT_ENABLE", "false", ""),
				Agent:  de.GetString("ALERT_AGENT", "TELEGRAM", ""),
				UrlBot: de.GetString("ALERT_AGENT_TELEGRAM_URL_BOT", "", ""),
				ChatId: de.GetString("ALERT_AGENT_TELEGRAM_CHAT_ID", "", ""),
			},
			API: API{
				WhiteList: de.GetStringArray("WALLET_WHITELIST", ",", ""),
				CountryId: de.GetString("COUNTRY_ID", "PER", ""),
				DomainId:  de.GetString("DOMAIN_ID", "basedomain", ""),
			},
			App: App{
				Stage:         de.GetString("APP_STAGE", "DEV", ""),
				APIKey:        de.GetString("APP_API_KEY", "LALALALA", ""),
				Address:       de.GetString("APP_ADDRESS", "0.0.0.0", ""),
				Protocol:      de.GetString("APP_PROTOCOL", "http", ""),
				Name:          de.GetString("APP_NAME", "BASE", ""),
				Port:          de.GetUint16("APP_PORT", "3000", ""),
				ServerTimeout: de.GetUint16("APP_SERVER_TIMEOUT", "4", ""),
				ExitTimeout:   de.GetUint16("APP_EXIT_TIMEOUT", "1", ""),
			},
			DB: DB{
				Mysql: MysqlEnv{
					Host:             de.GetString("DB_MYSQL_HOST", "127.0.0.1", ""),
					Port:             de.GetUint16("DB_MYSQL_PORT", "3306", ""),
					Name:             de.GetString("DB_MYSQL_NAME", "DB", ""),
					User:             de.GetString("DB_MYSQL_USER", "root", ""),
					Pass:             de.GetString("DB_MYSQL_PASS", "LALALALA", ""),
					MaxOpenConns:     de.GetUint("DB_MYSQL_MAX_OPEN_CONNS", "10", ""),
					MaxIdleConns:     de.GetUint("DB_MYSQL_MAX_IDLE_CONNS", "10", ""),
					MaxLifeTimeConns: de.GetUint("DB_MYSQL_MAX_LIFETIME_CONNS", "60", ""),
				},
				Mongo: MongoEnv{
					Host: de.GetString("DB_MONGO_HOST", "127.0.0.1", ""),
					Port: de.GetUint16("DB_MONGO_PORT", "3306", ""),
					Name: de.GetString("DB_MONGO_NAME", "DB", ""),
					User: de.GetString("DB_MONGO_USER", "root", ""),
					Pass: de.GetString("DB_MONGO_PASS", "LALALALA", ""),
				},
				Redis: RedisEnv{
					Host:   de.GetString("DB_REDIS_HOST", "127.0.0.1", ""),
					Port:   de.GetUint16("DB_REDIS_PORT", "3306", ""),
					Prefix: de.GetString("DB_REDIS_PREFIX", "", ""),
				},
			},
			Log: Log{
				Type:           de.GetString("LOG_TYPE", logger.DEFAULT_LOGGER, ""),
				ProjectId:      de.GetString("LOG_GCP_PROJECT_ID", "", ""),
				Name:           de.GetString("LOG_GCP_NAME", "", ""),
				MaxFiles:       de.GetUint16("LOG_ASYNC_MAX_FILES", "2", ""),
				FilePrefix:     de.GetString("LOG_ASYNC_FILE_PREFIX", "log_", ""),
				FileSize:       de.GetString("LOG_ASYNC_FILE_SIZE", "100 KB", ""),
				Path:           de.GetString("LOG_ASYNC_PATH", "", ""),
				ConsoleEnabled: de.GetBool("LOG_ASYNC_CONSOLE_ENABLED", "false", ""),
				FileEnabled:    de.GetBool("LOG_ASYNC_FILE_ENABLED", "false", ""),
				MaxQueue:       de.GetUint16("LOG_ASYNC_MAX_QUEUE", "20", ""),
				LogLevel:       de.GetString("LOG_LEVEL", logger.ERROR, ""),
			},
		},
	}
}
