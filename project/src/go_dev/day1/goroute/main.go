package main

import (
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		testGoroute(i)
	}

	time.Sleep(10 * time.Second)
}
