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
	Pharmacy         *Pharmacy         `json:"pharmacy"`          // 藥局資料
	PharmacyEmployee *PharmacyEmployee `json:"pharmacy_employee"` // 藥局員工資料
	Consumer         *Consumer         `json:"consumer"`          // 消費者資料
}

type SjenEmployee struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

type Pharmacy struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}
type PharmacyEmployee struct {
	ID         uint     `json:"id"`
	UserID     uint     `json:"user_id"`
	Name       string   `json:"name"`
	PharmacyID uint     `json:"pharmacy_id"`
	Pharmacy   Pharmacy `json:"pharmacy"`
}
type Consumer struct {
	ID             uint   `json:"id"`
	UserID         uint   `json:"user_id"`
	Name           string `json:"name"`
	FireBaseToken  string `json:"fire_base_token"`
	IdentityNumber string `json:"identity_number"`
	MobilePhone    string `json:"mobile_phone"`
	Email          string `json:"email"`
}
