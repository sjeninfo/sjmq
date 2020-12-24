package sjmq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Notifier struct {
	PublisherHost string
	Notification  []byte
	fireNotifyCh  chan bool
}

func NewNotifier(publisherHost, dbType, dsn string) *Notifier {
	notification, _ := json.Marshal(map[string]interface{}{
		"db_type": dbType,
		"dsn":     dsn,
	})
	return &Notifier{
		PublisherHost: publisherHost,
		Notification:  notification,
		fireNotifyCh:  make(chan bool, 0),
	}
}

func (n *Notifier) notifyHandler() {
	enable := false
	fn := func() {
		waitTime := time.Microsecond
		for enable {
			err := n.notifyPublisher()
			if err != nil {
				fmt.Println(err.Error())
				time.Sleep(waitTime)
				if waitTime < 10*time.Second {
					waitTime *= 10
				}
				continue
			}
			enable = false
		}
	}

	for {
		select {
		case <-n.fireNotifyCh:
			if enable {
				continue
			}
			enable = true
			go fn()
		}
	}
}

func (n *Notifier) notifyPublisher() error {
	resp, err := http.Post(
		fmt.Sprintf("%s/notify", n.PublisherHost),
		"application/json",
		bytes.NewBuffer(n.Notification),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.Errorf(string(body))
	}
	return nil
}

func (n *Notifier) Notify() {
	n.fireNotifyCh <- true
}

//func (n *Notifier) connectDb(dbType, dsn string) {
//	var db *gorm.DB
//	var err error
//	switch dbType {
//	case "mssql":
//		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	case "mysql":
//		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	case "postgres":
//		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	case "sqlite3":
//		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
//	default:
//		err = fmt.Errorf("unsupported database type")
//	}
//	if err != nil {
//		panic(err)
//	}
//	db.AutoMigrate(&Message{})
//	s.db = db
//}
