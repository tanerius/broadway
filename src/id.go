package actor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Bitfield int64

// int mapping
const (
	Bit_1  Bitfield = 1
	Bit_2  Bitfield = 2
	Bit_4  Bitfield = 4
	Bit_8  Bitfield = 8
	Bit_16 Bitfield = 16
	Bit_32 Bitfield = 32
	Bit_64 Bitfield = 64
)

type ipAddr struct {
	Query string
}

// The PID structure contains the ID information for an actor
type PID struct {
	address   string
	identity  string
	numericId int64
	typeflags Bitfield
	mu        sync.Mutex
	ipRetries int
}

func NewPID(identity string, numericId int64) *PID {
	p := &PID{
		address:   "",
		identity:  identity,
		numericId: numericId,
		typeflags: 0,
		ipRetries: 0,
	}
	go p.discoverRemoteAddress()
	return p
}

func (r *PID) RemoteAddress() string {
	var noRemote bool = false
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
		if noRemote && r.ipRetries <= 3 {
			go r.discoverRemoteAddress()
		}
	}()

	if r.address == "" {
		r.ipRetries++
		noRemote = true
		return ""
	}

	return fmt.Sprintf("/%s%s", r.address, r.Address())
}

func (r *PID) Address() string {
	return fmt.Sprintf("/%s/%d", r.identity, r.numericId)
}

func (r *PID) discoverRemoteAddress() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ipRetries++
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		r.address = ""
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		r.address = ""
	}

	var ip ipAddr
	json.Unmarshal(body, &ip)
	r.address = ip.Query
}

func (r *PID) TestFlag(flag Bitfield) bool {
	return r.typeflags&flag == flag
}

func (r *PID) ResetFlags() {
	r.typeflags = 0
}
