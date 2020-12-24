package events

import (
	"time"
)

type UpdateArea struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 區域代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 區域名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type UpdateBankAccount struct {
	ID             uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                          // ID
	CreatedAt      time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                   // 建檔日期
	UpdatedAt      time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                   // 修改日期
	No             string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"`    // 銀行帳號
	Name           string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                     // 銀行名稱
	QuickCode      string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                   // 簡碼
	AccountType    uint      `json:"account_type" mapstructure:"account_type" gorm:"Column:account_type"`             // 帳戶別(1.甲存 2.乙存 3.刷卡銀行)
	CheckType      uint      `json:"check_type" mapstructure:"check_type" gorm:"Column:check_type"`                   // 支票別(1.本票 2.支票)
	SubjectID      uint      `json:"subject_id" mapstructure:"subject_id" gorm:"Column:subject_id"`                   // 會計科目
	PhoneNo        string    `json:"phone_no" mapstructure:"phone_no" gorm:"Column:phone_no"`                         // 電話
	DepositAmount  float64   `json:"deposit_amount" mapstructure:"deposit_amount" gorm:"Column:deposit_amount"`       // 存款金額
	MarginOfSafety float64   `json:"margin_of_safety" mapstructure:"margin_of_safety" gorm:"Column:margin_of_safety"` // 安全餘額
	Blocked        bool      `json:"blocked" mapstructure:"blocked" gorm:"Column:blocked"`                            // 停用
	Used           bool      `json:"used" mapstructure:"used" gorm:"Column:used"`                                     // 使用過期標
	UpdatedByID    *uint     `json:"updated_by_id" mapstructure:"updated_by_id" gorm:"Column:updated_by_id"`          // 修改人關聯ID
	CreatedByID    *uint     `json:"created_by_id" mapstructure:"created_by_id" gorm:"Column:created_by_id"`          // 建檔人關聯ID
}

