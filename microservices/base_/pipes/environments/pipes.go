package environments

type UpdateItemInterSVC struct {
	KeyId *string     `json:"keyId" validate:"required,min=1"`
	Data  interface{} `json:"data" validate:"required,min=1"`
}

type SearchItemInterSVC struct {
	KeyId *string `param:"keyId" validate:"required,min=1"`
}
