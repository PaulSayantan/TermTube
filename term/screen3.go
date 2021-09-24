package term

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/PaulSayantan/TermTube/ytsearch"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Screen3 : Renders the search box in the terminal
func Screen3(filetype string) {
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

	empty := widgets.NewParagraph()
	empty.Text = " TYPE SOMETHING AND THEN PRESS ENTER !"
	empty.Border = true
	empty.SetRect(40, 13, 100, 16)

	query := ""
	st := &keypress{}
	prompt := func(st *keypress) {

		ui.Clear()
		p := widgets.NewParagraph()
		p.Title = "Search Box"
		p.Text = fmt.Sprintf("%s", st.queryRender)

		p.SetRect(40, 13, 100, 16)

		ui.Render(YTdesign, title, p)
	}

	prompt(st)

	menuEvents := ui.PollEvents()

	for {
		e := <-menuEvents

		if e.Type == ui.KeyboardEvent && e.ID == "<Enter>" {
			if query == "" {
				ui.Render(empty)
				time.Sleep(5 * time.Second)
				prompt(st)
				break
			}
			results := ytsearch.GetMusicList(query)
			QueryResults(query, results, filetype)
			query = ""
			break
		}
		if e.Type == ui.KeyboardEvent && e.ID == "<Escape>" {
			Screen2()
		}
		if e.Type == ui.KeyboardEvent && e.ID == "<Space>" {
			query += " "
			st.queryRender = query
			prompt(st)
		}
		if runtime.GOOS == "windows" {
			if e.Type == ui.KeyboardEvent && e.ID == "<C-<Backspace>>" {
				if query == "" {
					break
				}
				query = query[:len(query)-1]
				st.queryRender = query
				prompt(st)
			}
		} else {
			if e.Type == ui.KeyboardEvent && e.ID == "<Backspace>" {
				if query == "" {
					break
				}
				query = query[:len(query)-1]
				st.queryRender = query
				prompt(st)
			}
		}
		if runtime.GOOS == "windows" {
			if e.Type == ui.KeyboardEvent && e.ID != "<Escape>" && e.ID != "<Space>" && e.ID != "<C-<Backspace>>" {
				query += e.ID
				st.queryRender = query
				prompt(st)
			}
		} else {
			if e.Type == ui.KeyboardEvent && e.ID != "<Escape>" && e.ID != "<Space>" && e.ID != "<Backspace>" {
				query += e.ID
				st.queryRender = query
				prompt(st)
			}
		}

	}

}
