# go-generichandler
A simple json event handler using generics made to save hours of my life in tens of projects

## Get
`go get `

## Example:
```go
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
```

Basically I follow a convention in my projects, each handler has an input of type []byte (json.RawMessage).

With this package we remove the need to decode json types inside handler functions to their specific type