package generichandler

import (
	"errors"
	"sync"
)

var HandlerNotFound = errors.New("handler_not_found")

type Handlers struct {
	locker sync.Mutex

	data map[string]handler
}

func NewHandlers() *Handlers {
	return &Handlers{data: map[string]handler{}}
}

func (h *Handlers) AddHandler(eventType string, eventHandler handler) *Handlers {
	h.locker.Lock()
	defer h.locker.Unlock()

	h.data[eventType] = eventHandler

	return h
}

func (h *Handlers) Handle(eventType string, data []byte) error {
	handler, exists := h.getHandler(eventType)
	if !exists {
		return HandlerNotFound
	}

	return handler.handle(data)
}

func (h *Handlers) getHandler(eventType string) (handler, bool) {
	h.locker.Lock()
	defer h.locker.Unlock()

	if h, ok := h.data[eventType]; !ok {
		return nil, false
	} else {
		return h, true
	}
}
