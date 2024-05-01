package actor

import "fmt"

type ActorSystem struct {
	actors map[*PID]chan interface{}
}

func NewActorSystem() *ActorSystem {
	return &ActorSystem{
		actors: make(map[*PID]chan interface{}),
	}
}

func (system *ActorSystem) CreateActor(actor Actor) *PID {
	pid := newPid()
	messageChannel := make(chan interface{}, 10) // buffered channel
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

func (system *ActorSystem) Send(id *PID, message interface{}) error {
	if messageChannel, ok := system.actors[id]; ok {
		messageChannel <- message
		return nil
	}
	return fmt.Errorf("actor with id %s not found", id)
}
