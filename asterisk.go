package main

import (
	"log"
	"sync"

	"github.com/heltonmarx/goami/ami"
)

type Asterisk struct {
	socket *ami.Socket
	uuid   string

	events chan ami.Response
	stop   chan struct{}
	wg     sync.WaitGroup
}

// NewAsterisk initializes the AMI socket with a login and capturing the events.
func NewAsterisk(host string, username string, secret string) (*Asterisk, error) {

	socket, err := ami.NewSocket(host)
	if err != nil {
		return nil, err
	}

	uuid, err := ami.GetUUID()
	if err != nil {
		return nil, err
	}

	const events = "call,user"
	err = ami.Login(socket, username, secret, events, uuid)
	if err != nil {
		return nil, err
	}

	as := &Asterisk{
		socket: socket,
		uuid:   uuid,
		events: make(chan ami.Response),
		stop:   make(chan struct{}),
	}

	as.wg.Add(1)
	go as.run()

	return as, nil
}

// Logoff closes the current session with AMI.
func (as *Asterisk) Logoff() error {
	close(as.stop)
	as.wg.Wait()

	return ami.Logoff(as.socket, as.uuid)
}

// Events returns an channel with events received from AMI.
func (as *Asterisk) Events() <-chan ami.Response {
	return as.events
}

// run - listen events Asterisk
func (as *Asterisk) run() error {

	defer as.wg.Done()
	for {
		select {
		case <-as.stop:
			return nil
		default:
			events, err := ami.Events(as.socket)
			if err != nil {
				log.Printf("AMI events failed: %v\n", err)
				return err
			}
			as.events <- events
		}
	}
}
