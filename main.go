package main

import (
	"fmt"
	"github.com/kobshobe/minio_test/router"
)

func main() {
	fmt.Println("ok")
	engine := router.Init()
	err := engine.Run(":8089")
	if err != nil {
		panic(err)
	}
}
