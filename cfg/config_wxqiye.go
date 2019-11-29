package cfg

type configWxQy struct {
	Enable bool   `ini:"enable"`
	Corpid string `ini:"corpid"`
}

type configWxQyApp struct {
	Enable bool   `ini:"enable"`
	Secret string `ini:"secret"`
}
