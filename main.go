package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/router"
	"log"
)

func main() {

	// 初始化依赖模块
	fmt.Printf("初始化>>>>>>>>>>")
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	router.InitRouter()

}
