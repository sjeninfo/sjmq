package events

import (
	"gorm.io/gorm"
	"time"
)

type UpdateArea struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 區域代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 區域名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdateBankAccount struct {
	ID             uint      `json:"id" mapstructure:"id"`                             // ID
	CreatedAt      time.Time `json:"created_at" mapstructure:"created_at"`             // 建檔日期
	UpdatedAt      time.Time `json:"updated_at" mapstructure:"updated_at"`             // 修改日期
	No             string    `json:"no" mapstructure:"no"`                             // 銀行帳號
	Name           string    `json:"name" mapstructure:"name"`                         // 銀行名稱
	QuickCode      string    `json:"quick_code" mapstructure:"quick_code"`             // 簡碼
	AccountType    uint      `json:"account_type" mapstructure:"account_type"`         // 帳戶別(1.甲存 2.乙存 3.刷卡銀行)
	CheckType      uint      `json:"check_type" mapstructure:"check_type"`             // 支票別(1.本票 2.支票)
	SubjectID      uint      `json:"subject_id" mapstructure:"subject_id"`             // 會計科目
	PhoneNo        string    `json:"phone_no" mapstructure:"phone_no"`                 // 電話
	DepositAmount  float64   `json:"deposit_amount" mapstructure:"deposit_amount"`     // 存款金額
	MarginOfSafety float64   `json:"margin_of_safety" mapstructure:"margin_of_safety"` // 安全餘額
	Blocked        bool      `json:"blocked" mapstructure:"blocked"`                   // 停用
	Used           bool      `json:"used" mapstructure:"used"`                         // 使用過期標
	UpdatedByID    *uint     `json:"updated_by_id" mapstructure:"updated_by_id"`       // 修改人關聯ID
	CreatedByID    *uint     `json:"created_by_id" mapstructure:"created_by_id"`       // 建檔人關聯ID
}

type UpdateBasicItem struct {
	UpdateItem     `mapstructure:",squash"`
	ProviderID     uint    `json:"provider_id" mapstructure:"provider_id"`         // 主供應商關聯ID
	OnHandQty      float64 `json:"on_hand_qty" mapstructure:"on_hand_qty"`         // 現有存量
	InitQty        float64 `json:"init_qty" mapstructure:"init_qty"`               // 期初數量
	InitCost       float64 `json:"init_cost" mapstructure:"init_cost"`             // 期初成本
	BasicQty       float64 `json:"basic_qty" mapstructure:"basic_qty"`             // 基本採購量
	ShipDate       string  `json:"ship_date" mapstructure:"ship_date"`             // 出貨日期
	PurchaseDate   string  `json:"purchase_date" mapstructure:"purchase_date"`     // 進貨日期
	LicenseType    uint    `json:"license_type" mapstructure:"license_type"`       // 許可證類別(0.未填入 1.衛署藥製 2.衛署藥輸 3.衛署 4.內衛藥製 5.內衛藥輸 6.內衛)
	LicenseNo      string  `json:"license_no" mapstructure:"license_no"`           // 許可證字號
	AllowPurchase  bool    `json:"allow_purchase" mapstructure:"allow_purchase"`   // 允許進貨
	AllowInventory bool    `json:"allow_inventory" mapstructure:"allow_inventory"` // 允許盤點
}

type ContactDeleted struct {
	ID uint `json:"id" mapstructure:"id"`
}

