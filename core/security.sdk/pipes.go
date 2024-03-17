package securitysdk

type GuardOut struct {
	IsOk  bool
	Error string
	Data  interface{}
}

type CheckAPIKeyIn struct {
	APIKey string
	Trace  string
	Check  bool
}

type WhiteListIn struct {
	WhiteList []string
	Trace     string
	Check     bool
}

type CheckUserSessionIn struct {
	SecurityUrl string
	Trace       string
	Check       bool
	ReportError bool
}

type CheckUserPrivilegeIn struct {
	SecurityUrl string
	Trace       string
	EndpointId  string
	Check       bool
	ReportError bool
}

type CheckUserIsAuthorizedIn struct {
	SecurityUrl string
	Trace       string
	EndpointId  string
	Check       bool
	ReportError bool
}
