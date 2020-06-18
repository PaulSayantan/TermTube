package term

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type keypress struct {
	queryRender string
}

// type id struct {
// 	index int
// }

// Renders the screen in the terminal
func Initscreen1() {
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

	permitPane := widgets.NewTabPane("Search", "Quit")
	permitPane.SetRect(40, 10, 60, 13)
	permitPane.Border = true

	renderTabOne := func() {
		switch permitPane.ActiveTabIndex {
		case 0:
			Initscreen2()

		case 1:
			ui.Clear()
			ui.Close()
			os.Exit(0)
		}
	}

	ui.Render(YTdesign, title, permitPane)

	menuEvents := ui.PollEvents()
	for {
		e := <-menuEvents

		switch e.ID {
		case "<Escape>":
			ui.Clear()
			ui.Close()
			os.Exit(0)
		case "<Left>":
			permitPane.FocusLeft()
			ui.Clear()
			ui.Render(YTdesign, title, permitPane)
			if e.Type == ui.KeyboardEvent && e.ID == "<Enter>" {
				renderTabOne()
			}
		case "<Right>":
			permitPane.FocusRight()
			ui.Clear()
			ui.Render(YTdesign, title, permitPane)
			if e.Type == ui.KeyboardEvent && e.ID == "<Enter>" {
				renderTabOne()
			}
		case "<Enter>":
			renderTabOne()
		}

	}

	// for {
	// 	e := <-menuEvents

	// }

	// musicToolbar := widgets.NewTabPane("Play", "Pause", "Stop", "Seek Backward", "Jump Forward")
	// musicToolbar.SetRect(35, 13, 90, 16)
	// musicToolbar.Border = true

	// renderTab := func() {
	// 	switch musicToolbar.ActiveTabIndex {
	// 	case 0:
	// 		// write the playing music function here
	// 		ui.Render(play)
	// 	case 1:
	// 		// write the pause music function here
	// 		ui.Render(pause)
	// 	case 2:
	// 		// write the stop music function here
	// 		ui.Render(stop)
	// 	case 3:
	// 		// write the backward music function here
	// 		ui.Render(back)
	// 	case 4:
	// 		// write the jumping music function here
	// 		ui.Render(jump)
	// 	}
	// }

	// ui.Render(YTdesign, title, musicToolbar)

	// uiEvents := ui.PollEvents()

	// for {
	// 	e := <-uiEvents
	// 	switch e.ID {
	// 	case "q", "<C-c>":
	// 		return
	// 	case "left":
	// 		musicToolbar.FocusLeft()
	// 		ui.Clear()
	// 		ui.Render(YTdesign, musicToolbar)
	// 		renderTab()
	// 	case "right":
	// 		musicToolbar.FocusRight()
	// 		ui.Clear()
	// 		ui.Render(YTdesign, musicToolbar)
	// 		renderTab()
	// 	}
	// }
}
