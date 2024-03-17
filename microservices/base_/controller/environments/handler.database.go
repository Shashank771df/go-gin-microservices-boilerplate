package environments

import (
	"app/core/errors"
	"app/core/lang"
	"app/core/utils"
)

func (o Handler) ItemDatabase() ItemDatabase {
	return ItemDatabase{}
}

func (o Handler) UpdateDatabase(data interface{}) Update {
	var params ItemDatabase
	err := utils.UtilsMap{}.InterfaceToStruct(data, &params)

	if err != nil {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": data,
			},
			HideData: true,
		})
	}

	before := o.ItemDatabase()

	return Update{
		Before:  before,
		Current: o.ItemDatabase(),
	}
}
