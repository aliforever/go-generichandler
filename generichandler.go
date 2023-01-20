package generichandler

import "encoding/json"

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

	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}

	return h.handler(t)
}
