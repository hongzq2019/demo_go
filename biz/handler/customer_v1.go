package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	result "hertz_vpp/biz/config"
	"hertz_vpp/biz/modules"
	"hertz_vpp/biz/orm"
	"log"
	"strconv"
	"time"
)

func ListCustomer(ctx context.Context, c *app.RequestContext) {
	// 获取所有客户信息
	customers, err := orm.ListCustomers()
	if err != nil {
		return
	}
	//cmCustomerResp := make([]orm.CmCustomerResponse, 0, len(customers))
	//for i := range customers {
	//	cmCustomerResponse := orm.CmCustomerResponse{
	//		CustomerId:   customers[i].CustomerId,
	//		CustomerName: customers[i].CustomerName,
	//	}
	//	cmCustomerResp = append(cmCustomerResp, cmCustomerResponse)
	//}

	cmList := make([]map[string]interface{}, 0, len(customers))
	for _cm := range customers {
		cmList = append(cmList, map[string]interface{}{
			"customerId":   customers[_cm].CustomerId,
			"customerName": customers[_cm].CustomerName,
		})
	}

	typeList := make(map[int]string)
	typeList[2] = "删除"
	for i := 0; i <= 1; i++ {
		var value string
		if i == 1 {
			value = "新增"
		} else {
			value = "修改"
		}
		typeList[i] = value
	}

	results := map[string]interface{}{
		"customers": cmList,
		"typeList":  typeList,
	}
	log.Println(fmt.Sprintf("输出列表长度为: %d", len(cmList)))
	// 响应
	c.JSON(consts.StatusOK, result.SuccessData("操作成功", results))
}

// GetCustomerById 根据customerId查询客户信息
func GetCustomerById(ctx context.Context, c *app.RequestContext) {

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 企业客户Id
	customerIdStr := c.Query("customerId")
	if customerIdStr == "" {
		c.JSON(consts.StatusOK, result.Fail("参数请求错误", consts.StatusBadRequest))
		return
	}

	customerId, _ := strconv.Atoi(customerIdStr)
	customer, err := orm.QueryCustomerById(customerId)
	if err != nil {
		return
	}
	// 获取该企业客户下所有设备信息
	deviceInfos, _ := orm.GetDeviceInfosByCustomerId(customerId)
	resultMaps := map[string]interface{}{
		"customerInfo": customer,
		"deviceInfos":  deviceInfos,
	}

	// 存入Redis
	var rdb = orm.InitRedis()
	resultJson, err := json.Marshal(resultMaps)
	rdb.Set(ctx, "customerId", resultJson, time.Hour)

	redisResult, err := rdb.Get(ctx, "customerId").Result()
	if err == nil {
		log.Println(fmt.Sprintf("println Redis Json %s", redisResult))
	}

	// 返回
	c.JSON(consts.StatusOK, resultMaps)
}

// CreateCustomer 创建企业客户信息
func CreateCustomer(ctx context.Context, c *app.RequestContext) {

	// 入参
	customerRequest := &orm.CmCustomerRequest{}
	// 请求参数绑定
	if err := c.BindAndValidate(customerRequest); err != nil {
		// 参数错误
		c.JSON(consts.StatusOK, result.Fail("参数请求错误", consts.StatusBadRequest))
		return
	}
	// 赋值企业客户信息
	customer := &modules.CmCustomer{
		CustomerName: customerRequest.CustomerName,
		ShortName:    customerRequest.ShortName,
		TaxNo:        customerRequest.TaxNo,
		Status:       0,
		CreateDate:   time.Now(),
	}
	// 创建企业客户
	err := orm.CreateCustomer(customer)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, result.Fail("创建企业客户失败", consts.StatusInternalServerError))
		return
	}

	c.JSON(consts.StatusOK, result.SimpleSuccess())
}
