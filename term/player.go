package term

import (
	"log"
	"os/exec"
	"strings"

	"github.com/belikesayantan/youtube-tui/ytaudio"
	"github.com/belikesayantan/youtube-tui/ytvideo"
)

func Play(id, name, filetype string) {
	if filetype == "Audios" {
		uri, err := exec.Command("youtube-dl", "-g", id).CombinedOutput()
		if err != nil {
			log.Fatal(err)

		}
		audio := strings.Split(string(uri), "\n")[1]
		ytaudio.PlayAudio(audio, name)
	} else if filetype == "Videos" {
		uri, err := exec.Command("youtube-dl", "-g", "-f", "mp4", id).CombinedOutput()
		if err != nil {
			log.Fatal(err)

		}
		video := strings.Split(string(uri), "\n")[0]
		ytvideo.PlayVideo(video, name)
	}

}
