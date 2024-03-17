package config

type ActivateStatus struct {
	Active   uint8
	Inactive uint8
}

type EnvKeys struct {
	Alert    string
	Database string
	Guard    string
	SdkBase  string
	Log      string
}

type etc struct {
	ActiveStatus ActivateStatus
	EnvKeys      EnvKeys
}
