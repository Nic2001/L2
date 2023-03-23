package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {

	response, err := ntp.Query("ntp1.stratum2.ru")
	if err != nil {
		log.Fatalln(err)
	}

	time := time.Now().Add(response.ClockOffset)
	hour, min, sec := time.Clock()

	fmt.Printf("Current time is : %d:%d:%d\n", hour, min, sec)
}
