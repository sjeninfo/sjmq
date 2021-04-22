package events

import (
	"time"
)

type OrderUpdated struct {
	ID           uint          `json:"id"`            // 流水號
	CreatedAt    time.Time     `json:"created_at"`    // 建檔日期
	UpdatedAt    time.Time     `json:"updated_at"`    // 修改日期
	No           string        `json:"no"`            // 採購單號
	Date         string        `json:"date"`          // 採購日期
	ProviderID   uint          `json:"provider_id"`   // 廠商關聯ID
	AssigneeID   uint          `json:"assignee_id"`   // 承辦人關聯ID
	ValidityFrom string        `json:"validity_from"` // 效期起始日
	ValidityEnd  string        `json:"validity_end"`  // 效期終止日
	SystemHint   string        `json:"system_hint"`   // 系統提示
	Remark       string        `json:"remark"`        // 備註
	LockedAt     *time.Time    `json:"locked_at"`     // 結單時間
	ClosedAt     *time.Time    `json:"closed_at"`     // 結案時間
	Details      []OrderDetail `json:"details"`       // 訂單明細
	UpdatedByID  *uint         `json:"updated_by_id"` // 修改人關聯ID
	CreatedByID  *uint         `json:"created_by_id"` // 建檔人關聯ID
}

type OrderDetail struct {
	ID            uint    `json:"id"`             // 流水號
	ItemID        uint    `json:"item_id"`        // 商品關聯ID
	Quantity      float64 `json:"quantity"`       // 數量
	Giveaway      float64 `json:"giveaway"`       // 贈量
	Price         float64 `json:"price"`          // 單價
	Discount      float64 `json:"discount"`       // 折讓
	Sum           float64 `json:"sum"`            // 稅前合計(Sum + Tax - Discount)
	Total         float64 `json:"total"`          // 稅後總計
	TaxType       uint    `json:"tax_type"`       // 稅別
	Tax           float64 `json:"tax"`            // 稅額
	Remark        string  `json:"remark"`         // 備註
	OrderPriority uint    `json:"order_priority"` // 排序編號
	BasicItemID   uint    `json:"basic_item_id"`  // 過賬商品
	PostingQty    float64 `json:"posting_qty"`    // 過賬數量
}

type PurchaseUpdated struct {
	ID                          uint             `json:"id"`                             // 流水號
	CreatedAt                   time.Time        `json:"created_at"`                     // 建檔日期
	UpdatedAt                   time.Time        `json:"updated_at"`                     // 修改日期
	No                          string           `json:"no"`                             // 進貨單編號
	Date                        string           `json:"date"`                           // 進貨單日期
	WarehouseID                 uint             `json:"warehouse_id"`                   // 倉庫關聯ID
	ProviderID                  uint             `json:"provider_id"`                    // 廠商關聯ID
	AssigneeID                  uint             `json:"assignee_id"`                    // 承辦人關聯ID
	TranNo                      string           `json:"tran_no"`                        // 廠商貨單號
	Remark                      string           `json:"remark"`                         // 備註
	SystemHint                  string           `json:"system_hint"`                    // 系統資訊
	InvoiceNo                   string           `json:"invoice_no"`                     // 發票號碼
	InvoiceDate                 string           `json:"invoice_date"`                   // 發票日期
	OrderID                     *uint            `json:"order_id"`                       // 採購單號關聯ID
	ControlledDrugsTransferFlag bool             `json:"controlled_drugs_transfer_flag"` // 管藥受讓旗標
	LockedAt                    *time.Time       `json:"locked_at"`                      // 結單時間
	ClosedAt                    *time.Time       `json:"closed_at"`                      // 結案時間
	PayableClosedAt             *time.Time       `json:"payable_closed_at"`              // 應付結案時間
	Details                     []PurchaseDetail `json:"details"`                        // 訂單明細
	UpdatedByID                 *uint            `json:"updated_by_id"`                  // 修改人關聯ID
	CreatedByID                 *uint            `json:"created_by_id"`                  // 建檔人關聯ID
}

