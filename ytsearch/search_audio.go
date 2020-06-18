package ytsearch

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Returns user-input query
func GetQuery() string {

	fmt.Println("Search: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		queryErr := errors.New("something is wrong with your query input")
		log.Fatal(queryErr)
	}
	query = strings.TrimSpace(query)
	return query

}

// Returns a dict of title and id based on user-query
func GetMusicList(query string) map[string]string {

	qarg := `ytsearch20: ` + query

	audioList, err := exec.Command("youtube-dl", "-e", "--get-id", "--flat-playlist", qarg).CombinedOutput()
	if err != nil {
		log.Fatal(err)

	}
	list := strings.Split(string(audioList), "\n")

	titleid := make(map[string]string)

	for i := 0; i < len(list)-1; i += 2 {
		titleid[list[i]] = list[i+1]
	}

	return titleid

}
