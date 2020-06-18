package ytaudio

import (
	"log"
	"os/exec"
	"strings"
)

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"strings"
// )

func PlayAudio(id string) {
	uri, err := exec.Command("youtube-dl", "-g", id).CombinedOutput()
	if err != nil {
		log.Fatal(err)

	}
	audio := strings.Split(string(uri), "\n")[1]
	_, cmderr := exec.Command("ffplay", "-nostats", "-hide_banner", "-i", audio).CombinedOutput()
	if cmderr != nil {
		log.Fatal(cmderr)
	}
}
