package modules

import "time"

// CmCustomer 实体
type CmCustomer struct {
	// 主键,自增
	CustomerId   int       `column:"customer_id" json:"customerId" gorm:"primaryKey;autoIncrement"`
	CustomerName string    `column:"customer_name" json:"customerName"`
	ShortName    string    `column:"short_name" json:"shortName"`
	TaxNo        string    `column:"tax_no" json:"taxNo"`
	CreateDate   time.Time `column:"create_date" json:"createTime"`
	Status       int       `column:"status" json:"status"`
}

func (cm *CmCustomer) TableName() string {
	return "cm_customer"
}