type PurchaseDetail struct {
	ID            uint    `json:"id"`             // 流水號
	ItemID        uint    `json:"item_id"`        // 商品關聯ID
	Quantity      float64 `json:"quantity"`       // 數量
	Giveaway      float64 `json:"giveaway"`       // 贈量
	Price         float64 `json:"price"`          // 單價
	Discount      float64 `json:"discount"`       // 折讓
	Sum           float64 `json:"sum"`            // 稅前合計(Sum + Tax - Discount)
	Total         float64 `json:"total"`          // 稅後總計
	TaxType       uint    `json:"tax_type"`       // 稅別
	Tax           float64 `json:"tax"`            // 稅額
	Remark        string  `json:"remark"`         // 備註
	OrderPriority uint    `json:"order_priority"` // 排序編號
	LotNo         string  `json:"lot_no"`         // 批號
	ExpiryDate    string  `json:"expiry_date"`    // 效期
	BasicItemID   uint    `json:"basic_item_id"`  // 過賬商品
	PostingQty    float64 `json:"posting_qty"`    // 過賬數量
	//RequisitionID	*uint	requisition_id	請購編號
}

type PurchaseReturnUpdated struct {
	ID              uint                   `json:"id"`                  // 流水號
	CreatedAt       time.Time              `json:"created_at"`          // 建檔日期
	UpdatedAt       time.Time              `json:"updated_at"`          // 修改日期
	No              string                 `json:"no"`                  // 進退單編號
	Date            string                 `json:"date"`                // 進退單日期
	PurchaseID      uint                   `json:"purchase_id"`         // 進貨單關聯ID
	AssigneeID      uint                   `json:"assignee_id"`         // 承辦人關聯ID
	Remark          string                 `json:"remark"`              // 備註
	SystemHint      string                 `json:"system_hint"`         // 系統資訊
	ReturnedAt      *time.Time             `json:"returned_at"`         // 退貨處理時間
	ReturnInvoiceNo string                 `json:"returned_invoice_no"` // 進貨退出單編號
	LockedAt        *time.Time             `json:"locked_at"`           // 結單時間
	ClosedAt        *time.Time             `json:"closed_at"`           // 結案時間
	PayableClosedAt *time.Time             `json:"payable_closed_at"`   // 應付結案時間
	Details         []PurchaseReturnDetail `json:"details"`             // 訂單明細
	UpdatedByID     *uint                  `json:"updated_by_id"`       // 修改人關聯ID
	CreatedByID     *uint                  `json:"created_by_id"`       // 建檔人關聯ID
}

type PurchaseReturnDetail struct {
	ID            uint    `json:"id"`             // 流水號
	ItemID        uint    `json:"item_id"`        // 商品關聯ID
	Quantity      float64 `json:"quantity"`       // 數量
	Giveaway      float64 `json:"giveaway"`       // 贈量
	Price         float64 `json:"price"`          // 單價
	Discount      float64 `json:"discount"`       // 折讓
	Sum           float64 `json:"sum"`            // 稅前合計(Sum + Tax - Discount)
	Total         float64 `json:"total"`          // 稅後總計
	TaxType       uint    `json:"tax_type"`       // 稅別
	Tax           float64 `json:"tax"`            // 稅額
	Remark        string  `json:"remark"`         // 備註
	OrderPriority uint    `json:"order_priority"` // 排序編號
	LotNo         string  `json:"lot_no"`         // 批號
	BasicItemID   uint    `json:"basic_item_id"`  // 過賬商品
	PostingQty    float64 `json:"posting_qty"`    // 過賬數量
}

