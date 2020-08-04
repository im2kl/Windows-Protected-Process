package main

import (
	"fmt"
	"time"

	"github.com/im2kl/Windows-Protected-Process/process"
)

func main() {
	go process.Protect()

	// causes CPU spike adding time
	for {
		fmt.Printf("Try killing with non admin tool :)")
		time.Sleep(20 * time.Second)
	}
}
