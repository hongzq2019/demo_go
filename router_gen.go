package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	router "hertz_vpp/biz/router"
)

// register registers all routers.
func register(r *server.Hertz) {

	// 普通路由
	router.GeneratedRegister(r)

	// 业务路由
	customizedRegister(r)

	// 分组路由
	RegisterGroup(r)
}
