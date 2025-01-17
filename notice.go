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
	Intent *AndroidIntent `json:"intent,omitempty"`
	/*
		该字段用于指定开发者想要打开的 activity，值为 activity 节点的 “android:name”属性值;
		适配华为、小米、vivo厂商通道跳转；
		Jpush SDK≥V4.2.2，可不再填写本字段，仅设置intent字段即可
	*/
	UriActivity string `json:"uri_activity,omitempty"`
	/*
		该字段用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值;
		适配 oppo、fcm跳转；
		Jpush SDK≥V4.2.2，可不再填写本字段，仅设置intent字段即可，但若需兼容旧版SDK必须填写该字段
	*/
	UriAction string `json:"uri_action,omitempty"`
	/*
		角标数字，取值范围1-99
		此属性目前仅针对华为 EMUI 8.0 及以上、小米 MIUI 6 及以上设备生效；
		此字段如果不填，表示不改变角标数字（小米设备由于系统控制，不论推送走极光通道下发还是厂商通道下发，即使不传递依旧是默认+1的效果。）；
		否则下一条通知栏消息配置的badge_add_num数据会和之前角标数量进行增加； 建议badge_add_num配置为1；
		举例：badge_add_num配置1，应用之前角标数为2，发送此角标消息后，应用角标数显示为3。
	*/
	BadgeAddNum int `json:"badge_add_num,omitempty"`
	/*
		值为 "1" 时，APP 在前台会弹出通知栏消息；
		值为 "0" 时，APP 在前台不会弹出通知栏消息。
		注：默认情况下 APP 在前台会弹出通知栏消息。
	*/
	/*
		桌面图标对应的应用入口Activity类， 比如“com.test.badge.MainActivity”
		配合badge_add_num使用，二者需要共存，缺少其一不可；
		针对华为设备推送时生效（此值如果填写非主Activity类，走厂商推送以华为厂商限制逻辑为准；走极光通道下发，默认以APP的主Activity为准）
	*/
	BadgeClass        string                 `json:"badge_class,omitempty"`
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
