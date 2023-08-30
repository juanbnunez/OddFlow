package playlist

import (
	"Server/song"
	"fmt"
)

type Playlist struct {
	Songs []song.Song
	// You can add more fields here if needed
}

// NewPlaylist creates a new empty playlist.
func NewPlaylist() *Playlist {
	return &Playlist{}
}

// AddSong adds a song to the playlist.
func (p *Playlist) AddSong(s song.Song) {
	p.Songs = append(p.Songs, s)
}

// RemoveSong removes a song from the playlist by its name.
func (p *Playlist) RemoveSongByName(name string) {
	var updatedSongs []song.Song
	for _, s := range p.Songs {
		if s.Name != name {
			updatedSongs = append(updatedSongs, s)
		}
	}
	p.Songs = updatedSongs
}

// Example usage
func main() {
	playlist := NewPlaylist()

	song1 := song.Song{Name: "Song 1", FilePath: "path/to/song1.mp3"}
	song2 := song.Song{Name: "Song 2", FilePath: "path/to/song2.mp3"}

	playlist.AddSong(song1)
	playlist.AddSong(song2)

	// Display playlist songs
	for _, s := range playlist.Songs {
		fmt.Printf("Song Name: %s\n", s.Name)
		fmt.Printf("File Path: %s\n", s.FilePath)
		fmt.Println("----")
	}

	// Remove a song from the playlist
	playlist.RemoveSongByName("Song 1")

	// Display updated playlist songs
	for _, s := range playlist.Songs {
		fmt.Printf("Song Name: %s\n", s.Name)
		fmt.Printf("File Path: %s\n", s.FilePath)
		fmt.Println("----")
	}
}