type UpdateCustomerCategory struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 客戶類別代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 客戶類別名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type UpdateCustomer struct {
	ID              uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                          // 流水號
	CreatedAt       time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                   // 建檔日期
	UpdatedAt       time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                   // 修改日期
	No              string    `json:"no" mapstructure:"no" gorm:"Column:no"`                                           // 客戶編號
	InputNo         string    `json:"input_no" mapstructure:"input_no" gorm:"Column:input_no"`                         // 門市代號
	CategoryID      uint      `json:"category_id" mapstructure:"category_id" gorm:"Column:category_id"`                // 客戶類別
	Name            string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                     // 客戶名稱
	UseIndex        string    `json:"use_index" mapstructure:"use_index" gorm:"Column:use_index"`                      // 簡碼
	Telephone1      string    `json:"telephone1" mapstructure:"telephone1" gorm:"Column:telephone1"`                   // 電話1
	Telephone2      string    `json:"telephone2" mapstructure:"telephone2" gorm:"Column:telephone2"`                   // 電話2
	Mobilphone      string    `json:"mobilphone" mapstructure:"mobilphone" gorm:"Column:mobilphone"`                   // 行動電話
	Fax             string    `json:"fax" mapstructure:"fax" gorm:"Column:fax"`                                        // 傳真
	Contact         string    `json:"contact" mapstructure:"contact" gorm:"Column:contact"`                            // 連絡人
	Uniform         string    `json:"uniform" mapstructure:"uniform" gorm:"Column:uniform"`                            // 統一編號
	BillingAddress  string    `json:"billing_address" mapstructure:"billing_address" gorm:"Column:billing_address"`    // 帳單地址
	ZipCode         string    `json:"zip_code" mapstructure:"zip_code" gorm:"Column:zip_code"`                         // 郵遞區號
	Email           string    `json:"email" mapstructure:"email" gorm:"Column:email"`                                  // 電子郵件
	PostFlag        bool      `json:"post_flag" mapstructure:"post_flag" gorm:"Column:post_flag"`                      // 郵寄標籤
	TransactionDate string    `json:"transaction_date" mapstructure:"transaction_date" gorm:"Column:transaction_date"` // 最近交易日期
	AreaID          uint      `json:"area_id" mapstructure:"area_id" gorm:"Column:area_id"`                            // 區域別關聯ID
	DiscountID      uint      `json:"discount_id" mapstructure:"discount_id" gorm:"Column:discount_id"`                // 折扣別關聯ID
	Remark2         string    `json:"remark2" mapstructure:"remark2" gorm:"Column:remark2"`                            // 重要資訊
	Remark          string    `json:"remark" mapstructure:"remark" gorm:"Column:remark"`                               // 備 註
	InvoiceTitle    string    `json:"invoice_title" mapstructure:"invoice_title" gorm:"Column:invoice_title"`          // 發票抬頭
	InvoiceAddress  string    `json:"invoice_address" mapstructure:"invoice_address" gorm:"Column:invoice_address"`    // 發票地址
	DeliveryAddress string    `json:"delivery_address" mapstructure:"delivery_address" gorm:"Column:delivery_address"` // 送貨地址
	DeliveryID      uint      `json:"delivery_id" mapstructure:"delivery_id" gorm:"Column:delivery_id"`                // 配送類別
	MonthlyClose    string    `json:"monthly_close" mapstructure:"monthly_close" gorm:"Column:monthly_close"`          // 每月結帳日(每月歸帳日)
	MonthlyRequest  string    `json:"monthly_request" mapstructure:"monthly_request" gorm:"Column:monthly_request"`    // 每月請款日(會計請款日)
	OpeningAmount   float64   `json:"opening_amount" mapstructure:"opening_amount" gorm:"Column:opening_amount"`       // 期初金額
	CreditLine      float64   `json:"credit_line" mapstructure:"credit_line" gorm:"Column:credit_line"`                // 授信額度
	DeliveryRoute   string    `json:"delivery_route" mapstructure:"delivery_route" gorm:"Column:delivery_route"`       // 配送路線
	Url             string    `json:"url" mapstructure:"url" gorm:"Column:url"`                                        // 網 址
	License         string    `json:"license" mapstructure:"license" gorm:"Column:license"`                            // 登記證號
	Used            bool      `json:"used" mapstructure:"used" gorm:"Column:used"`                                     // 使用過期標
	UpdatedByID     *uint     `json:"updated_by_id" mapstructure:"updated_by_id" gorm:"Column:updated_by_id"`          // 修改人關聯ID
	CreatedByID     *uint     `json:"created_by_id" mapstructure:"created_by_id" gorm:"Column:created_by_id"`          // 建檔人關聯ID
}

type UpdateDelivery struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 配送代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 配送名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type UpdateDepartment struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 部門代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 部門名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type UpdateDiscount struct {
	ID          uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`
	No          string    `json:"no" mapstructure:"no" gorm:"Column:no"`
	Name        string    `json:"name" mapstructure:"name" gorm:"Column:name"`
	Description string    `json:"description" mapstructure:"description" gorm:"Column:description"`
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`
}

type UpdateEmployeeBonus struct {
	ID          uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No          string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 獎金代號(unique)
	Name        string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 獎金名稱
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
	Description string    `json:"description" mapstructure:"description" gorm:"Column:description"`             // 說明
}

