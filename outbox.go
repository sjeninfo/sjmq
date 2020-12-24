package sjmq

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                        // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"` // 建檔日期
	Channel   string    `json:"channel" mapstructure:"channel" gorm:"Column:channel"`          // 頻道(事件)名稱
	Event     string    `json:"event" mapstructure:"event" gorm:"Column:event"`                // 事件內容
}

func (Message) TableName() string {
	return "outbox"
}

type OutboxRepository struct {
	DBContext *GormDBContext
}

func (o *OutboxRepository) SetDBContext(ctx *GormDBContext) {
	o.DBContext = ctx
}

func NewOutboxRepository(ctx *GormDBContext) *OutboxRepository {
	o := new(OutboxRepository)
	o.DBContext = ctx
	return o
}

func (o *OutboxRepository) DB() *gorm.DB {
	return o.DBContext.DB()
}

func (o *OutboxRepository) AddEvent(event interface{}) error {
	channel := getChannelName(event)
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return o.DB().Save(&Message{Channel: channel, Event: string(data)}).Error
}

func (o *OutboxRepository) HasMessage() bool {
	var count int64
	o.DB().Model(&Message{}).Count(&count)
	return count > 0
}
