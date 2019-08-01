package main

import (
	"fmt"

	"github.com/heltonmarx/goami/ami"
)

type Bill struct {
	event   string
	billsec int
}

var bills map[string]Bill

func Billing(c <-chan ami.Response) {

	//chans := map[string]chan map[string]string{}
	for e := range c {

		fmt.Println(e.Get("Event"))

		uniqueId := e.Get("Uniqueid")
		linkedId := e.Get("Linkedid")

		switch e.Get("Event") {

		case "OriginateResponse":
			{
				fmt.Println("[DEBUG]: Originate Event")
				b := bills[uniqueId]
				b.event = "originate"
			}

		case "Newchannel":
			{
				fmt.Println("[DEBUG]: Newchannel Event")

				if b, ok := bills[linkedId]; ok {
					b.event = "originate"
				} else if b, ok := bills[uniqueId]; !ok {
					b.event = "manual"
				}

				fmt.Println(e.Get("AccountID"))
				fmt.Println(e.Get("LocalAddress"))
			}
		case "Hangup":
			{
				fmt.Println(e)
			}
		}
	}
}
