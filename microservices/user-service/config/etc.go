package config

// Etc variables globales de la aplicacion
var Etc = etc{}

func init() {
	Etc = etc{
		EnvKeys: EnvKeys{
			Alert:    "alert",
			Database: "database",
			Guard:    "guard",
			SdkBase:  "sdkBase",
			Log:      "log",
		},
		ActiveStatus: ActivateStatus{
			Active:   1,
			Inactive: 0,
		},
	}
}
