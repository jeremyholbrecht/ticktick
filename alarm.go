package main

import (
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

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
