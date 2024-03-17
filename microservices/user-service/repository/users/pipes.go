package users

type Users struct {
	UserId       *uint32 `db:"userId" json:"id"`
	Email        *string `db:"email" enableDefault:"true" json:"email"`
	FirstName    *string `db:"firstName" enableDefault:"true" json:"firstName"`
	LastName     *string `db:"lastName" enableDefault:"true" json:"lastName"`
	IcCode       *string `db:"icCode" enableDefault:"true" json:"icCode"`
	Phone        *string `db:"phone" enableDefault:"true" json:"phone"`
	Ip           *string `db:"ip" json:"ip"`
	BirthDate    *string `db:"birthDate" enableDefault:"true" json:"birthDate"`
	BirthYear    *uint16 `db:"birthYear" enableDefault:"true" json:"birthYear"`
	BirthMonth   *uint8  `db:"birthMonth" enableDefault:"true" json:"birthMonth"`
	CountryId    *string `db:"countryId" enableDefault:"true" json:"countryId"`
	Address      *string `db:"address" enableDefault:"true" json:"address"`
	DocIdType    *uint8  `db:"docIdType" enableDefault:"true" json:"docIdType"`
	DocId        *string `db:"docId" json:"docId"`
	Completed    *uint8  `db:"completed" json:"completed"`
	Active       *uint8  `db:"active" json:"active"`
	Tester       *uint8  `db:"tester" json:"tester"`
	RegUserId    *uint32 `db:"regUserId" enableDefault:"true" json:"regUserId"`
	RegDate      *string `db:"regDate" enableDefault:"true" json:"regDate"`
	RegDatetime  *string `db:"regDatetime" enableDefault:"true" json:"regDatetime"`
	RegTimestamp *int64  `db:"regTimestamp" enableDefault:"true" json:"regTimestamp"`
}

type Where struct {
	UserId       *uint32 `db:"userId" pattern:"="`
	Email        *string `db:"email" pattern:"="`
	IcCode       *string `db:"icCode" pattern:"="`
	Phone        *string `db:"phone" pattern:"="`
	Ip           *string `db:"ip" pattern:"="`
	BirthDate    *string `db:"birthDate" pattern:"="`
	BirthYear    *uint16 `db:"birthYear" pattern:"="`
	BirthMonth   *uint8  `db:"birthMonth" pattern:"="`
	CountryId    *string `db:"countryId" pattern:"="`
	DocIdType    *uint8  `db:"docIdType" pattern:"="`
	DocId        *string `db:"docId" pattern:"="`
	Completed    *uint8  `db:"completed" pattern:"="`
	Active       *uint8  `db:"active" pattern:"="`
	Tester       *uint8  `db:"tester" pattern:"="`
	RegUserId    *uint32 `db:"regUserId" pattern:"="`
	RegDate      *string `db:"regDate" pattern:"="`
	RegTimestamp *int64  `db:"regTimestamp" pattern:"="`
	Alias        string  `db:"alias" pattern:"="`
}

type ItemDB struct {
	UserId    *uint32
	Email     *string
	IcCode    *string
	Phone     *string
	DocIdType *uint8
	DocId     *string
	Label     string
	Trace     string
}

type ItemsDB struct {
	IcCode       *string
	Phone        *string
	Ip           *string
	BirthDate    *string
	BirthYear    *uint16
	BirthMonth   *uint8
	CountryId    *string
	DocIdType    *uint8
	Completed    *uint8
	Active       *uint8
	Tester       *uint8
	RegUserId    *uint32
	RegDate      *string
	RegTimestamp *int64
	OrderCol     string
	Order        string
	FilterVals   string
	OrderVals    string
	EnablePaging bool
	PagingSize   uint32
	PagingIndex  uint32
	Label        string
	Trace        string
}

type NewItemDB struct {
	Email      *string
	FirstName  *string
	LastName   *string
	IcCode     *string
	Phone      *string
	Ip         string
	BirthDate  *string
	BirthYear  *uint16
	BirthMonth *uint8
	CountryId  *string
	Address    *string
	DocIdType  *uint8
	DocId      *string
	RegUserId  *uint32
	Completed  *uint8
	Active     *uint8
	Tester     *uint8
	Label      string
	Trace      string
}

type ValidateNewItemDB struct {
	Email     *string
	IcCode    *string
	Phone     *string
	DocIdType *uint8
	DocId     *string
	Label     string
	Trace     string
}

type ValidateUpdateItemDB struct {
	UserId    *uint32
	Email     *string
	IcCode    *string
	Phone     *string
	DocIdType *uint8
	DocId     *string
	Label     string
	Trace     string
}

type UpdateItemDB struct {
	UserId     *uint32
	Email      *string
	FirstName  *string
	LastName   *string
	IcCode     *string
	Phone      *string
	BirthDate  *string
	BirthYear  *uint16
	BirthMonth *uint8
	CountryId  *string
	Address    *string
	DocIdType  *uint8
	DocId      *string
	Completed  *uint8
	Active     *uint8
	Tester     *uint8
	UpdUserId  uint32
	Label      string
	Trace      string
}

type RemoveItemDB struct {
	UserId    *uint32
	RmvUserId uint32
	Label     string
	Trace     string
}

type ItemsGroupByMonthDB struct {
	BirthMonth *string `db:"birthMonth" enableDefault:"false" json:"birthMonth"`
	Users      *uint32 `db:"users" enableDefault:"false" json:"users"`
}
