package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {

	set_timer := flag.Float64("t", 5, "Set the countdown timer")

	flag.Parse()

	if *set_timer != 0 {
		fmt.Printf("%v secs ticking... \n", *set_timer)
		timer := time.NewTimer(time.Duration(*set_timer * float64(time.Second)))
		<-timer.C
		fmt.Println("Time is up!")
		log.Fatal(play_alarm())

	}

}