type ShippingUpdated struct {
	ID                          uint             `json:"id"`                             // 流水號
	CreatedAt                   time.Time        `json:"created_at"`                     // 建檔日期
	UpdatedAt                   time.Time        `json:"updated_at"`                     // 修改日期
	No                          string           `json:"no"`                             // 出貨單編號
	Date                        string           `json:"date"`                           // 出貨單日期
	WarehouseID                 uint             `json:"warehouse_id"`                   // 倉庫關聯ID
	CustomerID                  uint             `json:"customer_id"`                    // 客戶關聯ID
	AssigneeID                  uint             `json:"assignee_id"`                    // 承辦人關聯ID
	Remark                      string           `json:"remark"`                         // 備註
	SystemHint                  string           `json:"system_hint"`                    // 系統資訊
	InvoiceNo                   string           `json:"invoice_no"`                     // 發票號碼
	ControlledDrugsTransferFlag bool             `json:"controlled_drugs_transfer_flag"` // 管藥受讓旗標
	LockedAt                    *time.Time       `json:"locked_at"`                      // 結單時間
	ClosedAt                    *time.Time       `json:"closed_at"`                      // 結案時間
	ReceivableClosedAt          *time.Time       `json:"receivable_closed_at"`           // 應收結案時間
	Details                     []ShippingDetail `json:"details"`                        // 訂單明細
	UpdatedByID                 *uint            `json:"updated_by_id"`                  // 修改人關聯ID
	CreatedByID                 *uint            `json:"created_by_id"`                  // 建檔人關聯ID
}

type ShippingDetail struct {
	ID            uint    `json:"id"`             // 流水號
	ItemID        uint    `json:"item_id"`        // 商品關聯ID
	Quantity      float64 `json:"quantity"`       // 數量
	Giveaway      float64 `json:"giveaway"`       // 贈量
	Price         float64 `json:"price"`          // 單價
	Discount      float64 `json:"discount"`       // 折讓
	Sum           float64 `json:"sum"`            // 稅前合計(Sum + Tax - Discount)
	Total         float64 `json:"total"`          // 稅後總計
	TaxType       uint    `json:"tax_type"`       // 稅別
	Tax           float64 `json:"tax"`            // 稅額
	Remark        string  `json:"remark"`         // 備註
	OrderPriority uint    `json:"order_priority"` // 排序編號
	LotNo         string  `json:"lot_no"`         // 批號
	ExpiryDate    string  `json:"expiry_date"`    // 效期
	BasicItemID   uint    `json:"basic_item_id"`  // 過賬商品
	PostingQty    float64 `json:"posting_qty"`    // 過賬數量
}

type ShippingReturnUpdated struct {
	ID                 uint                   `json:"id"`                   // 流水號
	CreatedAt          time.Time              `json:"created_at"`           // 建檔日期
	UpdatedAt          time.Time              `json:"updated_at"`           // 修改日期
	No                 string                 `json:"no"`                   // 出退單編號
	Date               string                 `json:"date"`                 // 出退單日期
	ShippingID         uint                   `json:"shipping_id"`          // 出貨單關聯ID
	AssigneeID         uint                   `json:"assignee_id"`          // 承辦人關聯ID
	Remark             string                 `json:"remark"`               // 備註
	SystemHint         string                 `json:"system_hint"`          // 系統資訊
	LockedAt           *time.Time             `json:"locked_at"`            // 結單時間
	ClosedAt           *time.Time             `json:"closed_at"`            // 結案時間
	ReceivableClosedAt *time.Time             `json:"receivable_closed_at"` // 應收結案時間
	Details            []ShippingReturnDetail `json:"details"`              // 訂單明細
	UpdatedByID        *uint                  `json:"updated_by_id"`        // 修改人關聯ID
	CreatedByID        *uint                  `json:"created_by_id"`        // 建檔人關聯ID
}

type ShippingReturnDetail struct {
	ID            uint    `json:"id"`             // 流水號
	ItemID        uint    `json:"item_id"`        // 商品關聯ID
	Quantity      float64 `json:"quantity"`       // 數量
	Giveaway      float64 `json:"giveaway"`       // 贈量
	Price         float64 `json:"price"`          // 單價
	Discount      float64 `json:"discount"`       // 折讓
	Sum           float64 `json:"sum"`            // 稅前合計(Sum + Tax - Discount)
	Total         float64 `json:"total"`          // 稅後總計
	TaxType       uint    `json:"tax_type"`       // 稅別
	Tax           float64 `json:"tax"`            // 稅額
	Remark        string  `json:"remark"`         // 備註
	OrderPriority uint    `json:"order_priority"` // 排序編號
	LotNo         string  `json:"lot_no"`         // 批號
	BasicItemID   uint    `json:"basic_item_id"`  // 過賬商品
	PostingQty    float64 `json:"posting_qty"`    // 過賬數量
}