type ContactUpdated struct {
	ID              uint      `json:"id" mapstructure:"id"`                             // 流水號
	CreatedAt       time.Time `json:"created_at" mapstructure:"created_at"`             // 建檔日期
	UpdatedAt       time.Time `json:"updated_at" mapstructure:"updated_at"`             // 修改日期
	No              string    `json:"no" mapstructure:"no"`                             // 通訊錄編號
	Name            string    `json:"name" mapstructure:"name"`                         // 通訊錄名稱
	CategoryID      uint      `json:"category_id" mapstructure:"category_id"`           // 通訊錄類別關聯ID
	QuickCode       string    `json:"quick_code" mapstructure:"quick_code"`             // 簡碼
	PhoneNo1        string    `json:"phone_no1" mapstructure:"phone_no1"`               // 電話1
	PhoneNo2        string    `json:"phone_no2" mapstructure:"phone_no2"`               // 電話2
	MobilePhone     string    `json:"mobile_phone" mapstructure:"mobile_phone"`         // 行動電話
	Fax             string    `json:"fax" mapstructure:"fax"`                           // 傳真
	Contact         string    `json:"contact" mapstructure:"contact"`                   // 連絡人
	Uniform         string    `json:"uniform" mapstructure:"uniform"`                   // 統一編號
	BillingAddress  string    `json:"billing_address" mapstructure:"billing_address"`   // 帳單地址
	Email           string    `json:"email" mapstructure:"email"`                       // 電子郵件
	ZipCode         string    `json:"zip_code" mapstructure:"zip_code"`                 // 郵遞區號(聯絡地址)
	ZipCode2        string    `json:"zip_code2" mapstructure:"zip_code2"`               // 郵遞區號(發票地址)
	ZipCode3        string    `json:"zip_code3" mapstructure:"zip_code3"`               // 郵遞區號(送貨地址)
	Hint            string    `json:"hint" mapstructure:"hint"`                         // 重要資訊
	Remark          string    `json:"remark" mapstructure:"remark"`                     // 備註
	InvoiceTitle    string    `json:"invoice_title" mapstructure:"invoice_title"`       // 發票抬頭
	InvoiceAddress  string    `json:"invoice_address" mapstructure:"invoice_address"`   // 發票地址
	DeliveryAddress string    `json:"delivery_address" mapstructure:"delivery_address"` // 送貨地址
	Website         string    `json:"website" mapstructure:"website"`                   // 網址
	UpdatedByID     *uint     `json:"updated_by_id" mapstructure:"updated_by_id"`       // 修改人關聯ID
	CreatedByID     *uint     `json:"created_by_id" mapstructure:"created_by_id"`       // 建檔人關聯ID
}

type ContactCategoryDeleted struct {
	ID uint `json:"id" mapstructure:"id"`
}

