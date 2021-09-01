package jpushclient

type Notice struct {
	Alert    string          `json:"alert,omitempty"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
}
type AndroidIntent struct {
	//"url": "intent:#Intent;component=com.jiguang.push/com.example.jpushdemo.SettingActivity;end"
	Url string `json:"url,omitempty"`
}
type AndroidNotice struct {
	Alert string `json:"alert"`
	Title string `json:"title,omitempty"`
	/*
		Android SDK 可设置通知栏样式，这里根据样式 ID 来指定该使用哪套样式，android 8.0 开始建议采用NotificationChannel配置。
	*/
	BuilderId int `json:"builder_id,omitempty"`
	/*
		不超过1000字节，Android 8.0开始可以进行NotificationChannel配置，这里根据channel ID 来指定通知栏展示效果。
	*/
	ChannelId string `json:"channel_id,omitempty"`
	/*
		默认为 0，还有 1，2，3 可选，用来指定选择哪种通知栏样式，其他值无效。有三种可选分别为 bigText=1，Inbox=2，bigPicture=3。
	*/
	Style int `json:"style,omitempty"`
	/*
		可选范围为 -1～7 ，对应 Notification.DEFAULT_ALL = -1 或者 Notification.DEFAULT_SOUND = 1,
		Notification.DEFAULT_VIBRATE = 2, Notification.DEFAULT_LIGHTS = 4 的任意 “or” 组合。默认按照 -1 处理。
	*/
	AlertType int `json:"alert_type,omitempty"`
	/*
		使用 intent 里的 url 指定点击通知栏后跳转的目标页面;
		SDK＜422的版本此字段值仅对走华硕通道和极光自有通道下发生效，不影响请求走其它厂商通道。
		SDK≥422的版本，API推送时建议填写intent字段（intent:#Intent;component=您的包名/Activity全名;end），否则点击通知可能无跳转动作。
	*/
	Intent AndroidIntent `json:"intent,omitempty"`
	/*
		值为 "1" 时，APP 在前台会弹出通知栏消息；
		值为 "0" 时，APP 在前台不会弹出通知栏消息。
		注：默认情况下 APP 在前台会弹出通知栏消息。
	*/
	DisplayForeground string                 `json:"display_foreground,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
}

type IOSNotice struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            string                 `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
}

type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

func (notice *Notice) SetAlert(alert string) {
	notice.Alert = alert
}

func (notice *Notice) SetAndroidNotice(n *AndroidNotice) {
	notice.Android = n
}

func (notice *Notice) SetIOSNotice(n *IOSNotice) {
	notice.IOS = n
}

func (notice *Notice) SetWinPhoneNotice(n *WinPhoneNotice) {
	notice.WINPhone = n
}
