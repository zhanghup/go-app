package event

import (
	"github.com/asaskevich/EventBus"
)

var bus = EventBus.New()

func EventPublish(topic string, args ...interface{}) {
	bus.Publish(topic, args...)
}

func EventSubscribe(topic string, fn interface{}) {
	_ = bus.Subscribe(topic, fn)
}

func EventSubscribeOnce(topic string, fn interface{}) {
	_ = bus.SubscribeOnce(topic, fn)
}

func EventUnsubscribe(topic string, fn interface{}) {
	_ = bus.Unsubscribe(topic, fn)
}

func init() {
}
