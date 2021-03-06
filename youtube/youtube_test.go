package youtube

import (
	"fmt"
	"strings"
	"testing"
)

var videoIDs = []string{"dQw4w9WgXcQ", "TYgOlqinH7A", "mOYZaiDZ7BM"}

var youtubeLongLinkExample = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
var youtubeShortLinkExample = "https://youtu.be/dQw4w9WgXcQ"

func ExampleCreatePlaylist() {
	fmt.Println(CreatePlaylist(videoIDs))
	// Output:
	// https://www.youtube.com/watch_videos?video_ids=dQw4w9WgXcQ,TYgOlqinH7A,mOYZaiDZ7BM
}

func TestGenerateURLFromIDs(t *testing.T) {
	expected := "https://www.youtube.com/watch_videos?video_ids=dQw4w9WgXcQ,TYgOlqinH7A,mOYZaiDZ7BM"
	got := generateURLFromIDs(videoIDs)

	if strings.Compare(expected, got) != 0 {
		t.Errorf("Playlist URL was incorrect, got: %v, expected: %v.", got, expected)
	}
}

func TestParseID(t *testing.T) {
	expected := "dQw4w9WgXcQ"
	got := parseID(youtubeLongLinkExample)

	if strings.Compare(expected, got) != 0 {
		t.Errorf("YouTube video ID was incorrect, got: %v, expected: %v.", got, expected)
	}

	expected = "dQw4w9WgXcQ"
	got = parseID(youtubeShortLinkExample)

	if strings.Compare(expected, got) != 0 {
		t.Errorf("YouTube video ID was incorrect, got: %v, expected: %v.", got, expected)
	}
}
