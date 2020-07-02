package ytvideo

import (
	"log"
	"os"

	vlc "github.com/adrg/libvlc-go"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

//PlayVideo : 	Function for video playback
func PlayVideo(uri string, name string) {
	if err := vlc.Init(); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	media, err := player.LoadMediaFromURL(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer media.Release()

	err = player.Play()
	if err != nil {
		log.Fatal(err)
	}

	manager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	NameBox := widgets.NewParagraph()
	NameBox.Text = "Audio :: " + name
	NameBox.Border = false
	NameBox.TextStyle.Fg = ui.ColorRed
	NameBox.SetRect(10, 9, 149, 12)

	ctrlList := widgets.NewList()
	ctrlList.Title = "CONTROLS"
	ctrlList.TitleStyle.Modifier = ui.ModifierBold
	ctrlList.TitleStyle.Fg = ui.ColorCyan
	ctrlList.Rows = []string{
		"[Space]	::	Toogle play/pause the music",
		"  [S]		::	Stop playing the current song and exit",
		"  [→]		::	seek forward 10s",
		"  [←]		::	seek backward 10s",
		"  [Q]		::	quit the player",
	}
	ctrlList.TextStyle = ui.NewStyle(ui.ColorRed)
	ctrlList.WrapText = true
	ctrlList.SetRect(20, 17, 110, 25)
	ctrlList.Border = true

	ui.Render(NameBox, ctrlList)

	uiEvents := ui.PollEvents()
	for player.WillPlay() {
		e := <-uiEvents
		switch e.ID {
		case "p", "<Space>":
			if player.IsPlaying() {
				player.SetPause(true)
			} else {
				player.SetPause(false)
			}
			break
		case "s", "<Escape>":
			player.Stop()
			ui.Clear()
			ui.Close()
			os.Exit(0)
			break
		case "<Left>", "a":
			sec, err := player.MediaTime()
			if err != nil {
				log.Fatal(err)
			}
			player.SetMediaTime(sec - 10000)
			player.Play()
			break
		case "<Right>", "d":
			sec, err := player.MediaTime()
			if err != nil {
				log.Fatal(err)
			}
			player.SetMediaTime(sec + 10000)
		case "q", "<C-c>":
			player.Release()
			ui.Clear()
			ui.Close()
			os.Exit(0)
		}

	}

	quit := make(chan struct{})
	eventCallback := func(event vlc.Event, userData interface{}) {
		close(quit)
	}

	eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Detach(eventID)

	<-quit
}
