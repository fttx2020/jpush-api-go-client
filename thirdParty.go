package jpushclient

/*
	"third_party_channel":{
	    "xiaomi":{
	              "distribution":"jpush",
	              "channel_id":"*******",
	              "large_icon":"jgmedia-2-14b23451-0001-41ce-89d9-987b465122da",
	              "small_icon_uri":"jgmedia-3-14b23451-0001-41ce-89d9-987b465122da",
	              "small_icon_color":"#ABCDEF",
	              "big_text":"testbigtext", // 可选，最多支持 128 个字符，配合小米 style 使用
	              "style":1,
	              "distribution_fcm":"fcm",
	              "skip_quota": true
	    },
	    "huawei":{
	              "distribution":"secondary_push",
	              "distribution_fcm":"jpush",
	              "distribution_customize":"first_ospush",
	              "sound":"/raw/shake",
	              "default_sound":false,
	              "importance":"NORMAL",
	              "large_icon":"jgmedia-2-14b23451-0001-41ce-89d9-987b465122da",
	              "small_icon_uri":"jgmedia-3-14b23451-0001-41ce-89d9-987b465122da",
	              "inbox": JSONObject, //可选，配合华为 style 使用
	              "style":2,
	              "only_use_vendor_style":true
	    },
	    "honor":{
	              "distribution":"secondary_push",
	              "distribution_fcm":"jpush",
	              "distribution_customize":"first_ospush",
	              "large_icon":"jgmedia-2-14b23451-0001-41ce-89d9-987b465122da",
	              "small_icon_uri":"jgmedia-3-14b23451-0001-41ce-89d9-987b465122da",
	              "style":1
	    },
	    "meizu":{
	              "distribution":"jpush",
	              "distribution_fcm":"pns"
	    },
	    "fcm":{   // 这个参数不支持 distribution_fcm 字段
	              "distribution":"jpush"
	    },
	    "oppo":{
	              "distribution":"ospush",
	              "channel_id":"*******",
	              "distribution_fcm":"secondary_fcm_push",
	              "large_icon":"jgmedia-2-14b23451-0001-41ce-89d9-987b465122da",
	              "big_pic_path":"jgmedia-1-14b23451-0001-41ce-89d9-987b465122da",
	              "style":1,
	              "skip_quota": true
	    },
	    "vivo":{
	            "distribution":"jpush",
	            "classification": 0,
	            "distribution_fcm":"secondary_pns_push",
	            "push_mode":0
	    }
	}
*/
type ThirdPartyChannel struct {
	Xiaomi *XiaomiChannel `json:"xiaomi,omitempty"`
	Huawei *HuaweiChannel `json:"huawei,omitempty"`
	Honor  *HonorChannel  `json:"honor,omitempty"`
	Meizu  *MeizuChannel  `json:"meizu,omitempty"`
	Fcm    *FcmChannel    `json:"fcm,omitempty"`
	Oppo   *OppoChannel   `json:"oppo,omitempty"`
	Vivo   *VivoChannel   `json:"vivo,omitempty"`
}

//小米
/*"xiaomi":{
"distribution":"jpush",
"channel_id":"*******",
"large_icon":"jgmedia-2-14b23451-0001-41ce-89d9-987b465122da",
"small_icon_uri":"jgmedia-3-14b23451-0001-41ce-89d9-987b465122da",
"small_icon_color":"#ABCDEF",
"big_text":"testbigtext", // 可选，最多支持 128 个字符，配合小米 style 使用
"style":1,
"distribution_fcm":"fcm",
"skip_quota": true
},*/
type XiaomiChannel struct {
	Distribution    string `json:"distribution,omitempty"`
	ChannelId       string `json:"channel_id,omitempty"`
	LargeIcon       string `json:"large_icon,omitempty"`
	SmallIconUri    string `json:"small_icon_uri,omitempty"`
	SmallIconColor  string `json:"small_icon_color,omitempty"`
	BigText         string `json:"big_text,omitempty"` // 可选，最多支持 128 个字符，配合小米 style 使用
	Style           int    `json:"style,omitempty"`
	DistributionFcm string `json:"distribution_fcm,omitempty"`
	SkipQuota       bool   `json:"skip_quota,omitempty"`
}

