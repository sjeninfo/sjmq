package events

import (
	"gorm.io/gorm"
	"time"
)

type TransferUpdated struct {
	ID                     uint              `json:"id" mapstructure:"id" gorm:"primaryKey"`                                                                  // 流水號
	CreatedAt              time.Time         `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                                           // 建檔日期
	UpdatedAt              time.Time         `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                                           // 修改日期
	DeletedAt              gorm.DeletedAt    `json:"-" mapstructure:"deleted_at" gorm:"Column:deleted_at;index"`                                              // 使用旗標
	No                     string            `json:"no" mapstructure:"no" gorm:"Column:no"`                                                                   // 移倉單編號
	Date                   string            `json:"date" mapstructure:"date" gorm:"Column:date"`                                                             // 移倉單日期
	StartWarehouseID       uint              `json:"start_warehouse_id" mapstructure:"start_warehouse_id" gorm:"Column:start_warehouse_id"`                   // 移出倉關聯ID
	DestinationWarehouseID uint              `json:"destination_warehouse_id" mapstructure:"destination_warehouse_id" gorm:"Column:destination_warehouse_id"` // 目的倉關聯ID
	AssigneeID             uint              `json:"assignee_id" mapstructure:"assignee_id" gorm:"Column:assignee_id"`                                        // 承辦人關聯ID
	DeliveryType           uint              `json:"delivery_type" mapstructure:"delivery_type" gorm:"Column:delivery_type"`                                  // 送貨方式
	Remark                 string            `json:"remark" mapstructure:"remark" gorm:"Column:remark"`                                                       // 備註
	SystemHint             string            `json:"system_hint" mapstructure:"system_hint" gorm:"Column:system_hint"`                                        // 系統資訊
	LockedAt               *time.Time        `json:"locked_at" mapstructure:"locked_at" gorm:"Column:locked_at"`                                              // 結單時間
	ClosedAt               *time.Time        `json:"closed_at" mapstructure:"closed_at" gorm:"Column:closed_at"`                                              // 結案時間
	Details                []*TransferDetail `json:"details" mapstructure:"details" gorm:"foreignKey:TransferID"`                                             // 移倉明細
	UpdatedByID            *uint             `json:"updated_by_id" mapstructure:"updated_by_id" gorm:"Column:updated_by_id"`                                  // 修改人關聯ID
	CreatedByID            *uint             `json:"created_by_id" mapstructure:"created_by_id" gorm:"Column:created_by_id"`                                  // 建檔人關聯ID
}

type TransferDetail struct {
	ID         uint           `json:"id" mapstructure:"id" gorm:"primaryKey"`                           // 流水號
	CreatedAt  time.Time      `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`    // 建檔日期
	UpdatedAt  time.Time      `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`    // 修改日期
	DeletedAt  gorm.DeletedAt `json:"-" mapstructure:"deleted_at" gorm:"Column:deleted_at;index"`       // 使用旗標
	TransferID uint           `json:"transfer_id" mapstructure:"transfer_id" gorm:"Column:transfer_id"` // 移倉關聯ID
	ItemID     uint           `json:"item_id" mapstructure:"item_id" gorm:"Column:item_id"`             // 商品關聯ID
	Quantity   float64        `json:"quantity" mapstructure:"quantity" gorm:"Columns:quantity"`         // 數量
	//過賬數量
	Remark        string `json:"remark" mapstructure:"remark" gorm:"Columns:remark"`                         //	備註
	OrderPriority uint   `json:"order_priority" mapstructure:"order_priority" gorm:"Columns:order_priority"` //	排序編號
	LotNo         string `json:"lot_no" mapstructure:"lot_no" gorm:"Column:lot_no"`                          // 批號
	ExpiryDate    string `json:"expiry_date" mapstructure:"expiry_date" gorm:"Column:expiry_date"`           // 效期
	//RequisitionID	*uint	requisition_id	請購編號
}
