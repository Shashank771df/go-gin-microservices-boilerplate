package database

// Patterns
const (
	EqualPattern              string = `%v = '%v'`
	NotEqualPattern           string = `%v != '%v'`
	LikePattern               string = `%v LIKE '%%%v%%'`
	BetweenPattern            string = `%v BETWEEN '%v' AND '%v'`
	InPattern                 string = `%v IN (%v)`
	GreaterThanPattern        string = `%v > '%v'`
	GreaterThanOrEqualPattern string = `%v >= '%v'`
	LessThanPattern           string = `%v < '%v'`
	LessThanOrEqualPattern    string = `%v <= '%v'`
)

// RunScript
type (
	RunScriptIn struct {
		Script string
		Trace  string
	}
	RunScriptOut struct {
		Error        error `json:"error"`
		AffectedRows int32 `json:"affectedRows"`
		LastInsertId int64 `json:"lastInsertId"`
		ChangedRows  int32 `json:"changedRows"`
	}
)

// SelectScript
type (
	SelectScriptIn struct {
		Script string
		Trace  string
		Dest   interface{}
	}
	SelectScriptOut struct {
		Error   error  `json:"error"`
		DataLen uint32 `json:"dataLen"`
	}
)

// ItemDB
type (
	ItemIn struct {
		Script          string
		CheckDuplicated bool
		Dest            interface{}
		Trace           string
		Label           string
	}
	ItemOut struct {
		Success   bool        `json:"success"`
		ItemFound bool        `json:"itemFound"`
		Message   string      `json:"message,omitempty"`
		Item      interface{} `json:"item,omitempty"`
	}
)

// ItemsCounterDB
type (
	ItemsCounterIn struct {
		Script string
		Trace  string
		Label  string
	}
	ItemsCounterOut struct {
		Success      bool    `json:"success"`
		Message      string  `json:"message"`
		ItemsCounter *uint32 `json:"itemsCounter"`
	}
)

// ItemsDB
type (
	FilterAllowed struct {
		Column  string                 `json:"-"`
		Order   string                 `json:"-"`
		Pattern string                 `json:"-"`
		Public  map[string]interface{} `json:"public,omitempty"`
	}

	OrdersAllowed struct {
		Column string                 `json:"-"`
		Order  string                 `json:"-"`
		Public map[string]interface{} `json:"public,omitempty"`
	}

	ItemsIn struct {
		SelectScript       string
		WhereScript        string
		CounterScript      string
		FiltersAllowed     map[string]FilterAllowed
		FilterVals         string
		OrdersAllowed      map[string]OrdersAllowed
		OrdersVals         string
		DefaultOrderBy     string
		EnablePaging       bool
		PagingSize         uint32
		PagingIndex        uint32
		EnableDefaultLimit bool
		DefaultLimit       uint32
		Dest               interface{}
		Trace              string
		Label              string
	}
	ItemsOut struct {
		Success        bool          `json:"success"`
		Message        string        `json:"message"`
		Items          interface{}   `json:"items"`
		HasFilter      bool          `json:"hasFilter"`
		HasPaging      bool          `json:"hasPaging"`
		ItemsCounter   uint32        `json:"itemsCounter"`
		PagingSize     uint32        `json:"pagingSize"`
		DefaultLimit   uint32        `json:"defaultLimit"`
		Applied        interface{}   `json:"applied,omitempty"`
		FiltersAllowed []interface{} `json:"filtersAllowed,omitempty"`
		OrdersAllowed  []interface{} `json:"ordersAllowed,omitempty"`
	}
)

// NewItemDB
type (
	NewItemIn struct {
		Script          string
		CheckDuplicated bool
		Dest            interface{}
		Trace           string
		Label           string
	}
	NewItemOut struct {
		Success  bool   `json:"success"`
		Message  string `json:"message"`
		InsertId int64  `json:"insertId"`
	}
)

// NewItemsDB
type (
	NewItemsIn struct {
		Script          string
		CheckDuplicated bool
		Dest            interface{}
		Trace           string
		Label           string
	}
	NewItemsOut struct {
		Success      bool   `json:"success"`
		Message      string `json:"message"`
		AffectedRows int32  `json:"affectedRows"`
	}
)

type (
	UpdateItemIn struct {
		Script               string
		AffectedRowsScript   string
		CheckDuplicated      bool
		EnableAffectedRows   bool
		ExpectedAffectedRows uint32
		Dest                 interface{}
		Trace                string
		Label                string
	}
	UpdateItemOut struct {
		Success      bool   `json:"success"`
		AffectedRows int32  `json:"affectedRows"`
		Message      string `json:"message"`
	}
)

// RemoveItemDB
type (
	RemoveItemIn struct {
		Script               string
		AffectedRowsScript   string
		CheckDuplicated      bool
		EnableAffectedRows   bool
		ExpectedAffectedRows uint32
		Dest                 interface{}
		Trace                string
		Label                string
	}
	RemoveItemOut struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

// HealthDB
type (
	HealthIn struct {
		Micro string
		Trace string
	}
	HealthOut struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Micro   string      `json:"micro"`
	}
	HealthData struct {
		Current *string `db:"now"`
	}
)
