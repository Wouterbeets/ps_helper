package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := flag.Int("number", 10, "amount of numbers generated")
	max := flag.Int("max", 1000, "max of gen number")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *num; i++ {
		fmt.Printf("%d ", rand.Intn(*max))
	}
	fmt.Printf("%d\n", rand.Intn(*max))
}