type UpdateEmployee struct {
	ID               uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                                // ID
	CreatedAt        time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                         // 建檔日期
	UpdatedAt        time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                         // 修改日期
	No               string    `json:"no" mapstructure:"no" gorm:"Column:no"`                                                 // 員工編號
	Name             string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                           // 員工名稱
	QuickCode        string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                         // 簡碼
	DepartmentID     uint      `json:"department_id" mapstructure:"department_id" gorm:"Column:department_id"`                // 部門關聯ID
	PositionID       uint      `json:"position_id" mapstructure:"position_id" gorm:"Column:position_id"`                      // 職位關聯ID
	IdentityNo       string    `json:"identity_no" mapstructure:"identity_no" gorm:"Column:identity_no"`                      // 身分證號
	Birthday         string    `json:"birthday" mapstructure:"birthday" gorm:"Column:birthday"`                               // 生日
	MailingAddress   string    `json:"mailing_address" mapstructure:"mailing_address" gorm:"Column:mailing_address"`          // 通訊地址
	ZipCode          string    `json:"zip_code" mapstructure:"zip_code" gorm:"Column:zip_code"`                               // 郵遞區號(通訊地址的)
	ResidentAddress  string    `json:"resident_address" mapstructure:"resident_address" gorm:"Column:resident_address"`       // 戶籍地址
	ContactNo        string    `json:"contact_no" mapstructure:"contact_no" gorm:"Column:contact_no"`                         // 連絡電話
	PermanentPhoneNo string    `json:"permanent_phone_no" mapstructure:"permanent_phone_no" gorm:"Column:permanent_phone_no"` // 戶籍電話
	Mobilphone       string    `json:"mobilphone" mapstructure:"mobilphone" gorm:"Column:mobilphone"`                         // 行動電話
	Email            string    `json:"email" mapstructure:"email" gorm:"Column:email"`                                        // 電子郵件
	Remark           string    `json:"remark" mapstructure:"remark" gorm:"Column:remark"`                                     // 備註
	BonusID          uint      `json:"bonus_id" mapstructure:"bonus_id" gorm:"Column:bonus_id"`                               // 獎金別關聯ID
	ShiftFlag        bool      `json:"shift_flag" mapstructure:"shift_flag" gorm:"Column:shift_flag"`                         // 排班旗標
	Education        string    `json:"education" mapstructure:"education" gorm:"Column:education"`                            // 學歷
	Experience       string    `json:"experience" mapstructure:"experience" gorm:"Column:experience"`                         // 經歷
	Gender           uint      `json:"gender" mapstructure:"gender" gorm:"Column:gender"`                                     // 性別(0.未填 1.男 2.女 3.不便透露)
	BloodType        uint      `json:"blood_type" mapstructure:"blood_type" gorm:"Column:blood_type"`                         // 血型(0.未填 1.A型 2.B型 3.AB型 4.O型 5.特殊血型)
	MarriedStatus    uint      `json:"married_status" mapstructure:"married_status" gorm:"Column:married_status"`             // 婚姻狀態(0.未填 1.未婚 2.已婚 3.離婚)
	NhiIcCard        string    `json:"nhi_ic_card" mapstructure:"nhi_ic_card" gorm:"Column:nhi_ic_card"`                      // 健保IC號碼
	InsuranceDate    string    `json:"insurance_date" mapstructure:"insurance_date" gorm:"Column:insurance_date"`             // 加保日期
	InsuranceCard    string    `json:"insurance_card" mapstructure:"insurance_card" gorm:"Column:insurance_card"`             // 加保證號
	OnBoardDate      string    `json:"on_board_date" mapstructure:"on_board_date" gorm:"Column:on_board_date"`                // 到職日期
	ResignationDate  string    `json:"resignation_date" mapstructure:"resignation_date" gorm:"Column:resignation_date"`       // 離職日期
	DependentNo      uint      `json:"dependent_no" mapstructure:"dependent_no" gorm:"Column:dependent_no"`                   // 扶養人數
	InsuranceNo      uint      `json:"insurance_no" mapstructure:"insurance_no" gorm:"Column:insurance_no"`                   // 眷保人數
	BirthPlace       string    `json:"birth_place" mapstructure:"birth_place" gorm:"Column:birth_place"`                      // 出身地(戶籍)
	Blocked          bool      `json:"blocked" mapstructure:"blocked" gorm:"Column:blocked"`                                  // 停用
	Used             bool      `json:"used" mapstructure:"used" gorm:"Column:used"`                                           // 使用過期標
	UpdatedByID      *uint     `json:"updated_by_id" mapstructure:"updated_by_id" gorm:"Column:updated_by_id"`                // 修改人關聯ID
	CreatedByID      *uint     `json:"created_by_id" mapstructure:"created_by_id" gorm:"Column:created_by_id"`                // 建檔人關聯ID
}

