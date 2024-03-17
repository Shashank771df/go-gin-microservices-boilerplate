package config

type (
	// Env variables de entorno de la aplicacion
	env struct {
		Micro Micro
	}

	// Micro .
	Micro struct {
		API         API
		Alert       Alert
		APIServices APIServices
		App         App
		DB          DB
		Log         Log
	}

	// App .
	App struct {
		Stage         string
		APIKey        string
		Address       string
		Port          uint16
		Protocol      string
		Name          string
		ServerTimeout uint16
		ExitTimeout   uint16
	}

	// API .
	API struct {
		WhiteList []string
		CountryId string
		DomainId  string
	}

	Alert struct {
		Enable bool
		Agent  string
		UrlBot string
		ChatId string
	}

	APIParams struct {
		Host string
	}

	APIServices struct {
		Base             APIParams
		SecurityAdmin    APIParams
		SecurityCustomer APIParams
	}

	// Files .
	Files struct {
		HostPath   string
		AppHostURL string
	}

	// DB .
	DB struct {
		Mysql MysqlEnv
		Mongo MongoEnv
		Redis RedisEnv
	}

	// MysqlEnv .
	MysqlEnv struct {
		Host             string
		Port             uint16
		Name             string
		User             string
		Pass             string
		MaxOpenConns     uint32
		MaxIdleConns     uint32
		MaxLifeTimeConns uint32
	}

	// MongoEnv .
	MongoEnv struct {
		Host string
		Port uint16
		Name string
		User string
		Pass string
	}

	// RedisEnv .
	RedisEnv struct {
		Host   string
		Port   uint16
		Prefix string
	}

	// Interservices .
	Interservices struct {
		Authorization Service
	}

	// ExternalServices .
	ExternalServices struct {
		Github Service
	}

	// Service Plantilla para crear un external service
	Service struct {
		// var1      string
		// endpoint1 string
	}

	//Logs
	Log struct {
		Type           string
		ProjectId      string
		Name           string
		MaxFiles       uint16
		FilePrefix     string
		FileSize       string
		Path           string
		ConsoleEnabled bool
		FileEnabled    bool
		MaxQueue       uint16
		LogLevel       string
	}
)
