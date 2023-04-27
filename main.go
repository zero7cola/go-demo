package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go1/bootstrap"
	baseconfig "go1/config"
	"go1/pkg/config"
)

func init() {
	baseconfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	router := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()

	bootstrap.SetupRoute(router)

	err := router.Run(":3000")

	if err != nil {
		fmt.Println(err.Error())
	}
}
