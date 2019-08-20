package main

import (
	"fmt"
	"time"

	"github.com/heltonmarx/goami/ami"
)

type Billing struct {
	bills map[string]Bill
}

type Bill struct {
	event      string
	id         int
	uniqueid   string
	linkedid   string
	callid     string
	calldate   time.Time
	from       string
	fromname   string
	to         string
	toname     string
	route      string
	channel    string
	dstchannel string
	status     string
	billsec    int
	callflow   []interface{}
}

var count map[string]int

const (
	Answered   = "ANSWERED"
	No_Answer  = "NO ANSWER"
	Busy       = "BUSY"
	Congestion = "CONGESTION"
	Failed     = "FAILED"
	Canceled   = "CANCELED"
	Invalid    = "INVALID"
)

func NewBilling() *Billing {
	bs := make(map[string]Bill)
	b := &Billing{
		bills: bs,
	}
	return b
}

func (bil *Billing) Listen(c <-chan ami.Response) {

	//chans := map[string]chan map[string]string{}
	for e := range c {

		//fmt.Println(e.Get("Event"))

		dt := time.Now()

		uniqueId := e.Get("Uniqueid")
		linkedId := e.Get("Linkedid")

		switch e.Get("Event") {

		case "OriginateResponse":
			{
				fmt.Println("[DEBUG]: Originate Event")
				b := bil.bills[uniqueId]
				b.event = "originate"
			}

		case "Newchannel":
			{
				fmt.Println("[DEBUG]: Newchannel Event")

				if b, ok := bil.bills[linkedId]; ok {
					b.event = "originate"
				} else if _, ok := bil.bills[uniqueId]; !ok {

					fmt.Println(e)

					//fmt.Println(b)
					count[uniqueId] = 1

					bil.bills[uniqueId] = Bill{
						event:      "manual",
						id:         count[uniqueId],
						callid:     e.Get("Uniqueid"),
						linkedid:   linkedId,
						calldate:   dt,
						from:       e.Get("CallerIDName"),
						to:         e.Get("Exten"),
						route:      "local",
						dstchannel: "",
						channel:    e.Get("Channel"),
						status:     e.Get("ChannelState"),
						toname:     "<unknown>",
						billsec:    0,
					}

					_, err := fmt.Println(bil.bills[uniqueId])
					if err != nil {
						panic(err)
					}
				}
				//fmt.Println(e.Get("AccountID"))
				//fmt.Println(e.Get("Channel"))
			}
		case "Hangup":
			{
				fmt.Println(e)
			}
		}
	}
}
