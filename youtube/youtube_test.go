package youtube

import (
	"strings"
	"testing"
)

var videoIDs []string = []string{"dQw4w9WgXcQ", "TYgOlqinH7A", "mOYZaiDZ7BM"}

var youtubeLongLinkExample string = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
var youtubeShortLinkExample string = "https://youtu.be/dQw4w9WgXcQ"

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
