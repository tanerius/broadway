package actor

import "fmt"

type ActorSystem struct {
	actors map[string]chan interface{}
}

func NewActorSystem() *ActorSystem {
	return &ActorSystem{
		actors: make(map[string]chan interface{}),
	}
}

func (system *ActorSystem) CreateActor(id string, actor Actor) {
	messageChannel := make(chan interface{}, 10) // buffered channel
	system.actors[id] = messageChannel
	go func() {
		actor.Init()
		for msg := range messageChannel {
			err := actor.Receive(msg)
			if err != nil {
				// Handle error, maybe restart the actor or log the error
				fmt.Println(err.Error())
				return
			}
		}
	}()
}

func (system *ActorSystem) Send(id string, message interface{}) error {
	if messageChannel, ok := system.actors[id]; ok {
		messageChannel <- message
		return nil
	}
	return fmt.Errorf("actor with id %s not found", id)
}
