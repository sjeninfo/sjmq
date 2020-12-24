package sjmq

import "reflect"

func getChannelName(v interface{}) string {
	var channel string
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		channel = reflect.TypeOf(v).Elem().Name()
	} else {
		channel = reflect.TypeOf(v).Name()
	}
	return channel
}
