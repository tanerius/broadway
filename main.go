package main

import (
	"fmt"

	actor "github.com/tanerius/broadway/src"
)

type BasicMessage struct {
	Data string
}

type ErrorMessage struct {
	Error error
}

type EchoActor struct{}

func (a *EchoActor) Init() {
	fmt.Println("EchoActor initialized")
}

func (a *EchoActor) Receive(message interface{}) error {
	fmt.Println("Received message:", message)
	return nil
}

func main() {
	fmt.Println("Starting...")
	system := actor.NewActorSystem()
	echoActor := &EchoActor{}
	// create am actor through a system
	system.CreateActor("echo1", echoActor)

	//send a couple of messages
	for i := 1; i < 5; i++ {
		fmt.Println("in loop: ", i)
		//send a mesdsage to the actor
		system.Send("echo1", BasicMessage{Data: "Hello, Actor Model " + fmt.Sprint(i)})
	}
}
