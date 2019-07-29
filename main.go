package main

import (
	"fmt"
	"os"

	"github.com/ivahaev/amigo"
)

// Default config
const (
	AsteriskHost = "127.0.0.1"
	AsteriskPort = "5038"
	AsteriskUser = "admin"
	AsteriskPass = "admin"
)

// Go AMI Asterisk
var a *amigo.Amigo

func main() {

	settings := &amigo.Settings{Host: AsteriskHost, Port: AsteriskPort, Username: AsteriskUser, Password: AsteriskPass}

	if e := os.Getenv("ASTERISK_HOST"); len(e) > 0 {
		fmt.Println(e)
		settings.Host = e
	}
	if e := os.Getenv("ASTERISK_PORT"); len(e) > 0 {
		settings.Port = e
	}
	if e := os.Getenv("ASTERISK_USERNAME"); len(e) > 0 {
		settings.Username = e
	}
	if e := os.Getenv("ASTERISK_PASSWORD"); len(e) > 0 {
		settings.Password = e
	}

	a = amigo.New(settings)

	a.On("connect", func(message string) {
		fmt.Println("Connected: ", message)
	})

	a.On("error", func(message string) {
		fmt.Println("Connection error: ", message)
	})

	a.Connect()
	c := make(chan map[string]string, 100)

	a.SetEventChannel(c)
	ProcessEvents(c)
}
