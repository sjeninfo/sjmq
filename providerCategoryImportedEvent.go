package sjmq

import "time"

type ProviderCategoryImportedEvent struct {
	CreatedAt   time.Time  `json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" mapstructure:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" mapstructure:"deleted_at"`
	No          string     `json:"no" mapstructure:"no"`
	Name        string     `json:"name" mapstructure:"name"`
	Description string     `json:"description" mapstructure:"description"`
	QuickCode   string     `json:"quick_code" mapstructure:"quick_code"`
}
