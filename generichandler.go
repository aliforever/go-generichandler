package generichandler

import (
	"encoding/json"
	"errors"
)

var EmptyDataErr = errors.New("empty_data")

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
	if len(data) == 0 {
		return EmptyDataErr
	}

	var t T

	if _, ok := any(t).(Nil); !ok {
		err := json.Unmarshal(data, &t)
		if err != nil {
			return err
		}
	}

	return h.handler(t)
}
