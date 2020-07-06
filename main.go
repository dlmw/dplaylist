package main

import (
	"bufio"
	"log"
	"os"

	"github.com/dlmw/dplaylist/youtube"
)

func main() {
	filePath := os.Args[1]
	lines, err := readLines(filePath)
	if err != nil {
		panic(err)
	}

	playlistURL := youtube.CreatePlaylist(lines)
	log.Println(playlistURL)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
