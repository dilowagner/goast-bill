package main

import (
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

func NewAsterisk(host string, username string, secret string) (*Asterisk, error) {

	socket, err := ami.NewSocket(host)
	if err != nil {
		return nil, err
	}

	uuid, err := ami.GetUUID()
	if err != nil {
		return nil, err
	}

	const events = "system,call,all,user"
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
	//go as.run()
	return as, nil
}
