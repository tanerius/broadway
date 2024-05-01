package actor

// The Actor interface that every actor type must implement. An actor can:
// - Send a finite number of messages
// - Create a finite number of actors
// - Designate the behaviour (message handler) to be used for the next message
// From Now on EVERYTHING is an actor.
type Actor interface {
	Receive(script any) error
	Init()
}
