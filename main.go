package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/zmb3/spotify/v2"
)

func main() {
	var seconds int
	flag.IntVar(&seconds, "seconds", 260, "The length of the liked songs to be shown, in seconds")
	flag.Parse()

	client := NewClient()
	tracks, err := getLikedSongs(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to get liked songs by length: %v", err)
	}
	filtered := filterByLength(tracks, seconds)

	log.Printf("Liked songs of length %v:\n", time.Duration(seconds)*time.Second)
	printTracks(filtered)
}

func getLikedSongs(ctx context.Context, client *spotify.Client) ([]spotify.SavedTrack, error) {
	savedTrackPage, err := client.CurrentUsersTracks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user's tracks: %w", err)
	}
	log.Printf("Total number of liked songs: %d", savedTrackPage.Total)

	log.Printf("Getting all liked songs...")
	tracks := make([]spotify.SavedTrack, 0, savedTrackPage.Total)
	for page := 1; ; page++ {
		tracks = append(tracks, savedTrackPage.Tracks...)
		err := client.NextPage(ctx, savedTrackPage)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next page of tracks: %w", err)
		}
	}
	return tracks, nil
}

// filterByLength returns a slice of tracks that are of the specified length.
// The precision is rounded to a second.
func filterByLength(tracks []spotify.SavedTrack, seconds int) []spotify.SavedTrack {
	var filtered []spotify.SavedTrack
	for _, track := range tracks {
		// Spotify UI seems to round down the duration when converting milliseconds to seconds.
		if track.Duration/1000 == seconds {
			filtered = append(filtered, track)
		}
	}
	return filtered
}

func printTracks(tracks []spotify.SavedTrack) {
	for _, track := range tracks {
		fmt.Printf("%s - %s: %s\n", track.Artists[0].Name, track.Name,
			fmt.Sprintf("https://open.spotify.com/track/%s", track.ID))
	}
}
