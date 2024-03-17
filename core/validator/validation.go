package validator

type FilterJoi struct {
	Key   string `json:"key" validate:"required,min=1"`
	Val   string `json:"val" validate:"required,min=1"`
	Val1  string `json:"val1" validate:"omitempty,min=1"`
	Order string `json:"order" validate:"omitempty,oneof=asc desc"`
}

type OrderJoi struct {
	Key   string `json:"key" validate:"required,min=1"`
	Order string `json:"order" validate:"omitempty,oneof=asc desc"`
}

type HeaderCredentialsFiltersPaging struct {
	UserId       uint32 `header:"userId" validate:"required,min=1" json:"-"`
	UserToken    string `header:"userToken" validate:"required,min=1" json:"-"`
	EnablePaging bool   `header:"enablePaging" validate:"omitempty" json:"-"`
	PagingSize   uint32 `header:"pagingSize" validate:"omitempty,min=5,max=100" json:"-"`
	PagingIndex  uint32 `header:"pagingIndex" validate:"omitempty,min=1" json:"-"`
	Filters      string `header:"filters" validate:"omitempty" json:"-"`
}

type HeaderCredentials struct {
	UserId    string `header:"userId" validate:"required,min=1" json:"-"`
	UserToken string `header:"userToken" validate:"required,min=1" json:"-"`
}

type HeaderCredentialsFiltersPagingOrders struct {
	UserId       uint32 `header:"userId" validate:"required,min=1" json:"-"`
	UserToken    string `header:"userToken" validate:"required,min=1" json:"-"`
	EnablePaging bool   `header:"enablePaging" validate:"omitempty" json:"-"`
	PagingSize   uint32 `header:"pagingSize" validate:"omitempty,min=2,max=100" json:"-"`
	PagingIndex  uint32 `header:"pagingIndex" validate:"omitempty,min=1" json:"-"`
	Filters      string `header:"filters" validate:"omitempty" json:"-"`
	Orders       string `header:"orders" validate:"omitempty" json:"-"`
}

type HeaderFiltersPagingOrders struct {
	EnablePaging bool   `header:"enablePaging" validate:"omitempty" json:"-"`
	PagingSize   uint32 `header:"pagingSize" validate:"omitempty,min=2,max=100" json:"-"`
	PagingIndex  uint32 `header:"pagingIndex" validate:"omitempty,min=1" json:"-"`
	Filters      string `header:"filters" validate:"omitempty" json:"-"`
	Orders       string `header:"orders" validate:"omitempty" json:"-"`
}

type HeaderApiKeyFiltersPaging struct {
	ApiKey  string `header:"apiKey" validate:"-" json:"-"`
	Filters string `header:"filters" validate:"-" json:"-"`
}
