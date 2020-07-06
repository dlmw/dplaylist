package youtube

import (
	"strings"
)

const longLinkPrefix string = "https://www.youtube.com/watch?v="
const shortLinkPrefix string = "https://youtu.be/"
const idLength int = 11

const playlistPrefix string = "https://www.youtube.com/watch_videos?video_ids="

// CreatePlaylist takes a list of short or long YouTube URLs
// as argument and "compiles" them in a single link (playlist)
func CreatePlaylist(videoIDs []string) string {
	cleanIDs := make([]string, 0)
	for _, videoID := range videoIDs {
		cleanIDs = append(cleanIDs, parseID(videoID))
	}

	return generateURLFromIDs(cleanIDs)
}

// generateURLFromIDs takes a list of, AND ONLY OF, video IDs
// and generates the playlist link
func generateURLFromIDs(videoIDs []string) string {
	var playlistLink strings.Builder
	playlistLink.WriteString(playlistPrefix)

	for _, videoID := range videoIDs {
		playlistLink.WriteString(videoID)
		playlistLink.WriteString(",")
	}

	return strings.TrimSuffix(playlistLink.String(), ",")
}

// parseID removes any prefix (long or short YouTube link) and
// suffix (URL parameters) and returns the ID of the video ONLY (11 characters long)
func parseID(videoID string) string {
	var parsed string
	if strings.HasPrefix(videoID, longLinkPrefix) {
		parsed = strings.TrimPrefix(videoID, longLinkPrefix)
	} else if strings.HasPrefix(videoID, shortLinkPrefix) {
		parsed = strings.TrimPrefix(videoID, shortLinkPrefix)
	} else {
		parsed = videoID
	}

	// this statement makes sure there is nothing on the right of the video ID (e.g. URL parameters)
	return strings.TrimSuffix(parsed, parsed[idLength:])
}
