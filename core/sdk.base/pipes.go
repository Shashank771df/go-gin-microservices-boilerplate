package sdkbase

import "app/core/network"

type SearchDataIn struct {
	UserId uint32 `json:"userId"`
}

type SearchItemIn struct {
	UserId     uint32 `json:"userId" validate:"required,min=1"`
	ThrowError bool
	Trace      string
}

type SearchDataOut struct {
	UserId       uint32 `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	IcCode       string `json:"icCode"`
	Phone        string `json:"phone"`
	Ip           string `json:"ip"`
	BirthDate    string `json:"birthDate"`
	BirthYear    uint16 `json:"birthYear"`
	BirthMonth   uint8  `json:"birthMonth"`
	CountryId    string `json:"countryId"`
	Address      string `json:"address"`
	DocIdType    uint8  `json:"docIdType"`
	DocId        string `json:"docId"`
	Completed    uint8  `json:"completed"`
	Active       uint8  `json:"active"`
	Tester       uint8  `json:"tester"`
	RegUserId    uint32 `json:"regUserId"`
	RegDate      string `json:"regDate"`
	RegDatetime  string `json:"regDatetime"`
	RegTimestamp int64  `json:"regTimestamp"`
}

type SearchItemOut struct {
	IsOk              bool
	Data              SearchDataOut
	ErrUserNotFound   bool
	ErrInternalServer bool
	Result            network.HttpClientResponse
}
