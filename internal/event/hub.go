package event

import (
	"sync"

	"github.com/gofrs/uuid"
)

type (
	Subscription uuid.UUID
	Topic[T any] map[Subscription]chan<- T
)

type Hub[T any] struct {
	mu     sync.RWMutex
	subs   map[string]Topic[T]
	closed bool
}

func (hub *Hub[_]) Close() error {
    hub.mu.Lock()
    defer hub.mu.Unlock()

    if hub.closed {
        return nil
    }

    hub.closed = true
    for _, subs := range hub.subs {
        for _, sub := range subs {
            close(sub)
        }
    }

    return nil
}

func (hub *Hub[T]) Publish(topic string, msg T) {
    hub.mu.RLock()
    defer hub.mu.RUnlock()

    if hub.closed {
        return
    }

    for _, ch := range hub.subs[topic] {
        go func(ch chan<- T) {
            ch <- msg
        }(ch)
    }
}

func (hub *Hub[T]) Subscribe(topic string, ch chan T) Subscription {
    hub.mu.Lock()
    defer hub.mu.Unlock()

    token, _ := uuid.NewV4()
    sub := Subscription(token)
    if hub.subs[topic] == nil {
        hub.subs[topic] = make(Topic[T])
    }

    hub.subs[topic][sub] = ch
    return sub
}

func (hub *Hub[T]) Unsubscribe(topic string, sub Subscription) {
    hub.mu.Lock()
    defer hub.mu.Unlock()

    close(hub.subs[topic][sub])
    delete(hub.subs[topic], sub)
}

func NewHub[T any]() *Hub[T] {
	return &Hub[T]{
		subs: make(map[string]Topic[T]),
	}
}
