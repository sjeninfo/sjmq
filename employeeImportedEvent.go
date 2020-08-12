package sjmq

import "time"

type EmployeeImportedEvent struct {
	EmployeeNo   string     `json:"employee_no" mapstructure:"employee_no"`
	Name         string     `json:"name" mapstructure:"name"`
	UseIndex     string     `json:"use_index" mapstructure:"use_index"`
	Department   string     `json:"department" mapstructure:"department"`
	InputNo      string     `json:"input_no" mapstructure:"input_no"`
	Position     string     `json:"position" mapstructure:"position"`
	IdentityCard string     `json:"identity_card" mapstructure:"identity_card"`
	DeletedAt    *time.Time `json:"deleted_at" mapstructure:"deleted_at"`
	CreatedAt    time.Time  `json:"created_at" mapstructure:"created_at"`
	CreatedBy    string     `json:"created_by" mapstructure:"created_by"`
	UpdatedAt    time.Time  `json:"updated_at" mapstructure:"updated_at"`
	UpdatedBy    string     `json:"updated_by" mapstructure:"updated_by"`
	Used         bool       `json:"used" mapstructure:"used"`
}