type UpdateItemBonus struct {
	ID          uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No          string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 獎金代號(unique)
	Name        string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 獎金名稱
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
	Description string    `json:"description" mapstructure:"description" gorm:"Column:description"`             // 說明
}

type UpdateItem struct {
	ID                uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                                   // ID
	CreatedAt         time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                            // 建檔日期
	UpdatedAt         time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                            // 修改日期
	No                string    `json:"no" mapstructure:"no" gorm:"Column:no"`                                                    // 編號
	MainCategoryID    uint      `json:"main_category_id" mapstructure:"main_category_id" gorm:"Column:main_category_id"`          // 大類關聯ID
	CategoryID        uint      `json:"category_id" mapstructure:"category_id" gorm:"Column:category_id"`                         // 中類關聯ID
	SubcategoryID     uint      `json:"subcategory_id" mapstructure:"subcategory_id" gorm:"Column:subcategory_id"`                // 小類關聯ID
	SpecialCategoryID uint      `json:"special_category_id" mapstructure:"special_category_id" gorm:"Column:special_category_id"` // 特殊分類
	TaxType           uint      `json:"tax_type" mapstructure:"tax_type" gorm:"Column:tax_type"`                                  // 稅別(1.應稅內含 2.應稅外加 3.零稅率 4.免稅)
	Name              string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                              // 名稱
	Spec              string    `json:"spec" mapstructure:"spec" gorm:"Column:spec"`                                              // 規格說明
	QuickCode         string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                            // 簡碼
	SkuNo             string    `json:"sku_no" mapstructure:"sku_no" gorm:"Column:sku_no"`                                        // 貨號
	BonusID           uint      `json:"bonus_id" mapstructure:"bonus_id" gorm:"Column:bonus_id"`                                  // 獎金別關聯ID
	BarcodeFlag       bool      `json:"barcode_flag" mapstructure:"barcode_flag" gorm:"Column:barcode_flag"`                      // 是否印出條碼
	ProviderID        uint      `json:"provider_id" mapstructure:"provider_id" gorm:"Column:provider_id"`                         // 主供應商關聯ID
	OnHandQty         float64   `json:"on_hand_qty" mapstructure:"on_hand_qty" gorm:"Column:on_hand_qty"`                         // 現有存量
	InitQty           float64   `json:"init_qty" mapstructure:"init_qty" gorm:"Column:init_qty"`                                  // 期初數量
	InitCost          float64   `json:"init_cost" mapstructure:"init_cost" gorm:"Column:init_cost"`                               // 期初成本
	AverageCost       float64   `json:"average_cost" mapstructure:"average_cost" gorm:"Column:average_cost"`                      // 平均成本
	WholesaleCost     float64   `json:"wholesale_cost" mapstructure:"wholesale_cost" gorm:"Column:wholesale_cost"`                // 批發成本
	UnitID            uint      `json:"unit_id" mapstructure:"unit_id" gorm:"Column:unit_id"`                                     // 單位關聯ID
	PromotionID       uint      `json:"promotion_id" mapstructure:"promotion_id" gorm:"Column:promotion_id"`                      // 促銷類別關聯ID
	BasicQty          float64   `json:"basic_qty" mapstructure:"basic_qty" gorm:"Column:basic_qty"`                               // 基本採購量
	ShipDate          string    `json:"ship_date" mapstructure:"ship_date" gorm:"Column:ship_date"`                               // 出貨日期
	PurchaseDate      string    `json:"purchase_date" mapstructure:"purchase_date" gorm:"Column:purchase_date"`                   // 進貨日期
	LicenseType       uint      `json:"license_type" mapstructure:"license_type" gorm:"Column:license_type"`                      // 許可證類別(0.未填入 1.衛署藥製 2.衛署藥輸 3.衛署 4.內衛藥製 5.內衛藥輸 6.內衛)
	LicenseNo         string    `json:"license_no" mapstructure:"license_no" gorm:"Column:license_no"`                            // 許可證字號
	UpcCode1          string    `json:"upc_code1" mapstructure:"upc_code1" gorm:"Column:upc_code1"`                               // 國際條碼1
	UpcCode2          string    `json:"upc_code2" mapstructure:"upc_code2" gorm:"Column:upc_code2"`                               // 國際條碼2
	UpcCode3          string    `json:"upc_code3" mapstructure:"upc_code3" gorm:"Column:upc_code3"`                               // 國際條碼3
	OriginID          uint      `json:"origin_id" mapstructure:"origin_id" gorm:"Column:origin_id"`                               // 產地關聯ID
	Remark            string    `json:"remark" mapstructure:"remark" gorm:"Column:remark"`                                        // 備註
	AllowPurchase     bool      `json:"allow_purchase" mapstructure:"allow_purchase" gorm:"Column:allow_purchase"`                // 允許進貨
	AllowSale         bool      `json:"allow_sale" mapstructure:"allow_sale" gorm:"Column:allow_sale"`                            // 允許銷售
	AllowInventory    bool      `json:"allow_inventory" mapstructure:"allow_inventory" gorm:"Column:allow_inventory"`             // 允許盤點
	AllowPosExchange  bool      `json:"allow_pos_exchange" mapstructure:"allow_pos_exchange" gorm:"Column:allow_pos_exchange"`    // 允許POS退換貨
	AllowPromotion    bool      `json:"allow_promotion" mapstructure:"allow_promotion" gorm:"Column:allow_promotion"`             // 允許促銷
	SkuFlag           bool      `json:"sku_flag" mapstructure:"sku_flag" gorm:"Column:sku_flag"`                                  // 計算SKU
	OrderMethod       uint      `json:"order_method" mapstructure:"order_method" gorm:"Column:order_method"`                      // 採購方式(1.自行(人工)訂貨-預設 2.庫存缺貨轉訂貨-自動補貨處理 3.銷售動態轉訂貨-保留尚無功能)
	AllowInquiry      bool      `json:"allow_inquiry" mapstructure:"allow_inquiry" gorm:"Column:allow_inquiry"`                   // 允許詢價
	AllowRequisition  bool      `json:"allow_requisition" mapstructure:"allow_requisition" gorm:"Column:allow_requisition"`       // 允許請購
	AllowDispatch     bool      `json:"allow_dispatch" mapstructure:"allow_dispatch" gorm:"Column:allow_dispatch"`                // 允許調撥
	AllowPricing      bool      `json:"allow_pricing" mapstructure:"allow_pricing" gorm:"Column:allow_pricing"`                   // 允許變價(開放售價)
	DeliveryMethod    uint      `json:"delivery_method" mapstructure:"delivery_method" gorm:"Column:delivery_method"`             // 配送方式(1.自行配送(總部物流) 2.廠商直送 3.區域經銷商配送 4.工廠配送 5.其他配送方式 6.總代理配送)
	PosProperty       uint      `json:"pos_property" mapstructure:"pos_property" gorm:"Column:pos_property"`                      // 銷售屬性(1.正常品 2.折現金/折價券 3.不扣庫商品 4.贈品點數商品 5.代收現金3 6.代收現金4 7.5代收現金5 8.5環保折讓 9.代用條碼 10.其他)
	DrugProperty      uint      `json:"drug_property" mapstructure:"drug_property" gorm:"Column:drug_property"`                   // 藥品屬性(1.非藥品 2.一般藥品 3.指示用藥 4.處方用藥)
	Turnover          float64   `json:"turnover" mapstructure:"turnover" gorm:"Column:turnover"`                                  // 周轉率
	Blocked           bool      `json:"blocked" mapstructure:"blocked" gorm:"Column:blocked"`                                     // 停用
	Used              bool      `json:"used" mapstructure:"used" gorm:"Column:used"`                                              // 使用過期標
	UpdatedByID       *uint     `json:"updated_by_id" mapstructure:"updated_by_id" gorm:"Column:updated_by_id"`                   // 修改人關聯ID
	CreatedByID       *uint     `json:"created_by_id" mapstructure:"created_by_id" gorm:"Column:created_by_id"`                   // 建檔人關聯ID
}

