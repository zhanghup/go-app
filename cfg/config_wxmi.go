package cfg

type configWxmi struct {
	Enable bool `ini:"enable"`

	Appid string `ini:"appid"`
	AppSecret string `ini:"appsecret"`
}

// api服务配置
func wxmi(flag ...bool) *configWxmi {
	if my.Wxmi == nil {
		panic("config.ini - [wxmi] - 配置文件wxmp信息尚未初始化完成")
	}
	if (len(flag) == 0 || flag[0]) && !my.Wxmp.Enable {
		panic("config.ini - [wxmi].enable 未启用")
	}
	return my.Wxmi
}
func Wxmi() *configWxmi {
	return wxmi()
}

func WxmiEnable() bool {
	return wxmi(false).Enable
}
