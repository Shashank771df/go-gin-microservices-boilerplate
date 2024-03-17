package users

type SearchItemInterSVC struct {
	UserId     *uint32 `json:"id"`
	Email      *string `json:"email"`
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	IcCode     *string `json:"icCode"`
	Phone      *string `json:"phone"`
	Ip         *string `json:"ip"`
	BirthDate  *string `json:"birthDate"`
	BirthYear  *uint16 `json:"birthYear"`
	BirthMonth *uint8  `json:"birthMonth"`
	CountryId  *string `json:"countryId"`
	Address    *string `json:"address"`
	DocIdType  *uint8  `json:"docIdType"`
	DocId      *string `json:"docId"`
	Completed  *uint8  `json:"completed"`
	Active     *uint8  `json:"active"`
	Tester     *uint8  `json:"tester"`
}
