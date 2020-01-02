package ctx

type configWxmp struct {
	Enable bool `ini:"enable"`

	Appid string `ini:"appid"`
	AppSecret string `ini:"appsecret"`
}

// api服务配置
func wxmp(flag ...bool) *configWxmp {
	if my.Wxmp == nil {
		panic("config.ini - [wxmp] - 配置文件wxmp信息尚未初始化完成")
	}
	if (len(flag) == 0 || flag[0]) && !my.Wxmp.Enable {
		panic("config.ini - [wxmp].enable 未启用")
	}
	return my.Wxmp
}
func Wxmp() *configWxmp {
	return wxmp()
}

func WxmpEnable() bool {
	return wxmp(false).Enable
}
