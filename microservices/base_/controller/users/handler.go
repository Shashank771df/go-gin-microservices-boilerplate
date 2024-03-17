package users

import (
	"app/core"
	"app/core/errors"
	"app/core/lang"
	"app/core/network"
	"app/core/utils"
	"app/microservices/base_/config"
	pipes "app/microservices/base_/pipes/users"
	repo "app/microservices/base_/repository/users"
	"net/http"
)

type Handler struct {
}

func (o Handler) NewItem(c core.AppContext) {
	var dataHandler repo.Data

	params, ok := c.Data.(*pipes.NewItem)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	//Validations
	itemDB := dataHandler.ValidateNewItemDB(repo.ValidateNewItemDB{
		Email:     params.Email,
		IcCode:    params.IcCode,
		Phone:     params.Phone,
		DocIdType: params.DocIdType,
		DocId:     params.DocId,
	})

	if itemDB.Success && itemDB.ItemsCounter >= 1 {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Base.USER_ALREADY_EXISTS,
			Data: map[string]interface{}{
				"itemDB": itemDB,
				"params": *params,
			},
			HideData: true,
		})
	}

	//Response
	c.JSON(http.StatusOK, dataHandler.NewItemDB(repo.NewItemDB{
		Email:      params.Email,
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		IcCode:     params.IcCode,
		Phone:      params.Phone,
		Ip:         network.UtilNetwork{}.RemoteIPAddress(c),
		BirthDate:  params.BirthDate,
		BirthYear:  params.BirthYear,
		BirthMonth: params.BirthMonth,
		CountryId:  params.CountryId,
		Address:    params.Address,
		DocIdType:  params.DocIdType,
		DocId:      params.DocId,
		RegUserId:  &c.Sec.Info.UserID,
	}))
}

func (o Handler) Item(c core.AppContext) {
	params, ok := c.Data.(*pipes.Item)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var dataHandler repo.Data

	c.JSON(http.StatusOK, dataHandler.ItemDB(repo.ItemDB{
		UserId: params.UserId,
	}))
}

func (o Handler) SearchItemInterSVC(c core.AppContext) {
	params, ok := c.Data.(*pipes.SearchItemInterSVC)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var dataHandler repo.Data

	userItemDB := dataHandler.ItemDB(repo.ItemDB{
		UserId: params.UserId,
	})

	if !userItemDB.ItemFound {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Base.USER_NOT_EXISTS,
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var userItem SearchItemInterSVC
	utils.UtilsMap{}.InterfaceToStruct(userItemDB.Item, &userItem)
	userItemDB.Item = userItem

	c.JSON(http.StatusOK, userItemDB)
}

func (o Handler) Items(c core.AppContext) {
	params, ok := c.Data.(*pipes.Items)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var dataHandler repo.Data

	c.JSON(http.StatusOK, dataHandler.ItemsDB(repo.ItemsDB{
		OrderVals:    params.Orders,
		FilterVals:   params.Filters,
		EnablePaging: params.EnablePaging,
		PagingSize:   params.PagingSize,
		PagingIndex:  params.PagingIndex,
	}))
}

func (o Handler) UpdateItem(c core.AppContext) {
	params, ok := c.Data.(*pipes.UpdateItem)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var dataHandler repo.Data

	//Validations
	validateDB := dataHandler.ValidateUpdateItemDB(repo.ValidateUpdateItemDB{
		UserId:    params.UserId,
		Email:     params.Email,
		IcCode:    params.IcCode,
		Phone:     params.Phone,
		DocIdType: params.DocIdType,
		DocId:     params.DocId,
	})

	if validateDB.Success && validateDB.ItemsCounter >= 1 {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Base.USER_INFORMATION_IN_USE,
			Data: map[string]interface{}{
				"validateDB": validateDB,
				"params":     *params,
			},
			HideData: true,
		})
	}

	itemDB := dataHandler.ItemDB(repo.ItemDB{
		UserId: params.UserId,
	})

	if !itemDB.ItemFound {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Base.USER_NOT_EXISTS,
			Data: map[string]interface{}{
				"itemDB": itemDB,
				"params": *params,
			},
			HideData: true,
		})
	}

	//Using UpdateItemDB
	c.JSON(http.StatusOK, dataHandler.UpdateItemDB(repo.UpdateItemDB{
		Email:      params.Email,
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		IcCode:     params.IcCode,
		Phone:      params.Phone,
		BirthDate:  params.BirthDate,
		BirthYear:  params.BirthYear,
		BirthMonth: params.BirthMonth,
		CountryId:  params.CountryId,
		Address:    params.Address,
		DocIdType:  params.DocIdType,
		DocId:      params.DocId,
		UserId:     params.UserId,
		UpdUserId:  c.Sec.Info.UserID,
	}))
}

func (o Handler) RemoveItem(c core.AppContext) {
	params, ok := c.Data.(*pipes.RemoveItem)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var dataHandler repo.Data

	//Validations
	itemDB := dataHandler.ItemDB(repo.ItemDB{
		UserId: params.UserId,
	})

	if !itemDB.ItemFound {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Base.USER_NOT_EXISTS,
			Data: map[string]interface{}{
				"itemDB": itemDB,
				"params": *params,
			},
			HideData: true,
		})
	}

	//Using RemoveItemDB
	c.JSON(http.StatusOK, dataHandler.RemoveItemDB(repo.RemoveItemDB{
		UserId:    params.UserId,
		RmvUserId: c.Sec.Info.UserID,
	}))
}

func (o Handler) ActivateItem(c core.AppContext) {
	params, ok := c.Data.(*pipes.Activate)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var dataHandler repo.Data

	//Validations
	itemDB := dataHandler.ItemDB(repo.ItemDB{
		UserId: params.UserId,
	})

	if !itemDB.ItemFound {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Base.USER_NOT_EXISTS,
			Data: map[string]interface{}{
				"itemDB": itemDB,
				"params": *params,
			},
			HideData: true,
		})
	}

	//Using UpdateItemDB
	c.JSON(http.StatusOK, dataHandler.UpdateItemDB(repo.UpdateItemDB{
		UserId:    params.UserId,
		Active:    &config.Etc.ActiveStatus.Active,
		UpdUserId: c.Sec.Info.UserID,
	}))
}

func (o Handler) DeactivateItem(c core.AppContext) {
	params, ok := c.Data.(*pipes.Deactivate)
	if !ok {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "Invalid pipe data",
			},
			Data: map[string]interface{}{
				"data": c.Data,
			},
			HideData: true,
		})
	}

	var dataHandler repo.Data

	//Validations
	itemDB := dataHandler.ItemDB(repo.ItemDB{
		UserId: params.UserId,
	})

	if !itemDB.ItemFound {
		errors.HTTPErrors.Conflict(errors.HTTPErrorConfig{
			Message: lang.Base.USER_NOT_EXISTS,
			Data: map[string]interface{}{
				"itemDB": itemDB,
				"params": *params,
			},
			HideData: true,
		})
	}

	//Using NewUserDB
	c.JSON(http.StatusOK, dataHandler.UpdateItemDB(repo.UpdateItemDB{
		UserId:    params.UserId,
		Active:    &config.Etc.ActiveStatus.Inactive,
		UpdUserId: c.Sec.Info.UserID,
	}))
}

func (o Handler) UsersGroupByMonth(c core.AppContext) {
	var dataHandler repo.Data

	c.JSON(http.StatusOK, dataHandler.ItemsGroupByMonthDB())
}
