package main

import (
	"gin_oj.cn/config"
	"gin_oj.cn/router"
)

func main() {
	config.InitConfig()
	r := router.Router()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