type UpdateItemCategory struct {
	ID          uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`
	No          string    `json:"no" mapstructure:"no" gorm:"Column:no"`
	Name        string    `json:"name" mapstructure:"name" gorm:"Column:name"`
	Description string    `json:"description" mapstructure:"description" gorm:"Column:description"`
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`
}

type UpdateItemMainCategory struct {
	ID          uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`
	No          string    `json:"no" mapstructure:"no" gorm:"Column:no"`
	Name        string    `json:"name" mapstructure:"name" gorm:"Column:name"`
	Description string    `json:"description" mapstructure:"description" gorm:"Column:description"`
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`
}

type UpdateItemSubcategory struct {
	ID          uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`
	No          string    `json:"no" mapstructure:"no" gorm:"Column:no"`
	Name        string    `json:"name" mapstructure:"name" gorm:"Column:name"`
	Description string    `json:"description" mapstructure:"description" gorm:"Column:description"`
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`
}

type UpdateOrigin struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 產地代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 產地名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type UpdatePicture struct {
	ID            uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                    // 流水號
	CreatedAt     time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`             // 建檔日期
	UpdatedAt     time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`             // 修改日期
	OwnerID       uint      `json:"-"`                                                                         // GORM關聯用欄位
	OwnerType     string    `json:"-"`                                                                         // GORM關聯用欄位
	Data          string    `json:"data" mapstructure:"data" gorm:"Column:data;type:string"`                   // Base64圖檔
	OrderPriority uint      `json:"order_priority" mapstructure:"order_priority" gorm:"Column:order_priority"` // 排序優先度
}