type ContactCategoryUpdated struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 通訊錄類別代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 通訊錄類別名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdateCustomerCategory struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 客戶類別代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 客戶類別名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdateCustomer struct {
	ID                  uint           `json:"id" mapstructure:"id"`                                       // 流水號
	CreatedAt           time.Time      `json:"created_at" mapstructure:"created_at"`                       // 建檔日期
	UpdatedAt           time.Time      `json:"updated_at" mapstructure:"updated_at"`                       // 修改日期
	DeletedAt           gorm.DeletedAt `json:"-" mapstructure:"deleted_at"`                                // 使用旗標
	No                  string         `json:"no" mapstructure:"no"`                                       // 客戶編號
	Name                string         `json:"name" mapstructure:"name"`                                   // 客戶名稱
	CategoryID          uint           `json:"category_id" mapstructure:"category_id"`                     // 客戶類別關聯ID
	QuickCode           string         `json:"quick_code" mapstructure:"quick_code"`                       // 簡碼
	PhoneNo1            string         `json:"phone_no1" mapstructure:"phone_no1"`                         // 電話1
	PhoneNo2            string         `json:"phone_no2" mapstructure:"phone_no2"`                         // 電話2
	Mobilphone          string         `json:"mobilphone" mapstructure:"mobilphone"`                       // 行動電話
	Fax                 string         `json:"fax" mapstructure:"fax"`                                     // 傳真
	Contact             string         `json:"contact" mapstructure:"contact"`                             // 連絡人
	Uniform             string         `json:"uniform" mapstructure:"uniform"`                             // 統一編號
	TaxIdentificationNo string         `json:"tax_identification_no" mapstructure:"tax_identification_no"` // 稅籍編號
	BillingAddress      string         `json:"billing_address" mapstructure:"billing_address"`             // 帳單地址
	Email               string         `json:"email" mapstructure:"email"`                                 // 電子郵件
	ZipCode             string         `json:"zip_code" mapstructure:"zip_code"`                           // 郵遞區號(聯絡地址)
	ZipCode2            string         `json:"zip_code2" mapstructure:"zip_code2"`                         // 郵遞區號(發票地址)
	ZipCode3            string         `json:"zip_code3" mapstructure:"zip_code3"`                         // 郵遞區號(送貨地址)
	PostFlag            bool           `json:"post_flag" mapstructure:"post_flag"`                         // 郵寄旗標
	TransactionDate     string         `json:"transaction_date" mapstructure:"transaction_date"`           // 最近交易日
	AreaID              uint           `json:"area_id" mapstructure:"area_id"`                             // 區域別關聯ID
	DiscountID          uint           `json:"discount_id" mapstructure:"discount_id"`                     // 折扣別關聯ID
	Hint                string         `json:"hint" mapstructure:"hint"`                                   // 重要資訊
	Remark              string         `json:"remark" mapstructure:"remark"`                               // 備註
	InvoiceType         uint           `json:"invoice_type" mapstructure:"invoice_type"`                   // 開立發票方式
	InvoiceTitle        string         `json:"invoice_title" mapstructure:"invoice_title"`                 // 發票抬頭
	InvoiceAddress      string         `json:"invoice_address" mapstructure:"invoice_address"`             // 發票地址
	DeliveryAddress     string         `json:"delivery_address" mapstructure:"delivery_address"`           // 送貨地址
	DeliveryID          uint           `json:"delivery_id" mapstructure:"delivery_id"`                     // 配送類別關聯ID
	Remitter            string         `json:"remitter" mapstructure:"remitter"`                           // 匯款人
	WireAccount         string         `json:"wire_account" mapstructure:"wire_account"`                   // 匯款帳號
	BankAccountID       uint           `json:"bank_account_id" mapstructure:"bank_account_id"`             // 付款銀行關聯ID
	PaymentTerms        string         `json:"payment_terms" mapstructure:"payment_terms"`                 // 付款條件
	SalesmanID          *uint          `json:"salesman_id" mapstructure:"salesman_id"`                     // 業務員關聯ID
	AttendantID         *uint          `json:"attendant_id" mapstructure:"attendant_id"`                   // 服務員關聯ID
	MonthlyClose        string         `json:"monthly_close" mapstructure:"monthly_close"`                 // 每月結帳日
	MonthlyRequest      string         `json:"monthly_request" mapstructure:"monthly_request"`             // 每月請款日
	Prepayment          float64        `json:"prepayment" mapstructure:"prepayment"`                       // 預收金額
	OpeningAmount       float64        `json:"opening_amount" mapstructure:"opening_amount"`               // 期初金額
	CreditLine          float64        `json:"credit_line" mapstructure:"credit_line"`                     // 授信額度
	Website             string         `json:"website" mapstructure:"website"`                             // 網址
	MedicalNo           string         `json:"medical_no" mapstructure:"medical_no"`                       // 健保藥局代號
	License             string         `json:"license" mapstructure:"license"`                             // 登記證號
	Used                bool           `json:"used" mapstructure:"used"`                                   // 使用過期標
	UpdatedByID         *uint          `json:"updated_by_id" mapstructure:"updated_by_id"`                 // 修改人關聯ID
	CreatedByID         *uint          `json:"created_by_id" mapstructure:"created_by_id"`                 // 建檔人關聯ID
}

type UpdateDelivery struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 配送代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 配送名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdateDepartment struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 部門代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 部門名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdateDiscount struct {
	ID          uint      `json:"id" mapstructure:"id"`
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at"`
	No          string    `json:"no" mapstructure:"no"`
	Name        string    `json:"name" mapstructure:"name"`
	Description string    `json:"description" mapstructure:"description"`
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code"`
}

type UpdateEmployeeBonus struct {
	ID          uint      `json:"id" mapstructure:"id"`                   // 流水號
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at"`   // 建檔日期
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at"`   // 修改日期
	No          string    `json:"no" mapstructure:"no"`                   // 獎金代號(unique)
	Name        string    `json:"name" mapstructure:"name"`               // 獎金名稱
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code"`   // 簡碼
	Description string    `json:"description" mapstructure:"description"` // 說明
}

