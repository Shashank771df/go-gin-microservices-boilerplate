package users

import "app/core/validator"

type Item struct {
	UserId *uint32 `param:"id" validate:"required,min=1"`
}

type Items struct {
	validator.HeaderCredentialsFiltersPagingOrders
}

type NewItem struct {
	Email      *string `json:"email" validate:"omitempty,min=1,max=250"`
	FirstName  *string `json:"firstName" validate:"required,min=1,max=100"`
	LastName   *string `json:"lastName" validate:"required,min=1,max=100"`
	IcCode     *string `json:"icCode" validate:"omitempty,min=1,max=10"`
	Phone      *string `json:"phone" validate:"omitempty,min=1,max=20"`
	BirthDate  *string `json:"birthDate" validate:"omitempty"`
	BirthYear  *uint16 `json:"birthYear" validate:"omitempty,min=1900,max=3000"`
	BirthMonth *uint8  `json:"birthMonth" validate:"omitempty,min=1,max=12"`
	CountryId  *string `json:"countryId" validate:"omitempty,min=3,max=3"`
	Address    *string `json:"address" validate:"omitempty,min=1,max=4096"`
	DocIdType  *uint8  `json:"docIdType" validate:"required,min=1,max=20"`
	DocId      *string `json:"docId" validate:"required,min=1,max=50"`
}

type UpdateItem struct {
	UserId     *uint32 `param:"id" validate:"required,min=1"`
	Email      *string `json:"email" validate:"omitempty,min=1,max=250"`
	FirstName  *string `json:"firstName" validate:"omitempty,min=1,max=100"`
	LastName   *string `json:"lastName" validate:"omitempty,min=1,max=100"`
	IcCode     *string `json:"icCode" validate:"omitempty,min=1,max=10"`
	Phone      *string `json:"phone" validate:"omitempty,min=1,max=20"`
	BirthDate  *string `json:"birthDate" validate:"omitempty"`
	BirthYear  *uint16 `json:"birthYear" validate:"omitempty,min=1900,max=3000"`
	BirthMonth *uint8  `json:"birthMonth" validate:"omitempty,min=1,max=12"`
	CountryId  *string `json:"countryId" validate:"omitempty,min=3,max=3"`
	Address    *string `json:"address" validate:"omitempty,min=1,max=4096"`
	DocIdType  *uint8  `json:"docIdType" validate:"omitempty,min=1,max=20"`
	DocId      *string `json:"docId" validate:"omitempty,min=1,max=50"`
}

type RemoveItem struct {
	UserId *uint32 `param:"id" validate:"required,min=1"`
}

type Activate struct {
	UserId *uint32 `param:"id" validate:"required,min=1,max=30"`
}

type Deactivate struct {
	UserId *uint32 `param:"id" validate:"required,min=1,max=30"`
}

type SearchItemInterSVC struct {
	UserId *uint32 `param:"userId" validate:"required,min=1"`
}
