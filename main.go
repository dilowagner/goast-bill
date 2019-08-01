package main

import (
	"flag"
	"log"
)

var (
	user   = flag.String("user", "admin", "AMI username")
	secret = flag.String("secret", "admin", "AMI secret")
	host   = flag.String("host", "127.0.0.1:5038", "AMI host address")
)

func main() {

	flag.Parse()

	asterisk, err := NewAsterisk(*host, *user, *secret)
	if err != nil {
		log.Fatal(err)
	}
	defer asterisk.Logoff()

	log.Printf("connected with asterisk\n")

	events := asterisk.Events()
	Billing(events)
	//result := <-events
	//log.Printf("EVENT: %v\n", result["Event"])
}