type UpdateEmployee struct {
	ID               uint      `json:"id" mapstructure:"id"`                                 // ID
	CreatedAt        time.Time `json:"created_at" mapstructure:"created_at"`                 // 建檔日期
	UpdatedAt        time.Time `json:"updated_at" mapstructure:"updated_at"`                 // 修改日期
	No               string    `json:"no" mapstructure:"no"`                                 // 員工編號
	Name             string    `json:"name" mapstructure:"name"`                             // 員工名稱
	QuickCode        string    `json:"quick_code" mapstructure:"quick_code"`                 // 簡碼
	DepartmentID     uint      `json:"department_id" mapstructure:"department_id"`           // 部門關聯ID
	PositionID       uint      `json:"position_id" mapstructure:"position_id"`               // 職位關聯ID
	IdentityNo       string    `json:"identity_no" mapstructure:"identity_no"`               // 身分證號
	Birthday         string    `json:"birthday" mapstructure:"birthday"`                     // 生日
	MailingAddress   string    `json:"mailing_address" mapstructure:"mailing_address"`       // 通訊地址
	ZipCode          string    `json:"zip_code" mapstructure:"zip_code"`                     // 郵遞區號(通訊地址的)
	ResidentAddress  string    `json:"resident_address" mapstructure:"resident_address"`     // 戶籍地址
	ContactNo        string    `json:"contact_no" mapstructure:"contact_no"`                 // 連絡電話
	PermanentPhoneNo string    `json:"permanent_phone_no" mapstructure:"permanent_phone_no"` // 戶籍電話
	Mobilphone       string    `json:"mobilphone" mapstructure:"mobilphone"`                 // 行動電話
	Email            string    `json:"email" mapstructure:"email"`                           // 電子郵件
	Remark           string    `json:"remark" mapstructure:"remark"`                         // 備註
	BonusID          uint      `json:"bonus_id" mapstructure:"bonus_id"`                     // 獎金別關聯ID
	ShiftFlag        bool      `json:"shift_flag" mapstructure:"shift_flag"`                 // 排班旗標
	Education        string    `json:"education" mapstructure:"education"`                   // 學歷
	Experience       string    `json:"experience" mapstructure:"experience"`                 // 經歷
	Gender           uint      `json:"gender" mapstructure:"gender"`                         // 性別(0.未填 1.男 2.女 3.不便透露)
	BloodType        uint      `json:"blood_type" mapstructure:"blood_type"`                 // 血型(0.未填 1.A型 2.B型 3.AB型 4.O型 5.特殊血型)
	MarriedStatus    uint      `json:"married_status" mapstructure:"married_status"`         // 婚姻狀態(0.未填 1.未婚 2.已婚 3.離婚)
	NhiIcCard        string    `json:"nhi_ic_card" mapstructure:"nhi_ic_card"`               // 健保IC號碼
	InsuranceDate    string    `json:"insurance_date" mapstructure:"insurance_date"`         // 加保日期
	InsuranceCard    string    `json:"insurance_card" mapstructure:"insurance_card"`         // 加保證號
	OnBoardDate      string    `json:"on_board_date" mapstructure:"on_board_date"`           // 到職日期
	ResignationDate  string    `json:"resignation_date" mapstructure:"resignation_date"`     // 離職日期
	DependentNo      uint      `json:"dependent_no" mapstructure:"dependent_no"`             // 扶養人數
	InsuranceNo      uint      `json:"insurance_no" mapstructure:"insurance_no"`             // 眷保人數
	BirthPlace       string    `json:"birth_place" mapstructure:"birth_place"`               // 出身地(戶籍)
	Blocked          bool      `json:"blocked" mapstructure:"blocked"`                       // 停用
	Used             bool      `json:"used" mapstructure:"used"`                             // 使用過期標
	UpdatedByID      *uint     `json:"updated_by_id" mapstructure:"updated_by_id"`           // 修改人關聯ID
	CreatedByID      *uint     `json:"created_by_id" mapstructure:"created_by_id"`           // 建檔人關聯ID
}

type UpdateItemBonus struct {
	ID          uint      `json:"id" mapstructure:"id"`                   // 流水號
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at"`   // 建檔日期
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at"`   // 修改日期
	No          string    `json:"no" mapstructure:"no"`                   // 獎金代號(unique)
	Name        string    `json:"name" mapstructure:"name"`               // 獎金名稱
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code"`   // 簡碼
	Description string    `json:"description" mapstructure:"description"` // 說明
}

