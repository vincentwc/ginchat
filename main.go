package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	// 初始化配置文件
	utils.InitConfig()
	// 初始化数据库
	utils.InitMySQL()

	r := router.Router()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
