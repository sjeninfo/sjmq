package sjmq

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type InMessage struct {
	ID        uint      `json:"id" mapstructure:"id" gorm:"primaryKey"`                        // 流水號
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"` // 建檔日期
	MessageID string    `json:"message_id" mapstructure:"message_id" gorm:"Column:message_id"` // Message ID
	Channel   string    `json:"channel" mapstructure:"channel" gorm:"Column:channel"`          // 頻道(事件)名稱
	Event     string    `json:"event" mapstructure:"event" gorm:"Column:event"`                // 事件內容
	Times     uint      `gorm:"Column:times"`                                                  // 處理次數
}

func (InMessage) TableName() string {
	return "inbox"
}

type InboxRepository struct {
	DBContext *GormDBContext
}

func NewInboxRepository(ctx *GormDBContext) *InboxRepository {
	i := new(InboxRepository)
	i.DBContext = ctx
	return i
}

func (i *InboxRepository) DB() *gorm.DB {
	return i.DBContext.DB()
}

func (i *InboxRepository) InsertMessage(messageId, channel, body string) error {
	return i.DB().Save(&InMessage{
		MessageID: messageId,
		Channel:   channel,
		Event:     body,
	}).Error
}

func (i *InboxRepository) GetTopMessage() (*InMessage, error) {
	var message InMessage
	err := i.DB().Clauses(clause.Locking{Strength: "UPDATE"}).First(&message).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, nil
	}
	return &message, nil
}

func (i *InboxRepository) DeleteMessage(message *InMessage) error {
	return i.DB().Delete(message).Error
}

func (i *InboxRepository) AddTimes(message *InMessage) error {
	message.Times += 1
	return i.DB().Save(message).Error
}

// mark times.
// transfer to hospital
