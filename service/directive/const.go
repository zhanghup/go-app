package directive

type TokenType string

const (
	TokenPc  TokenType = "pc"     // 网页登录
	TokenMp  TokenType = "wx-mp"  // 微信公众号
	TokenMi TokenType = "wx-mi" // 微信小程序
)
