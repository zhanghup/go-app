package cfg

import "fmt"

type configWxQy struct {
	Enable bool   `ini:"enable"`
	Corpid string `ini:"corpid"`
}

type configWxQyApp struct {
	Enable bool   `ini:"enable"`
	Secret string `ini:"secret"`
}

// 微信企业号配置
func wxqy(flag ...bool) *configWxQy {
	if my.Wxqy == nil {
		panic("config.ini - [wxqy] - 配置文件微信企业号信息尚未初始化完成")
	}

	if (len(flag) == 0 || flag[0]) && !my.Wxqy.Enable {
		panic("config.ini - [wxqy].enable 未启用")
	}
	return my.Wxqy
}
func Wxqy() *configWxQy {
	return wxqy()
}
func WxqyEnable() bool {
	return wxqy(false).Enable
}

// 微信企业应用配置
func wxqyApp(agentid string, flag ...bool) *configWxQyApp {
	if my.Wxqy == nil {
		panic("config.ini - [wxqy-app] - 配置文件微信企业号应用信息尚未初始化完成")
	}
	obj, ok := my.WxqyApp[agentid]
	if !ok {
		panic(fmt.Sprintf(`config.ini - [wxqy-app "%s"] - 没有找到该应用的agentid`, agentid))
	}

	if (len(flag) == 0 || flag[0]) && !obj.Enable {
		panic(fmt.Sprintf(`config.ini - [wxqy-app "%s"].enable 未启用`, agentid))
	}
	return &obj
}
func WxqyApp(agentid string) *configWxQyApp {
	return wxqyApp(agentid)
}

func WxqyAppEnable(agentid string) bool {
	return wxqyApp(agentid, false).Enable
}