type UpdatePosition struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 職別代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 職別名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type UpdatePriceCategory struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                             // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                      // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                      // 修改日期
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name;index:,unique,where:deleted_at IS NULL"` // 價格名稱
}

type UpdatePriceConfig struct {
	NormalPriceID    uint `json:"normal_price_id" mapstructure:"normal_price_id" gorm:"Column:normal_price_id"`          // 正常價(預設價)關聯ID
	MemberPriceID    uint `json:"member_price_id" mapstructure:"member_price_id" gorm:"Column:member_price_id"`          // 會員價關聯ID
	EmployeePriceID  uint `json:"employee_price_id" mapstructure:"employee_price_id" gorm:"Column:employee_price_id"`    // 員工價關聯ID
	VipPriceID       uint `json:"vip_price_id" mapstructure:"vip_price_id" gorm:"Column:vip_price_id"`                   // VIP價關聯ID
	WholesalePriceID uint `json:"wholesale_price_id" mapstructure:"wholesale_price_id" gorm:"Column:wholesale_price_id"` // 批發價關聯ID
}

type UpdatePromotion struct {
	ID          uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No          string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 促銷代號(unique)
	Name        string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 促銷名稱
	QuickCode   string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
	Description string    `json:"description" mapstructure:"description" gorm:"Column:description"`             // 說明
}