//华为
/*"huawei":{
"distribution":"secondary_push",
"distribution_fcm":"jpush",
"distribution_customize":"first_ospush",
"sound":"/raw/shake",
"default_sound":false,
"importance":"NORMAL",
"large_icon":"jgmedia-2-14b23451-0001-41ce-89d9-987b465122da",
"small_icon_uri":"jgmedia-3-14b23451-0001-41ce-89d9-987b465122da",
"inbox": JSONObject, //可选，配合华为 style 使用
"style":2,
"only_use_vendor_style":true
},*/
type HuaweiChannel struct {
	Category              string                 `json:"category,omitempty"`
	Distribution          string                 `json:"distribution,omitempty"`
	DistributionFcm       string                 `json:"distribution_fcm,omitempty"`
	DistributionCustomize string                 `json:"distribution_customize,omitempt"`
	Sound                 string                 `json:"sound,omitempty"`
	DefaultSound          bool                   `json:"default_sound,omitempty"`
	Importance            string                 `json:"importance,omitempty"`
	LargeIcon             string                 `json:"large_icon,omitempty"`
	SmallIconUri          string                 `json:"small_icon_uri,omitempty"`
	Inbox                 map[string]interface{} `json:"inbox,omitempty"` // 可选，配合华为 style 使用
	Style                 int                    `json:"style,omitempty"`
	OnlyUseVendorStyle    bool                   `json:"only_use_vendor_style,omitempty"`
}

//荣耀
/*"honor":{
"distribution":"secondary_push",
"distribution_fcm":"jpush",
"distribution_customize":"first_ospush",
"large_icon":"jgmedia-2-14b23451-0001-41ce-89d9-987b465122da",
"small_icon_uri":"jgmedia-3-14b23451-0001-41ce-89d9-987b465122da",
"style":1
},*/
type HonorChannel struct {
	Distribution          string `json:"distribution,omitempty"`
	DistributionFcm       string `json:"distribution_fcm,omitempty"`
	DistributionCustomize string `json:"distribution_customize,omitempt"`
	LargeIcon             string `json:"large_icon,omitempty"`
	SmallIconUri          string `json:"small_icon_uri,omitempty"`
	Style                 int    `json:"style,omitempty"`
}

//魅族
/*"meizu":{
"distribution":"jpush",
"distribution_fcm":"pns"
},*/
type MeizuChannel struct {
	Distribution    string `json:"distribution,omitempty"`
	DistributionFcm string `json:"distribution_fcm,omitempty"`
}

//fcm
/*"fcm":{
"distribution":"jpush"
},*/
type FcmChannel struct {
	Distribution string `json:"distribution,omitempty"`
}

//OPPO
/*"oppo":{
"distribution":"ospush",
"channel_id":"*******",
"distribution_fcm":"secondary_fcm_push",
"large_icon":"jgmedia-2-14b23451-0001-41ce-89d9-987b465122da",
"big_pic_path":"jgmedia-1-14b23451-0001-41ce-89d9-987b465122da",
"style":1,
"skip_quota": true
},*/
type OppoChannel struct {
	Distribution    string `json:"distribution,omitempty"`
	ChannelId       string `json:"channel_id,omitempty"`
	DistributionFcm string `json:"distribution_fcm,omitempty"`
	LargeIcon       string `json:"large_icon,omitempty"`
	BigPicPath      string `json:"big_pic_path,omitempty"`
	Style           int    `json:"style,omitempty"`
	SkipQuota       bool   `json:"skip_quota,omitempty"`
}

//VIVO
/*"vivo":{
"distribution":"jpush",
"classification": 0,
"distribution_fcm":"secondary_pns_push",
"push_mode":0
}*/
type VivoChannel struct {
	Category        string `json:"category,omitempty"`
	Distribution    string `json:"distribution,omitempty"`
	Classification  int    `json:"classification,omitempty"`
	DistributionFcm string `json:"distribution_fcm,omitempty"`
	PushMode        int    `json:"push_mode,omitempty"`
}
