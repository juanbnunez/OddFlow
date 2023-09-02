package song

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Song struct {
	ID       int
	Name     string
	FilePath string
	// Otros campos relevantes
}

// LoadSongsFromDirectory loads songs from a directory and returns a slice of songs.
func LoadSongsFromDirectory(directoryPath string) ([]Song, error) {
	var songs []Song

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ".mp3") {
			songPath := filepath.Join(directoryPath, file.Name())
			song := Song{
				Name:     strings.TrimSuffix(file.Name(), ".mp3"),
				FilePath: songPath,
			}
			songs = append(songs, song)
		}
	}

	return songs, nil
}

// Example usage
func main() {
	songs, err := LoadSongsFromDirectory("/path/to/songs/directory")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, song := range songs {
		fmt.Printf("Song Name: %s\n", song.Name)
		fmt.Printf("File Path: %s\n", song.FilePath)
		fmt.Println("----")
	}
}
