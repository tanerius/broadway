package main

import (
	"fmt"

	actor "github.com/tanerius/broadway/src"
)

func main() {
	fmt.Println("Starting...")
	system := actor.NewActorSystem()
	echoActor := &actor.EchoActor{}

	// create am actor through a system
	pid := system.Create(echoActor)
	//system.CreateActor("echo2", echoActor)

	//send a couple of messages
	for i := 1; i < 5; i++ {
		//send a mesdsage to the actor
		system.Send(pid, actor.BasicMessage{Data: "Hello, Actor echo1 " + fmt.Sprint(i)})
		//system.Send("echo2", actor.BasicMessage{Data: "Hello, Actor echo2 " + fmt.Sprint(i)})
	}
}
