package main

import (
	"fmt"
	"github.com/aliforever/go-generichandler"
	"time"
)

type ping struct {
	Time time.Time `json:"time"`
}

func PingHandler(p ping) error {
	fmt.Println(p.Time)
	return nil
}

func main() {
	handlers := generichandler.NewHandlers()

	handlers.AddHandler("ping", generichandler.Handler[ping](PingHandler))

	event := []byte(fmt.Sprintf(`{"time": "%s"}`, time.Now().Format(time.RFC3339)))

	err := handlers.Handle("ping", event)
	if err != nil {
		panic(err)
	}
}
