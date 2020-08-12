package sjmq

import "time"

type CustomerImportedEvent struct {
	CustomerNo      string     `json:"customer_no" mapstructure:"customer_no"`
	InputNo         string     `json:"input_no" mapstructure:"input_no"`
	Category        string     `json:"category" mapstructure:"category"`
	Name            string     `json:"name" mapstructure:"name"`
	UseIndex        string     `json:"use_index" mapstructure:"use_index"`
	Telephone1      string     `json:"telephone1" mapstructure:"telephone1"`
	Telephone2      string     `json:"telephone2" mapstructure:"telephone2"`
	Mobilphone      string     `json:"mobilphone" mapstructure:"mobilphone"`
	Fax             string     `json:"fax" mapstructure:"fax"`
	Contact         string     `json:"contact" mapstructure:"contact"`
	Uniform         string     `json:"uniform" mapstructure:"uniform"`
	BillingAddress  string     `json:"billing_address" mapstructure:"billing_address"`
	ZipCode         string     `json:"zip_code" mapstructure:"zip_code"`
	Email           string     `json:"email" mapstructure:"email"`
	PostFlag        bool       `json:"post_flag" mapstructure:"post_flag"`
	TransactionDate string     `json:"transaction_date" mapstructure:"transaction_date"`
	Area            string     `json:"area" mapstructure:"area"`
	Discount        string     `json:"discount" mapstructure:"discount"`
	Remark2         string     `json:"remark2" mapstructure:"remark2"`
	Remark          string     `json:"remark" mapstructure:"remark"`
	InvoiceTitle    string     `json:"invoice_title" mapstructure:"invoice_title"`
	InvoiceAddress  string     `json:"invoice_address" mapstructure:"invoice_address"`
	DeliveryAddress string     `json:"delivery_address" mapstructure:"delivery_address"`
	MonthlyClose    string     `json:"monthly_close" mapstructure:"monthly_close"`
	MonthlyRequest  string     `json:"monthly_request" mapstructure:"monthly_request"`
	OpeningAmount   float64    `json:"opening_amount" mapstructure:"opening_amount"`
	CreditLine      float64    `json:"credit_line" mapstructure:"credit_line"`
	DeliveryRoute   string     `json:"delivery_route" mapstructure:"delivery_route"`
	Url             string     `json:"url" mapstructure:"url"`
	License         string     `json:"license" mapstructure:"license"`
	DeletedAt       *time.Time `json:"deleted_at" mapstructure:"deleted_at"`
	CreatedAt       time.Time  `json:"created_at" mapstructure:"created_at"`
	CreatedBy       string     `json:"created_by" mapstructure:"created_by"`
	UpdatedAt       time.Time  `json:"updated_at" mapstructure:"updated_at"`
	UpdatedBy       string     `json:"updated_by" mapstructure:"updated_by"`
	Used            bool       `json:"used" mapstructure:"used"`
}
