package main

import (
	"flag"
)

var (
	user   = flag.String("user", "admin", "AMI username")
	secret = flag.String("secret", "admin", "AMI secret")
	host   = flag.String("host", "127.0.0.1:5038", "AMI host address")
)

func main() {

	flag.Parse()

}
