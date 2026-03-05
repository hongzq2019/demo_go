package modules

// DeviceInfo 设备信息
type DeviceInfo struct {
	Id           string `column:"id"`
	DeviceId     string `column:"device_id"`
	CustomerId   string `column:"customer_id"`
	CustomerName string `column:"customer_name"`
	DeviceName   string `column:"device_name"`
	DeviceTypeId string `column:"device_type_id"`
	CreateDate   string `column:"create_date"`
}

// TableName 表名
func (device *DeviceInfo) TableName() string {
	return "cm_customer_device"
}
