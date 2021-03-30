package sjmq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
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

	n := &Notifier{
		PublisherHost: publisherHost,
		Notification:  notification,
		fireNotifyCh:  make(chan bool, 0),
	}
	go n.notifyHandler()
	return n
}

func (n *Notifier) notifyHandler() {
	enable := false
	fn := func() {
		waitTime := time.Microsecond
		for enable {
			err := n.notifyPublisher()
			if err != nil {
				log.Println(err)
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
		fmt.Sprintf("http://%s/notify", n.PublisherHost),
		"application/json",
		bytes.NewBuffer(n.Notification),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.Errorf(string(body))
	}
	return nil
}

func (n *Notifier) Notify() {
	n.fireNotifyCh <- true
}
