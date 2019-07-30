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
