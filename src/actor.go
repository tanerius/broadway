package actor

import (
	"net"

	"github.com/google/uuid"
)

// The Actor interface that every actor type must implement. An actor can:
// - Send a finite number of messages
// - Create a finite number of actors
// - Designate the behaviour (message handler) to be used for the next message
// From Now on EVERYTHING is an actor.
type Actor interface {
	Receive(script any) error
	Init()
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

func (r *PID) Equals(other *PID) bool {
	if other == nil {
		return false
	}

	return (r.GetId() == other.GetId()) && (r.GetAddress() == r.GetAddress())
}
