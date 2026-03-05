package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz_vpp/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {

	// 获取所有企业客户数据信息
	r.Handle("GET", "/ListCustomer", handler.ListCustomer)

}

// RegisterGroup . 分组接口
func RegisterGroup(r *server.Hertz) {

	// v1版本
	v1 := r.Group("/v1")
	{
		// 根据Id获取企业客户信息
		v1.GET("getCustomerById", handler.GetCustomerById)

		// 创建企业客户信息
		v1.POST("createCustomer", handler.CreateCustomer)
	}

}
