package config

var AppInfo = appInfo{}

type appInfo struct {
	Name        string
	Version     string
	Description string
}

func init() {
	AppInfo = appInfo{
		Name:        "User Service",
		Version:     "v1",
		Description: "Microservicio encargado de la gestion de usuarios base",
	}
}