type UpdateItem struct {
	ID                uint      `json:"id" mapstructure:"id"`                                   // ID
	CreatedAt         time.Time `json:"created_at" mapstructure:"created_at"`                   // 建檔日期
	UpdatedAt         time.Time `json:"updated_at" mapstructure:"updated_at"`                   // 修改日期
	No                string    `json:"no" mapstructure:"no"`                                   // 編號
	Name              string    `json:"name" mapstructure:"name"`                               // 名稱
	SubcategoryID     uint      `json:"subcategory_id" mapstructure:"subcategory_id"`           // 小類關聯ID
	SpecialCategoryID uint      `json:"special_category_id" mapstructure:"special_category_id"` // 特殊分類
	TaxType           uint      `json:"tax_type" mapstructure:"tax_type"`                       // 稅別(1.應稅內含 2.應稅外加 3.零稅率 4.免稅)
	Spec              string    `json:"spec" mapstructure:"spec"`                               // 規格說明
	QuickCode         string    `json:"quick_code" mapstructure:"quick_code"`                   // 簡碼
	SkuNo             string    `json:"sku_no" mapstructure:"sku_no"`                           // 貨號
	BonusID           uint      `json:"bonus_id" mapstructure:"bonus_id"`                       // 獎金別關聯ID
	BarcodeFlag       bool      `json:"barcode_flag" mapstructure:"barcode_flag"`               // 是否印出條碼
	AverageCost       float64   `json:"average_cost" mapstructure:"average_cost"`               // 平均成本
	WholesaleCost     float64   `json:"wholesale_cost" mapstructure:"wholesale_cost"`           // 批發成本
	UnitID            uint      `json:"unit_id" mapstructure:"unit_id"`                         // 單位關聯ID
	PromotionID       uint      `json:"promotion_id" mapstructure:"promotion_id"`               // 促銷類別關聯ID
	UpcCode1          string    `json:"upc_code1" mapstructure:"upc_code1"`                     // 國際條碼1
	UpcCode2          string    `json:"upc_code2" mapstructure:"upc_code2"`                     // 國際條碼2
	UpcCode3          string    `json:"upc_code3" mapstructure:"upc_code3"`                     // 國際條碼3
	OriginID          uint      `json:"origin_id" mapstructure:"origin_id"`                     // 產地關聯ID
	Remark            string    `json:"remark" mapstructure:"remark"`                           // 備註
	AllowSale         bool      `json:"allow_sale" mapstructure:"allow_sale"`                   // 允許銷售
	AllowPosExchange  bool      `json:"allow_pos_exchange" mapstructure:"allow_pos_exchange"`   // 允許POS退換貨
	AllowPromotion    bool      `json:"allow_promotion" mapstructure:"allow_promotion"`         // 允許促銷
	SkuFlag           bool      `json:"sku_flag" mapstructure:"sku_flag"`                       // 計算SKU
	OrderMethod       uint      `json:"order_method" mapstructure:"order_method"`               // 採購方式(1.自行(人工)訂貨-預設 2.庫存缺貨轉訂貨-自動補貨處理 3.銷售動態轉訂貨-保留尚無功能)
	AllowInquiry      bool      `json:"allow_inquiry" mapstructure:"allow_inquiry"`             // 允許詢價
	AllowRequisition  bool      `json:"allow_requisition" mapstructure:"allow_requisition"`     // 允許請購
	AllowDispatch     bool      `json:"allow_dispatch" mapstructure:"allow_dispatch"`           // 允許調撥
	AllowPricing      bool      `json:"allow_pricing" mapstructure:"allow_pricing"`             // 允許變價(開放售價)
	DeliveryMethod    uint      `json:"delivery_method" mapstructure:"delivery_method"`         // 配送方式(1.自行配送(總部物流) 2.廠商直送 3.區域經銷商配送 4.工廠配送 5.其他配送方式 6.總代理配送)
	PosProperty       uint      `json:"pos_property" mapstructure:"pos_property"`               // 銷售屬性(1.正常品 2.折現金/折價券 3.不扣庫商品 4.贈品點數商品 5.代收現金3 6.代收現金4 7.5代收現金5 8.5環保折讓 9.代用條碼 10.其他)
	DrugProperty      uint      `json:"drug_property" mapstructure:"drug_property"`             // 藥品屬性(1.非藥品 2.一般藥品 3.指示用藥 4.處方用藥)
	Turnover          float64   `json:"turnover" mapstructure:"turnover"`                       // 周轉率
	PackageType       uint      `json:"package_type" mapstructure:"package_type"`               // 包裝類型(1.基本商品 2.大包裝 3.組合包裝)
	Blocked           bool      `json:"blocked" mapstructure:"blocked"`                         // 停用
	Used              bool      `json:"used" mapstructure:"used"`                               // 使用過期標
	UpdatedByID       *uint     `json:"updated_by_id" mapstructure:"updated_by_id"`             // 修改人關聯ID
	CreatedByID       *uint     `json:"created_by_id" mapstructure:"created_by_id"`             // 建檔人關聯ID
}

