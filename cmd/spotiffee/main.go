package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jostrzol/spotiffee/lib/inhibitor"
	"github.com/leberKleber/go-mpris"
)

const (
	SpotifyDbusName = "org.mpris.MediaPlayer2.spotify"
	InhibitReason   = "Music playing"
	CheckPeriod     = time.Second * 5
)

func main() {
	player, err := mpris.NewPlayer(SpotifyDbusName)
	if err != nil {
		fmt.Printf("failed to create gompris.Player: %s\n", err)
		os.Exit(1)
	}
	defer player.Close()

	inhibitor, err := inhibitor.New()
	if err != nil {
		fmt.Printf("failed to create inhibitor.Inhibitor: %s\n", err)
		os.Exit(1)
	}
	defer inhibitor.Close()

	for {
		s, err := player.PlaybackStatus()
		if err != nil {
			fmt.Printf("failed to get playback status: %s\n", err)
		}

		if s == mpris.PlaybackStatusPlaying && !inhibitor.IsInhibited() {
			fmt.Println("Inhibiting suspension")
			err = inhibitor.Inhibit(InhibitReason)

			if err != nil {
				fmt.Printf("failed to inhibit: %s\n", err)
			}

		} else if s != mpris.PlaybackStatusPlaying && inhibitor.IsInhibited() {
			fmt.Println("Uninhibiting suspension")
			err = inhibitor.Uninhibit()

			if err != nil {
				fmt.Printf("failed to uninhibit: %s\n", err)
			}
		}

		time.Sleep(CheckPeriod)
	}
}
