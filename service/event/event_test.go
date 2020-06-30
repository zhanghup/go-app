package event_test

import (
	"fmt"
	"github.com/zhanghup/go-app/service/event"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	event.EventSubscribe("test", func(a, b int) {
		fmt.Println(a + b)
	})

	for {
		event.EventPublish("test", 1, 2)
		time.Sleep(time.Second)
	}
}

func TestUserLogin(t *testing.T) {

	event.UserLoginSubscribe(func(ty, id string) {
		fmt.Printf("type:%s id: %s \n", ty, id)
	})

	for {
		event.UserLogin("web", "123")
		time.Sleep(time.Second)
	}
}
