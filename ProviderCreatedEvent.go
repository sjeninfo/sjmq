package sjmq

import (
	"time"
)

type ProviderCreatedEvent struct {
	ID              uint       `json:"id" mapstructure:"id"`
	CreatedAt       time.Time  `json:"created_at" mapstructure:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" mapstructure:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at" mapstructure:"deleted_at"`
	ProviderNumber  string     `json:"provider_no" mapstructure:"provider_no"`
	CategoryID      uint       `json:"category_id" mapstructure:"category_id"`
	Name            string     `json:"name" mapstructure:"name"`
	Remark1         string     `json:"remark1" mapstructure:"remark1"`
	UseIndex        string     `json:"use_index" mapstructure:"use_index"`
	InvoiceTitle    string     `json:"invoice_title" mapstructure:"invoice_title"`
	Uniform         string     `json:"uniform" mapstructure:"uniform"`
	Zip_code        string     `json:"zip_code" mapstructure:"zip_code"`
	InvoiceAddress  string     `json:"invoice_address" mapstructure:"invoice_address"`
	Address         string     `json:"address" mapstructure:"address"`
	Telephone1      string     `json:"telephone1" mapstructure:"telephone1"`
	Telephone2      string     `json:"telephone2" mapstructure:"telephone2"`
	Mobilphone      string     `json:"mobilphone" mapstructure:"mobilphone"`
	Fax             string     `json:"fax" mapstructure:"fax"`
	Contact         string     `json:"contact" mapstructure:"contact"`
	Email           string     `json:"email" mapstructure:"email"`
	Website         string     `json:"website" mapstructure:"website"`
	Remark2         string     `json:"remark2" mapstructure:"remark2"`
	Prepayment      float64    `json:"prepayment" mapstructure:"prepayment"`
	PostFlag        bool       `json:"post_flag" mapstructure:"post_flag"`
	WireAccount     string     `json:"wire_account" mapstructure:"wire_account"`
	BankAccount     string     `json:"bank_account" mapstructure:"bank_account"`
	License         string     `json:"license" mapstructure:"license"`
	TransactionDate string     `json:"transaction_date" mapstructure:"transaction_date"`
	Remark          string     `json:"remark" mapstructure:"remark"`
	OpeningAmount   float64    `json:"opening_amount" mapstructure:"opening_amount"`
	MonthlyClose    string     `json:"monthly_close" mapstructure:"monthly_close"`
	MonthlyRequest  string     `json:"monthly_request" mapstructure:"monthly_request"`
	UpdatedBy       string     `json:"updated_by" mapstructure:"updated_by"`
	CreatedBy       string     `json:"created_by" mapstructure:"created_by"`
	Used            bool       `json:"used" mapstructure:"used"`
}
