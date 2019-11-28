package main

import "github.com/zhanghup/go-app/boot"

func main() {
	boot.Boot(boot.BootOption{
		DB: boot.BODB{
			Type:    "mysql",
			Uri:     "root:123@/test?charset=utf8",
			ShowSql: true,
		},
		Web: boot.BOWeb{
			Port: "8899",
		},
	})
}