type UpdateItemCategory struct {
	ID                 uint      `json:"id" mapstructure:"id"`                                       // 流水號
	CreatedAt          time.Time `json:"created_at" mapstructure:"created_at"`                       // 建檔日期
	UpdatedAt          time.Time `json:"updated_at" mapstructure:"updated_at"`                       // 修改日期
	No                 string    `json:"no" mapstructure:"no"`                                       // 類別代號(unique)
	Name               string    `json:"name" mapstructure:"name"`                                   // 類別名稱
	QuickCode          string    `json:"quick_code" mapstructure:"quick_code"`                       // 簡碼
	ItemMainCategoryID uint      `json:"item_main_category_id" mapstructure:"item_main_category_id"` // 大類關聯ID
}

type UpdateItemMainCategory struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 大類代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 大類名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdateItemSubcategory struct {
	ID             uint      `json:"id" mapstructure:"id"`                             // 流水號
	CreatedAt      time.Time `json:"created_at" mapstructure:"created_at"`             // 建檔日期
	UpdatedAt      time.Time `json:"updated_at" mapstructure:"updated_at"`             // 修改日期
	No             string    `json:"no" mapstructure:"no"`                             // 小類代號(unique)
	Name           string    `json:"name" mapstructure:"name"`                         // 小類名稱
	QuickCode      string    `json:"quick_code" mapstructure:"quick_code"`             // 簡碼
	ItemCategoryID uint      `json:"item_category_id" mapstructure:"item_category_id"` // 類別關聯ID
}

type UpdateOrigin struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 產地代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 產地名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdatePackagedItem struct {
	UpdateItem    `mapstructure:",squash"`
	BasicItemID   uint    `json:"basic_item_id" mapstructure:"basic_item_id"`   // 基本商品關聯ID
	Quantity      float64 `json:"quantity" mapstructure:"quantity"`             // 包裝量
	AllowPurchase bool    `json:"allow_purchase" mapstructure:"allow_purchase"` // 允許進貨
}

type UpdatePicture struct {
	ID            uint      `json:"id" mapstructure:"id"`                         // 流水號
	CreatedAt     time.Time `json:"created_at" mapstructure:"created_at"`         // 建檔日期
	UpdatedAt     time.Time `json:"updated_at" mapstructure:"updated_at"`         // 修改日期
	OwnerID       uint      `json:"-"`                                            // GORM關聯用欄位
	OwnerType     string    `json:"-"`                                            // GORM關聯用欄位
	Data          string    `json:"data" mapstructure:"data"`                     // Base64圖檔
	OrderPriority uint      `json:"order_priority" mapstructure:"order_priority"` // 排序優先度
}

type UpdatePosition struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 職別代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 職別名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdatePriceCategory struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	Name      string    `json:"name" mapstructure:"name"`             // 價格名稱
}

type UpdatePriceConfig struct {
	NormalPriceID    uint `json:"normal_price_id" mapstructure:"normal_price_id"`       // 正常價(預設價)關聯ID
	MemberPriceID    uint `json:"member_price_id" mapstructure:"member_price_id"`       // 會員價關聯ID
	EmployeePriceID  uint `json:"employee_price_id" mapstructure:"employee_price_id"`   // 員工價關聯ID
	VipPriceID       uint `json:"vip_price_id" mapstructure:"vip_price_id"`             // VIP價關聯ID
	WholesalePriceID uint `json:"wholesale_price_id" mapstructure:"wholesale_price_id"` // 批發價關聯ID
}

type UpdatePromotion struct {
	ID          uint      `json:"id" mapstructure:"id"`                   // 流水號
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at"`   // 建檔日期
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at"`   // 修改日期
	No          string    `json:"no" mapstructure:"no"`                   // 促銷代號(unique)
	Name        string    `json:"name" mapstructure:"name"`               // 促銷名稱
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code"`   // 簡碼
	Description string    `json:"description" mapstructure:"description"` // 說明
}

