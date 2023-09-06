package playlist

import (
	"Server/song"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

// ValidateSong Find the path of the file with the song, validate that it is of type mp3 and call the AddSong function
func (p *Playlist) ValidateSong(projectDir string, musicFolder string) {
	// Create a new playlist instance
	//list := NewPlaylist()

	err := os.Chdir(projectDir)
	if err != nil {
		fmt.Printf("Error changing the working directory to %s: %v\n", projectDir, err)
		return
	}

	// Get the list of files in the "Music" folder
	files, err := os.ReadDir(musicFolder)
	if err != nil {
		fmt.Printf("Error reading the directory %s: %v\n", musicFolder, err)
		return
	}

	// Iterate through the files
	for _, file := range files {
		// Check if the file is a music file (e.g., with a .mp3 extension)
		if strings.HasSuffix(file.Name(), ".mp3") {
			// Create an instance of the song.Song structure
			songName := strings.TrimSuffix(file.Name(), ".mp3")
			songFilePath := filepath.Join(musicFolder, file.Name())
			songInstance := song.Song{Name: songName, FilePath: songFilePath}

			// Add the song to the playlist
			p.AddSong(songInstance)
		}

	}

}

// BrowsePlayList Browse and print the songs
func (p *Playlist) BrowsePlayList() {
	// Display playlist songs
	for _, s := range p.Songs {
		fmt.Printf("Song Name: %s\n", s.Name)
		fmt.Printf("File Path: %s\n", s.FilePath)
		fmt.Println("----")
	}
}

func (p *Playlist) GetSongs() []song.Song {
	return p.Songs
}

// RemoveSongByName removes a song from the playlist by its name.
func (p *Playlist) RemoveSongByName(name string) {
	var updatedSongs []song.Song
	for _, s := range p.Songs {
		if s.Name != name {
			updatedSongs = append(updatedSongs, s)
		}
	}
	p.Songs = updatedSongs
}

// SearchSong Search for songs by name
func (p *Playlist) SearchSong(name string) (string, error) {
	for _, s := range p.Songs {
		if s.Name == name {
			return s.FilePath, nil
		}
	}
	return "", errors.New("Can't find the song.")
}
