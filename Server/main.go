package main

import (
	"Server/playlist"
)

func main() {

	// Set the project directory path
	projectDir := "D:\\TEC\\ProyectoLenguajes\\OddFlow"

	// Get the absolute path of the "Music" folder relative to the executable's location
	musicFolder := "Music"

	// Crear una instancia de la playlist
	myPlaylist := playlist.NewPlaylist()

	// Agregar canciones desde el directorio de música
	myPlaylist.ValidateSong(projectDir, musicFolder)

	// Imprimir la lista de canciones
	myPlaylist.BrowsePlayList()

}