type UpdateProviderCategory struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 廠商類別代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 廠商類別名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type UpdateProvider struct {
	ID              uint      `json:"id" mapstructure:"id"`                             // ID
	CreatedAt       time.Time `json:"created_at" mapstructure:"created_at"`             // 建檔日期
	UpdatedAt       time.Time `json:"updated_at" mapstructure:"updated_at"`             // 修改日期
	No              string    `json:"no" mapstructure:"no"`                             // 廠商編號
	Name            string    `json:"name" mapstructure:"name"`                         // 廠商名稱
	CategoryID      uint      `json:"category_id" mapstructure:"category_id"`           // 廠商類別關聯ID
	QuickCode       string    `json:"quick_code" mapstructure:"quick_code"`             // 簡碼
	Uniform         string    `json:"uniform" mapstructure:"uniform"`                   // 統一編號
	InvoiceTitle    string    `json:"invoice_title" mapstructure:"invoice_title"`       // 發票抬頭
	InvoiceAddress  string    `json:"invoice_address" mapstructure:"invoice_address"`   // 發票地址
	Prepayment      float64   `json:"prepayment" mapstructure:"prepayment"`             // 預收金額
	ZipCode         string    `json:"zip_code" mapstructure:"zip_code"`                 // 郵遞區號
	PostFlag        bool      `json:"post_flag" mapstructure:"post_flag"`               // 郵寄旗標
	Address         string    `json:"address" mapstructure:"address"`                   // 公司地址
	Contact         string    `json:"contact" mapstructure:"contact"`                   // 聯絡人
	PhoneNo1        string    `json:"phone_no1" mapstructure:"phone_no1"`               // 電話1
	PhoneNo2        string    `json:"phone_no2" mapstructure:"phone_no2"`               // 電話2
	Mobilphone      string    `json:"mobilphone" mapstructure:"mobilphone"`             // 行動電話
	Fax             string    `json:"fax" mapstructure:"fax"`                           // 傳真
	Email           string    `json:"email" mapstructure:"email"`                       // 電子郵件
	Website         string    `json:"website" mapstructure:"website"`                   // 公司網址
	WireAccount     string    `json:"wire_account" mapstructure:"wire_account"`         // 匯款帳號
	BankAccountID   uint      `json:"bank_account_id" mapstructure:"bank_account_id"`   // 付款銀行關聯ID
	PaymentTerms    string    `json:"payment_terms" mapstructure:"payment_terms"`       // 付款條件
	License         string    `json:"license" mapstructure:"license"`                   // 管藥登記證號
	TransactionDate string    `json:"transaction_date" mapstructure:"transaction_date"` // 最近交易日
	OpeningAmount   float64   `json:"opening_amount" mapstructure:"opening_amount"`     // 期初金額
	MonthlyClose    string    `json:"monthly_close" mapstructure:"monthly_close"`       // 每月結帳日
	MonthlyRequest  string    `json:"monthly_request" mapstructure:"monthly_request"`   // 每月請款日
	Remark          string    `json:"remark" mapstructure:"remark"`                     // 備註
	Hint            string    `json:"hint" mapstructure:"hint"`                         // 重要資訊
	Blocked         bool      `json:"blocked" mapstructure:"blocked"`                   // 停用
	Used            bool      `json:"used" mapstructure:"used"`                         // 使用過期標
	UpdatedByID     *uint     `json:"updated_by_id" mapstructure:"updated_by_id"`       // 修改人關聯ID
	CreatedByID     *uint     `json:"created_by_id" mapstructure:"created_by_id"`       // 建檔人關聯ID
}

type UpdateStoreItem struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	StoreID   uint
	ItemID    uint
}

type UpdateUnit struct {
	ID        uint      `json:"id" mapstructure:"id"`                 // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"` // 修改日期
	No        string    `json:"no" mapstructure:"no"`                 // 單位代號(unique)
	Name      string    `json:"name" mapstructure:"name"`             // 單位名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code"` // 簡碼
}

