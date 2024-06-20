package main

import (
	"gin-ranking/router"
)

func main() {
	r := router.Router()

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
