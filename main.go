package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {

	timer := flag.Float64("t", 5, "Set the countdown timer")

	flag.Parse()

	switch *timer {
	case 1:
		fmt.Println("1sec ticking....")
	case 2:
		fmt.Println("2secs ticking....")
		timer := time.NewTimer(2 * time.Second)
		<-timer.C
		fmt.Println("Time is up!")
	case 3:
		fmt.Println("3secs ticking....")
		timer := time.NewTimer(3 * time.Second)
		<-timer.C
		fmt.Println("Time is up!")
		log.Fatal(play_alarm())

	case 4:
		fmt.Println("4sec ticking....")
	case 5:
		fmt.Println("5sec ticking....")

	}

}
