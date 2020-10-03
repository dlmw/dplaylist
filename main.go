package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/dlmw/dplaylist/youtube"
)

var filePath string

func init() {
	const (
		filePathShort   = "f"
		filePathLong    = "file"
		filePathDefault = "youtube.txt"
		filePathUsage   = "The path of the file containing YouTube IDs. " + filePathShort + " and " + filePathLong + " are equivalent."
	)

	flag.StringVar(&filePath, filePathLong, filePathDefault, filePathUsage)
	flag.StringVar(&filePath, filePathShort, filePathDefault, filePathUsage)
}

func main() {
	flag.Parse()
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
