package main

import (
	"flag"
	"fmt"
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
		play_alarm()

	case 4:
		fmt.Println("4sec ticking....")
	case 5:
		fmt.Println("5sec ticking....")

	}

}

func play_alarm() error {
	f, err := os.Open("alarm.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}