type StocktakeUpdated struct {
	ID          uint              `json:"id"`            // 流水號
	CreatedAt   time.Time         `json:"created_at"`    // 建檔日期
	UpdatedAt   time.Time         `json:"updated_at"`    // 修改日期
	No          string            `json:"no"`            // 盤點單編號
	Date        string            `json:"date"`          // 盤點單日期
	FormType    uint              `json:"formType"`      // 單據類型(1.一般盤點單 2.商品調整單)
	WarehouseID uint              `json:"warehouse_id"`  // 倉庫關聯ID
	AssigneeID  uint              `json:"assignee_id"`   // 承辦人關聯ID
	Remark      string            `json:"remark"`        // 備註
	SystemHint  string            `json:"system_hint"`   // 系統資訊
	LockedAt    *time.Time        `json:"locked_at"`     // 結單時間
	ClosedAt    *time.Time        `json:"closed_at"`     // 結案時間
	Details     []StocktakeDetail `json:"details"`       // 訂單明細
	UpdatedByID *uint             `json:"updated_by_id"` // 修改人關聯ID
	CreatedByID *uint             `json:"created_by_id"` // 建檔人關聯ID
}

type StocktakeDetail struct {
	ID                    uint    `json:"id"`                      // 流水號
	ItemID                uint    `json:"item_id"`                 // 商品關聯ID
	OnHandQty             float64 `json:"on_hand_qty"`             // 存量
	VerificationQty       float64 `json:"verification_qty"`        // 盤量
	PostingQty            float64 `json:"posting_qty"`             // 盈虧量
	Remark                string  `json:"remark"`                  // 備註
	LotNo                 string  `json:"lot_no"`                  // 批號
	PreviousStocktakeDate string  `json:"previous_stocktake_date"` // 上次盤點日
	OrderPriority         uint    `json:"order_priority"`          // 排序編號
}

type TransferUpdated struct {
	ID                     uint             `json:"id"`                       // 流水號
	CreatedAt              time.Time        `json:"created_at"`               // 建檔日期
	UpdatedAt              time.Time        `json:"updated_at"`               // 修改日期
	No                     string           `json:"no"`                       // 移倉單編號
	Date                   string           `json:"date"`                     // 移倉單日期
	StartWarehouseID       uint             `json:"start_warehouse_id"`       // 移出倉關聯ID
	DestinationWarehouseID uint             `json:"destination_warehouse_id"` // 目的倉關聯ID
	AssigneeID             uint             `json:"assignee_id"`              // 承辦人關聯ID
	DeliveryType           uint             `json:"delivery_type"`            // 送貨方式
	Remark                 string           `json:"remark"`                   // 備註
	SystemHint             string           `json:"system_hint"`              // 系統資訊
	LockedAt               *time.Time       `json:"locked_at"`                // 結單時間
	ClosedAt               *time.Time       `json:"closed_at"`                // 結案時間
	Details                []TransferDetail `json:"details"`                  // 移倉明細
	UpdatedByID            *uint            `json:"updated_by_id"`            // 修改人關聯ID
	CreatedByID            *uint            `json:"created_by_id"`            // 建檔人關聯ID
}

type TransferDetail struct {
	ID            uint    `json:"id"`             // 流水號
	ItemID        uint    `json:"item_id"`        // 商品關聯ID
	Quantity      float64 `json:"quantity"`       // 數量
	Remark        string  `json:"remark"`         // 備註
	OrderPriority uint    `json:"order_priority"` // 排序編號
	LotNo         string  `json:"lot_no"`         // 批號
	ExpiryDate    string  `json:"expiry_date"`    // 效期
	BasicItemID   uint    `json:"basic_item_id"`  // 過賬商品
	PostingQty    float64 `json:"posting_qty"`    // 過賬數量
	//RequisitionID	*uint	requisition_id	請購編號
}
