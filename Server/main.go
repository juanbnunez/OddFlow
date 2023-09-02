package main

import (
	"Server/server"
)

func main() {
	//musicserver
	_, err := server.NewMusicServer("localhost:8080")
	if err != nil {
		// Manejar el error
		return
	}

	// Iniciar la gestión de solicitudes del servidor
	// Por ejemplo, aceptar conexiones y manejar las solicitudes de los clientes
}
