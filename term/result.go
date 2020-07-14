package term

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// QueryResults :	recieves query string and sends id to player
func QueryResults(query string, results map[string]string, filetype string) {
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
	title.Text = "-- A Terminal YouTube App"
	title.Border = false
	title.TextStyle.Fg = ui.ColorRed
	title.SetRect(105, 9, 149, 12)

	resultlist := widgets.NewList()
	resultlist.Title = "Results for: " + query
	resultlist.TitleStyle.Modifier = ui.ModifierBold
	resultlist.TitleStyle.Fg = ui.ColorCyan
	var rlist []string
	for titles := range results {
		if len(titles) > 11 {
			rlist = append(rlist, titles)
		}

	}
	resultlist.Rows = rlist
	resultlist.TextStyle = ui.NewStyle(ui.ColorRed)
	resultlist.WrapText = true
	resultlist.SetRect(20, 15, 110, 33)
	resultlist.Border = false

	st := &id{}
	sendid := func(st *id) string {
		tid := results[rlist[st.index]]
		return tid
	}

	ui.Render(YTdesign, title, resultlist)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Clear()
			ui.Close()
			Screen2()
		case "j", "<Down>":
			resultlist.ScrollDown()
		case "k", "<Up>":
			resultlist.ScrollUp()
		case "<C-d>":
			resultlist.ScrollHalfPageDown()
		case "<C-u>":
			resultlist.ScrollHalfPageUp()
		case "<C-f>":
			resultlist.ScrollPageDown()
		case "<C-b>":
			resultlist.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				resultlist.ScrollTop()
			}
		case "<Home>":
			resultlist.ScrollTop()
		case "G", "<End>":
			resultlist.ScrollBottom()
		case "<Enter>":
			i := resultlist.SelectedRow
			st.index = i
			name := rlist[st.index]
			Play(sendid(st), name, filetype)
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(YTdesign, title, resultlist)
	}

}
