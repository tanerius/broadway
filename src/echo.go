package actor

import (
	"log"
	"os"
)

type BasicMessage struct {
	Data string
}

type ErrorMessage struct {
	Error error
}

type EchoActor struct {
	logger log.Logger
}

func (a *EchoActor) Init() {
	a.logger = *log.New(os.Stdout, "echo: ", 0)
	a.logger.Println("EchoActor initialized")
}

func (a *EchoActor) Receive(message interface{}) error {
	a.logger.Println("Received message:", message)
	return nil
}
