package term

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Screen2 : renders the searching Menu in terminal
func Screen2() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	YTdesign := widgets.NewParagraph()
	YTdesign.SetRect(40, 0, 400, 9)
	YTdesign.Border = false
	YTdesign.TextStyle.Fg = ui.ColorRed
	YTdesign.Text = `
	██╗   ██╗████████╗      ██████╗ ██╗      █████╗ ██╗   ██╗███████╗██████╗ 
	╚██╗ ██╔╝╚══██╔══╝      ██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝██╔══██╗
	 ╚████╔╝    ██║         ██████╔╝██║     ███████║ ╚████╔╝ █████╗  ██████╔╝
	  ╚██╔╝     ██║         ██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██╔══╝  ██╔══██╗
	   ██║      ██║         ██║     ███████╗██║  ██║   ██║   ███████╗██║  ██║
	   ╚═╝      ╚═╝         ╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝`

	title := widgets.NewParagraph()
	title.Text = "-- A Online Terminal Music App"
	title.Border = false
	title.TextStyle.Fg = ui.ColorRed
	title.SetRect(105, 9, 149, 12)

	searcherPane := widgets.NewTabPane("Audio", "Video", "Playlists", "Go Back", "Quit")
	searcherPane.SetRect(40, 10, 90, 13)
	searcherPane.Border = true

	renderTabTwo := func() {
		switch searcherPane.ActiveTabIndex {
		case 0:
			Screen3("Audios")
		case 1:
			Screen3("Videos")
		case 2:
			Screen3("Playlists")
		case 3:
			Screen2()
			break
		case 4:
			ui.Clear()
			ui.Close()
			os.Exit(0)
		}
	}

	ui.Render(YTdesign, title, searcherPane)

	menuEvents := ui.PollEvents()

	for {
		e := <-menuEvents

		switch e.ID {
		case "<Escape>":
			return
		case "<Left>":
			searcherPane.FocusLeft()
			ui.Clear()
			ui.Render(YTdesign, title, searcherPane)
			if e.Type == ui.KeyboardEvent && e.ID == "<Enter>" {
				renderTabTwo()
			}
		case "<Right>":
			searcherPane.FocusRight()
			ui.Clear()
			ui.Render(YTdesign, title, searcherPane)
			if e.Type == ui.KeyboardEvent && e.ID == "<Enter>" {
				renderTabTwo()
			}
		case "<Enter>":
			renderTabTwo()
		}

	}
}
