package main

type Bill struct {
	event   string
	billsec int
}

var bills map[string]Bill

func ProcessEvents(c chan map[string]string) {

	chans := map[string]chan map[string]string{}
	for e := range c {

		uniqueID := e["UniqueID"]
		if len(uniqueID) == 0 {
			uniqueID = e["UniqueID"]
		}

		switch e["Event"] {

		case "Dial":
			{
				if ch, ok := chans[uniqueID]; ok {
					ch <- e
					continue
				}

				ch := make(chan map[string]string, 3)
				ch <- e

			}
		}
	}
}