type UpdateProviderCategory struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 廠商類別代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 廠商類別名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type UpdateProvider struct {
	ID              uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                          // ID
	CreatedAt       time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                   // 建檔日期
	UpdatedAt       time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                   // 修改日期
	No              string    `json:"no" mapstructure:"no" gorm:"Column:no"`                                           // 廠商編號
	Name            string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                     // 廠商名稱
	CategoryID      uint      `json:"category_id" mapstructure:"category_id" gorm:"Column:category_id"`                // 廠商類別關聯ID
	QuickCode       string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                   // 簡碼
	Uniform         string    `json:"uniform" mapstructure:"uniform" gorm:"Column:uniform"`                            // 統一編號
	InvoiceTitle    string    `json:"invoice_title" mapstructure:"invoice_title" gorm:"Column:invoice_title"`          // 發票抬頭
	InvoiceAddress  string    `json:"invoice_address" mapstructure:"invoice_address" gorm:"Column:invoice_address"`    // 發票地址
	Prepayment      float64   `json:"prepayment" mapstructure:"prepayment" gorm:"Column:prepayment"`                   // 預收金額
	ZipCode         string    `json:"zip_code" mapstructure:"zip_code" gorm:"Column:zip_code"`                         // 郵遞區號
	PostFlag        bool      `json:"post_flag" mapstructure:"post_flag" gorm:"Column:post_flag"`                      // 郵寄旗標
	Address         string    `json:"address" mapstructure:"address" gorm:"Column:address"`                            // 公司地址
	Contact         string    `json:"contact" mapstructure:"contact" gorm:"Column:contact"`                            // 聯絡人
	PhoneNo1        string    `json:"phone_no1" mapstructure:"phone_no1" gorm:"Column:phone_no1"`                      // 電話1
	PhoneNo2        string    `json:"phone_no2" mapstructure:"phone_no2" gorm:"Column:phone_no2"`                      // 電話2
	Mobilphone      string    `json:"mobilphone" mapstructure:"mobilphone" gorm:"Column:mobilphone"`                   // 行動電話
	Fax             string    `json:"fax" mapstructure:"fax" gorm:"Column:fax"`                                        // 傳真
	Email           string    `json:"email" mapstructure:"email" gorm:"Column:email"`                                  // 電子郵件
	Website         string    `json:"website" mapstructure:"website" gorm:"Column:website"`                            // 公司網址
	WireAccount     string    `json:"wire_account" mapstructure:"wire_account" gorm:"Column:wire_account"`             // 匯款帳號
	BankAccountID   uint      `json:"bank_account_id" mapstructure:"bank_account_id" gorm:"Column:bank_account_id"`    // 付款銀行關聯ID
	PaymentTerms    string    `json:"payment_terms" mapstructure:"payment_terms" gorm:"Column:payment_terms"`          // 付款條件
	License         string    `json:"license" mapstructure:"license" gorm:"Column:license"`                            // 管藥登記證號
	TransactionDate string    `json:"transaction_date" mapstructure:"transaction_date" gorm:"Column:transaction_date"` // 最近交易日
	OpeningAmount   float64   `json:"opening_amount" mapstructure:"opening_amount" gorm:"Column:opening_amount"`       // 期初金額
	MonthlyClose    string    `json:"monthly_close" mapstructure:"monthly_close" gorm:"Column:monthly_close"`          // 每月結帳日
	MonthlyRequest  string    `json:"monthly_request" mapstructure:"monthly_request" gorm:"Column:monthly_request"`    // 每月請款日
	Remark          string    `json:"remark" mapstructure:"remark" gorm:"Column:remark"`                               // 備註
	Hint            string    `json:"hint" mapstructure:"hint" gorm:"Column:hint"`                                     // 重要資訊
	Blocked         bool      `json:"blocked" mapstructure:"blocked" gorm:"Column:blocked"`                            // 停用
	Used            bool      `json:"used" mapstructure:"used" gorm:"Column:used"`                                     // 使用過期標
	UpdatedByID     *uint     `json:"updated_by_id" mapstructure:"updated_by_id" gorm:"Column:updated_by_id"`          // 修改人關聯ID
	CreatedByID     *uint     `json:"created_by_id" mapstructure:"created_by_id" gorm:"Column:created_by_id"`          // 建檔人關聯ID
}

type UpdateStoreItem struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                        // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"` // 修改日期
	StoreID   uint
	ItemID    uint
}

type UpdateStore struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 門市代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 門市名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type UpdateUnit struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	No        string    `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 單位代號(unique)
	Name      string    `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 單位名稱
	QuickCode string    `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}

type DeleteArea struct {
	ID uint `json:"id" mapstructure:"id"`
}

type DeleteBankAccount struct {
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

type DeleteItem struct {
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
