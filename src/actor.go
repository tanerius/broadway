package actor

import (
	"net"

	"github.com/google/uuid"
)

// The Actor interface that every actor type must implement. An actor can:
// - Send a finite number of messages (If it implements Sender)
// - Create a finite number of actors (if it implements Creator)
// - Designate the behaviour (message handler) to be used for the next message
// From Now on EVERYTHING is an actor.
type Actor interface {
	Receive(Script) error
	Init()
}

// An interface for a basic message. We call it a script in light of Broadway
type Script interface{}

// An actor that implements Sender will also be able to send messages
type Sender interface {
	Send(*PID, Actor)
}

// An actor that implements Creator will be able to create new actors
type Creator interface {
	Create(Actor) *PID
}

func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

func newPid() *PID {
	ip, err := getOutboundIP()
	if err != nil {
		return nil
	}

	uuid := uuid.NewString()

	return &PID{
		Address: ip.To4().String(),
		Id:      uuid,
	}
}

// A function to determine whether two PIDs are deeply equal
func (r *PID) Equals(other *PID) bool {
	if other == nil {
		return false
	}

	return (r.GetId() == other.GetId()) && (r.GetAddress() == r.GetAddress())
}
