package main

import (

	"github.com/zhanghup/go-app/api/server/engine"
)

func main() {
	engine.Router().Run(":8899")
}