type StoreUpdated struct {
	ID                uint      `json:"id" mapstructure:"id"`                                   // 流水號
	CreatedAt         time.Time `json:"created_at" mapstructure:"created_at"`                   // 建檔日期
	UpdatedAt         time.Time `json:"updated_at" mapstructure:"updated_at"`                   // 修改日期
	No                string    `json:"no" mapstructure:"no"`                                   // 門市代號
	Name              string    `json:"name" mapstructure:"name"`                               // 門市名稱
	QuickCode         string    `json:"quick_code" mapstructure:"quick_code"`                   // 簡碼
	AreaID            uint      `json:"area_id" mapstructure:"area_id"`                         // 區域別id
	StoreType         uint      `json:"store_type" mapstructure:"store_type"`                   // 店系類別
	Uniform           string    `json:"uniform" mapstructure:"uniform"`                         // 統一編號
	InvoiceTitle      string    `json:"invoice_title" mapstructure:"invoice_title"`             // 發票抬頭
	Contact           string    `json:"contact" mapstructure:"contact"`                         // 聯絡人
	IdentityNo        string    `json:"identity_no" mapstructure:"identity_no"`                 // 身分證
	PhoneNo1          string    `json:"phone_no1" mapstructure:"phone_no1"`                     // 電話1
	PhoneNo2          string    `json:"phone_no2" mapstructure:"phone_no2"`                     // 電話2
	Fax               string    `json:"fax" mapstructure:"fax"`                                 // 傳真
	BillingAddress    string    `json:"billing_address" mapstructure:"billing_address"`         // 帳單地址
	InvoiceAddress    string    `json:"invoice_address" mapstructure:"invoice_address"`         // 公司發票地址
	DeliveryAddress   string    `json:"delivery_address" mapstructure:"delivery_address"`       // 送貨地址
	BillingZipCode    string    `json:"billing_zip_code" mapstructure:"billing_zip_code"`       // 郵遞區號(帳單地址)
	InvoiceZipCode    string    `json:"invoice_zip_code" mapstructure:"invoice_zip_code"`       // 郵遞區號(公司發票地址)
	DeliveryZipCode   string    `json:"delivery_zip_code" mapstructure:"delivery_zip_code"`     // 郵遞區號(送貨地址)
	Email             string    `json:"email" mapstructure:"email"`                             // 電子郵件
	BusinessStartDate string    `json:"business_start_date" mapstructure:"business_start_date"` // 開始營業日
	BusinessEndDate   string    `json:"business_end_date" mapstructure:"business_end_date"`     // 結束營業日
	Remark            string    `json:"remark" mapstructure:"remark"`                           // 備註
	Blocked           bool      `json:"blocked" mapstructure:"blocked"`                         // 停用
	UpdatedByID       *uint     `json:"updated_by_id" mapstructure:"updated_by_id"`             // 修改人關聯ID
	CreatedByID       *uint     `json:"created_by_id" mapstructure:"created_by_id"`             // 建檔人關聯ID
}

type WarehouseUpdated struct {
	ID          uint      `json:"id" mapstructure:"id"`                       // 流水號
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at"`       // 建檔日期
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at"`       // 修改日期
	No          string    `json:"no" mapstructure:"no"`                       // 區域代號(unique)
	Name        string    `json:"name" mapstructure:"name"`                   // 區域名稱
	StoreID     uint      `json:"store_id" mapstructure:"store_id"`           // 職位關聯ID
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code"`       // 簡碼
	Remark      string    `json:"remark" mapstructure:"remark"`               // 備註
	Blocked     bool      `json:"blocked" mapstructure:"blocked"`             // 停用
	UpdatedByID *uint     `json:"updated_by_id" mapstructure:"updated_by_id"` // 修改人關聯ID
	CreatedByID *uint     `json:"created_by_id" mapstructure:"created_by_id"` // 建檔人關聯ID
}

type DeleteArea struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteBankAccount struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteBasicItem struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteCustomerCategory struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteCustomer struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteDelivery struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteDepartment struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteDiscount struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteEmployeeBonus struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteEmployee struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteItemBonus struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteItemCategory struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteItemMainCategory struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteItemSubcategory struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteOrigin struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeletePackagedItem struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeletePicture struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeletePosition struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeletePriceCategory struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeletePriceConfig struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeletePromotion struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteProviderCategory struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteProvider struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteStoreItem struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteStore struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteUnit struct {
	ID uint `json:"id" mapstructure:"id"`
}
