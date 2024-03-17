package sdkbase

import (
	"app/core/errors"
	"app/core/lang"
	"app/core/network"
	u "app/core/utils"
	"app/core/validator"
	"strconv"
	"strings"
)

type basePaths struct {
	usersSearchItem string
}

type BaseSDK struct {
	paths  basePaths
	host   string
	apiKey string
	val    *validator.Validate
}

func New(host string, apiKey string) BaseSDK {
	var item BaseSDK
	item.val = validator.New()

	if host == "" {
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "BaseSDK - host is required",
			},
		})
	}

	if apiKey == "" {
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "BaseSDK - apiKey is required",
			},
		})
	}

	item.paths.usersSearchItem = "/v1/users/search-item/interservices"
	item.apiKey = apiKey
	item.host = host

	return item
}

func (o *BaseSDK) SetHost(data string) {

}

func (o BaseSDK) UsersSearchItem(data SearchItemIn) SearchItemOut {
	err := o.val.Struct(data)
	if err != nil {
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: err.Error(),
			},
		})
	}

	var mutil u.UtilsMap
	client := network.HttpClient{}
	result := client.Call(network.HttpClientRequest{
		Method: "POST",
		Url:    o.host + strings.ReplaceAll(o.paths.usersSearchItem, "{id}", strconv.FormatUint(uint64(data.UserId), 10)),
		Headers: map[string]string{
			"trace":  data.Trace,
			"apiKey": o.apiKey,
		},
		Data: SearchDataIn{
			UserId: data.UserId,
		},
	})
	if result.IsStatusIn2xx {
		//+ success
		var value SearchDataOut
		data := mutil.GetMapValue(result.Data, "item")
		err := mutil.InterfaceToStruct(data, &value)

		if err != nil {
			errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
				Message: lang.Message{
					Message: err.Error(),
				},
			})
		}
		return SearchItemOut{
			IsOk: result.IsStatusIn2xx,
			Data: value,
		}
	}

	// + 409
	if !result.IsStatusIn2xx && mutil.GetStringValue(result.Data, "messageId") == lang.Base.USER_NOT_EXISTS.ID {
		if data.ThrowError {
			errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
				Message: lang.Base.USER_NOT_EXISTS,
			})
		}
		return SearchItemOut{
			IsOk:            result.IsStatusIn2xx,
			ErrUserNotFound: true,
		}
	}

	if data.ThrowError {
		errors.HTTPErrors.InternalServerError(errors.HTTPErrorConfig{
			Message: lang.Message{
				Message: "UsersSearchItem: General error",
			},
		})
	}
	return SearchItemOut{
		IsOk:              result.IsStatusIn2xx,
		ErrInternalServer: true,
		Result:            result,
	}
}
