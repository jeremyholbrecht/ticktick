package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
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

func play_alarm() error {
	// Open the file
	file, err := os.Open("alarm.mp3")
	if err != nil {
		return err
	}

	defer file.Close()

	//decode the file
	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	//create a new player that will handle the sound, paused by default
	player := c.NewPlayer(decoder)
	defer player.Close()
	player.Play()

	for {
		time.Sleep(time.Second)
		if !player.IsPlaying() {
			break
		}
	}

	return nil
}
