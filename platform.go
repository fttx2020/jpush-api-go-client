package jpushclient

import (
	"errors"
)

const (
	IOS      = "ios"
	ANDROID  = "android"
	WINPHONE = "winphone"
)

type Platform struct {
	Os     interface{}
	osArry []string
}

func (platform *Platform) All() {
	platform.Os = "all"
}

func (platform *Platform) Add(os string) error {
	if platform.Os == nil {
		platform.osArry = make([]string, 0, 3)
	} else {
		switch platform.Os.(type) {
		case string:
			return errors.New("platform is all")
		default:
		}
	}

	//判断是否重复
	for _, value := range platform.osArry {
		if os == value {
			return nil
		}
	}

	switch os {
	case IOS:
		fallthrough
	case ANDROID:
		fallthrough
	case WINPHONE:
		platform.osArry = append(platform.osArry, os)
		platform.Os = platform.osArry
	default:
		return errors.New("unknow platform")
	}

	return nil
}

func (platform *Platform) AddIOS() {
	platform.Add(IOS)
}

func (platform *Platform) AddAndrid() {
	platform.Add(ANDROID)
}

func (platform *Platform) AddWinphone() {
	platform.Add(WINPHONE)
}
