package term

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Play(filetype string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	rcvmsg := widgets.NewParagraph()
	rcvmsg.Text = fmt.Sprintf("Retrieving your %s file from Youtube.....")

}
