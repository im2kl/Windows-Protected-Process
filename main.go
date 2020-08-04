package main

import (
	"time"

	"github.com/im2kl/Windows-Protected-Process/process"
)

func main() {
	go process.Protect()

	// causes CPU spike adding time
	for {
		time.Sleep(20 * time.Second)
	}
}
