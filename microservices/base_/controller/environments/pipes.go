package environments

type (
	Items struct {
		Alert    ItemAlert    `json:"alert"`
		Database ItemDatabase `json:"database"`
		Guard    ItemGuard    `json:"guard"`
		SdkBase  ItemSdkBase  `json:"sdkBase"`
		Log      ItemLog      `json:"log"`
	}
	Item struct {
		KeyId string      `json:"keyId"`
		Item  interface{} `json:"item"`
	}

	Update struct {
		Current interface{} `json:"current"`
		Before  interface{} `json:"before"`
	}
)

type (
	ItemAlert struct {
		Enable *bool   `json:"enable"`
		Url    *string `json:"url"`
		ChatId *string `json:"chatId"`
	}

	ItemDatabase struct {
	}

	ItemGuard struct {
		WhiteList *[]string `json:"whiteList"`
	}

	ItemSdkBase struct {
		Host *string `json:"host"`
	}

	ItemLog struct {
		LogLevel *string `json:"logLevel"`
	}
)
