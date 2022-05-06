package mq

import (
	"context"
	carpb "happy-car/car/api/gen/v1"
)

// Publisher defines the publishing interface.
type Publisher interface {
	Publish(context.Context, *carpb.CarEntity) error
}

// Subscriber defines a car update subscriber.
type Subscriber interface {
	Subscribe(context.Context) (ch chan *carpb.CarEntity, cleanUp func(), err error)
}
