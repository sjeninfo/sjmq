package sjmq

import "time"

type BankAccountImportedEvent struct {
	CreatedAt      time.Time  `json:"created_at" mapstructure:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" mapstructure:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" mapstructure:"deleted_at"`
	No             string     `json:"no" mapstructure:"no"`
	Name           string     `json:"name" mapstructure:"name"`
	UseIndex       string     `json:"use_index" mapstructure:"use_index"`
	CatId          string     `json:"cat_id" mapstructure:"cat_id"`
	CheckType      string     `json:"check_type" mapstructure:"check_type"`
	SubjectId      string     `json:"subject_id" mapstructure:"subject_id"`
	Telephone      string     `json:"telephone" mapstructure:"telephone"`
	DepositAmount  float64    `json:"deposit_amount" mapstructure:"deposit_amount"`
	MarginOfSafety float64    `json:"margin_of_safety" mapstructure:"margin_of_safety"`
	CreatedBy      string     `json:"created_by" mapstructure:"created_by"`
	UpdatedBy      string     `json:"updated_by" mapstructure:"updated_by"`
	Used           bool       `json:"used" mapstructure:"used"`
}
