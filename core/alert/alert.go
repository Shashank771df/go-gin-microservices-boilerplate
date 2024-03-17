package alert

const (
	TELEGRAM = "TELEGRAM"
)

type AlertInfo struct{
	Key		string
	Value		string
	Extra		interface{}
}

type AlertProps struct{
	Url 		string
	ChatId		string
	ParseMode	string
	Enable		bool
}

type Alert interface{
	SendMessage(log AlertInfo) error
	Initialize(log AlertProps) error
}

func New(logger string, props AlertProps) Alert{
	var item Alert

	switch (logger){
		case TELEGRAM:
			item = &TelegramSDK{}
	}

	item.Initialize(props)

	return item
}