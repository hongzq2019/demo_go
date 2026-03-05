package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz_vpp/biz/orm"
)

func main() {
	// 初始化postgres数据库
	orm.InitPostgres()

	// 初始化redis数据库
	orm.InitRedis()

	h := server.New(server.WithHostPorts(":8091"))
	// 注册路由
	register(h)
	// 接受信号可退出服务
	h.Spin()
}
