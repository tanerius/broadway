package actor

import "fmt"

type ActorSystem struct {
	actors map[*PID]chan Script
}

func NewActorSystem() *ActorSystem {
	system := &ActorSystem{}
	system.Init()
	return system
}

// Implementing Init from Actor
func (system *ActorSystem) Init() {
	system.actors = make(map[*PID]chan Script)
}

// Implementing Creator
func (system *ActorSystem) Create(actor Actor) *PID {
	pid := newPid()
	messageChannel := make(chan Script, 10) // buffered channel
	system.actors[pid] = messageChannel
	go func() {
		actor.Init()
		for msg := range messageChannel {
			err := actor.Receive(msg)
			if err != nil {
				// Handle error, maybe restart the actor or log the error
				fmt.Println("ERROR")
				return
			}
		}
	}()
	return pid
}

// Implementing Receive from Actor
func (system *ActorSystem) Receive(script Script) error {
	return nil
}

// Implementing Sender
func (system *ActorSystem) Send(id *PID, message Script) error {
	if messageChannel, ok := system.actors[id]; ok {
		messageChannel <- message
		return nil
	}
	return fmt.Errorf("actor with id %s not found", id)
}
