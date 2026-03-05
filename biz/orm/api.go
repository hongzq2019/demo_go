package orm

import "hertz_vpp/biz/modules"

// CmCustomerRequest . 入参请求
type CmCustomerRequest struct {
	CustomerName string `json:"customerName,required" query:"customerName"`
	ShortName    string `json:"shortName" query:"shortName"`
	TaxNo        string `json:"taxNo" query:"taxNo"`
}

// CmCustomerResponse 企业客户响应参数
type CmCustomerResponse struct {
	CustomerId   int    `json:"customerId"`
	CustomerName string `json:"customerName"`
}

// CreateCustomer 创建企业客户
func CreateCustomer(cm *modules.CmCustomer) error {
	return DB.Create(cm).Error
}

func ListCustomers() (customers []*modules.CmCustomer, err error) {
	db := DB.Model(modules.CmCustomer{})
	return customers, db.Find(&customers).Error
}

// QueryCustomerById 根据customerId查询企业客户
func QueryCustomerById(customerId int) (*modules.CmCustomer, error) {
	db := DB.Model(modules.CmCustomer{})
	if customerId != 0 {
		db = db.Where("customer_id =?", customerId)
	}
	var res *modules.CmCustomer
	if err := db.First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetDeviceInfosByCustomerId 根据customerId查询所有设备信息
func GetDeviceInfosByCustomerId(customerId int) (devices []*modules.DeviceInfo, err error) {
	db := DB.Model(modules.DeviceInfo{})
	if customerId != 0 {
		db = db.Where("customer_id =?", customerId)
	}
	return devices, db.Find(&devices).Error
}
