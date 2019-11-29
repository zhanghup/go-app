package cfg

import (
	"bytes"
	"io"
	"os"
)

var FileConfigString = `
[web]
enable = false
port = 8899

[database]
enable = false
type = mysql
uri = root:123@/test?charset=utf8

[wxqy]
enable = false
corpid = ww87b71522a68e1b3f

[wxqy-app "1000002"]
enable = false
secret = FJiFUfgSI66DTYeZtLhxUS0Wb8UWF3mS_29HY1elzBM
`

func InitConfFile() {
	_, err := os.Open("conf")
	if err == nil || !os.IsNotExist(err) {
		return
	}
	err = os.Mkdir("conf", os.ModePerm)
	if err != nil {
		panic("创建conf文件夹异常")
	}
	f, err := os.Create("conf/config.ini")
	if err != nil {
		panic("创建config.ini文件异常")
	}
	_, err = io.Copy(f, bytes.NewBuffer([]byte(FileConfigString)))
	if err != nil {
		panic("写入config.ini文件异常")
	}
}
