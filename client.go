package sjmq

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	Receiver *Receiver
	Notifier *Notifier
	DB       *gorm.DB
}

func NewClient(
	publisherHost string,
	dbType string,
	dsn string,
	mqHost string,
	group string,
) *Client {
	c := new(Client)
	db, err := initialDB(dbType, dsn)
	if err != nil {
		panic(err)
	}
	c.DB = db
	c.Notifier = NewNotifier(publisherHost, dbType, dsn)
	c.Receiver = NewReceiver(db, mqHost, group)
	outboxRepo := NewOutboxRepository(NewGormDBContext(db))
	if outboxRepo.HasMessage() {
		c.Notifier.Notify()
	}
	return c
}

func (c *Client) NewSender() (ISender, *SenderContext) {
	ctx := NewSenderContext(c.DB, c.Notifier.Notify)
	s := NewSender(ctx)
	return s, ctx
}

func initialDB(dbType, dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch dbType {
	case "mssql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite3":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		err = errors.New("unsupported database type")
	}
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&OutMessage{}, &InMessage{})
	return db, nil
}
