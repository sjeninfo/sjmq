package events

import (
	"time"
)

type UserDeleted struct {
	ID uint `json:"id"`
}

type UserUpdated struct {
	ID               uint              `json:"id"`                // 流水號
	CreatedAt        time.Time         `json:"created_at"`        // 建檔日期
	UpdatedAt        time.Time         `json:"updated_at"`        // 修改日期
	Account          string            `json:"account"`           // 帳號
	Password         string            `json:"password"`          // 密碼
	VerifiedStatus   uint              `json:"verified_status"`   // 認證狀態
	Identity         string            `json:"identity"`          // 身分別
	BlockedAt        *time.Time        `json:"blocked_at"`        // 停用時間
	SjenEmployee     *SjenEmployee     `json:"sjen_employee"`     // 先勁員工資料
	Headquarters     *Headquarters     `json:"headquarters"`      // 總部資料
	HeadquartersEmp  *HeadquartersEmp  `json:"headquarters_emp"`  // 總部員工資料
	Pharmacy         *Pharmacy         `json:"pharmacy"`          // 藥局資料
	PharmacyEmployee *PharmacyEmployee `json:"pharmacy_employee"` // 藥局員工資料
	Consumer         *Consumer         `json:"consumer"`          // 消費者資料
}

type SjenEmployee struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

type Headquarters struct {
	UserID uint   `json:"user_id"` // 用戶主檔關聯ID
	Name   string `json:"name"`    // 姓名
}
type HeadquartersEmp struct {
	UserID             uint   `json:"user_id" gorm:"primaryKey;Column:user_id"`                // 用戶主檔關聯ID
	Name               string `json:"name" gorm:"Column:name;size:20"`                         // 姓名
	HeadquartersUserID uint   `json:"headquarters_user_id" gorm:"Column:headquarters_user_id"` // 藥局關聯ID
}
type Pharmacy struct {
	UserID             uint    `json:"user_id"`
	HeadquartersUserID uint    `json:"headquarters_user_id"`
	Name               string  `json:"name"`
	Remark             string  `json:"remark"`  // 備註
	Address            string  `json:"address"` // 地址
	Phone              string  `json:"phone"`   // 電話
	Email              string  `json:"email"`
	Latitude           float64 `json:"latitude"`  // 緯度
	Longitude          float64 `json:"longitude"` // 經度
}
type PharmacyEmployee struct {
	UserID         uint   `json:"user_id"`
	Name           string `json:"name"`
	PharmacyUserID uint   `json:"pharmacy_user_id"`
}
type Consumer struct {
	UserID        uint      `json:"user_id"`
	Name          string    `json:"name"`
	FireBaseToken string    `json:"fire_base_token"`
	LineToken     string    `json:"line_token"`
	Gender        string    `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	MobilePhone   string    `json:"mobile_phone"`
	Email         string    `json:"email"`
}

type LicenseDeleted struct {
	ID uint `json:"id"`
}
type LicenseUpdated struct {
	ID        uint             `json:"id"`
	CreatedAt time.Time        `json:"created_at" `
	UpdatedAt time.Time        `json:"updated_at"`
	UserID    uint             `json:"user_id"`
	Blocked   bool             `json:"blocked"`
	Details   []*LicenseDetail `json:"details" `
}
type LicenseDetailDeleted struct {
	ID uint `json:"id"`
}
type LicenseDetail struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LicenseID uint      `json:"license_id"`
	ProductID uint      `json:"product_id"`
	StartAt   time.Time `json:"start_at"`
	EndAt     time.Time `json:"end_at"`
}
type ProductUpdated struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}
type ProductDeleted struct {
	ID uint `json:"id"`
}
