package generichandler

import "encoding/json"

type Nil struct{}

type handler interface {
	handle(message []byte) error
}

type h[T any] struct {
	handler func(data T) error
}

func Handler[T any](handler func(T) error) handler {
	return h[T]{handler: handler}
}

func (h h[T]) handle(data []byte) error {
	var t T

	if _, ok := any(t).(Nil); !ok {
		err := json.Unmarshal(data, &t)
		if err != nil {
			return err
		}
	}

	return h.handler(t)
}
