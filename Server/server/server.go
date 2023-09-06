package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	//"github.com/OddFlow/oto"


)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	fmt.Println("*** Server STARTED ***")
	connectionNumber := 1
	// Antes del bucle for en la función main
	playerContext, err := oto.NewContext(44100, 2, 2, 8192)
	if err != nil {
		log.Fatal(err)
	}
	defer playerContext.Close()
	player := playerContext.NewPlayer()

	for {
		fmt.Println("Wainting for connections...")
		con, err := listener.Accept()
		//EN LA LINEA 24 SE ESPERA HASTA QUE SE ENCUENTRE UNA NUEVA CONEXIÓN. SOLO ANTE UNA NUEVA SE SIGUE EN LA EJECUCIÓN

		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Connection#", connectionNumber, " started!")

		//UNA VEZ ESTABLECIDA UNA CONEXIÓN SIN ERRORES, SE EJECUTA EL MANEJO DE ENVÍO DE MENSAJES CON DICHA CONEXIÓN
		//NÓTESE QUE SE HACE CON UNA go routine PARA QUE SE EJECUTE EN PARALELO Y EL CICLO DE ESTE MÉTODO PUEDA CONTINUAR Y ESPERAR MÁS CONEXIONES
		go handleClientRequest(con, connectionNumber)
		connectionNumber += 1
	}
}

/* codigo del profe******************************
func handleClientRequest(con net.Conn, cn int) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	for {
		// Waiting for the client request
		// MIENTRAS QUE LA CONEXIÓN QUE SE HIZO Y QUE SE MANDÓ POR PARÁMETRO ESTÉ ACTIVA
		// LA go routine SE ESPERARÁ ACÁ HASTA QUE PUEDA LEER UN MENSAJE QUE VENGA DEL CLIENTE DE ESA CONEXIÓN
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			if clientRequest == ":QUIT" {
				log.Println("client requested server to close the connection so closing")
				return
			} else {
				log.Println("CLIENT#", cn, " says: ", clientRequest)

				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(5000)
				time.Sleep(time.Duration(n) * time.Millisecond)

				// Responding to the client request
				// SE LE ENVÍA RESPUESTA AL CLIENTE Y SE CONTINÚA CON EL CICLO ESPERANDO CUALQUIER OTRO MENSAJE DE CLIENTE
				// EN EL HIPOTÉTICO DE QUE NO DEBA HABER MÁS MENSAJES CON EL CLIENTE, LA FORMA DE CERRAR LA CONEXIÓN SERÁ SIEMPRE
				// SALIR DE ESTE CICLO, YA QUE TERMINARÍA LA go routine Y ANTES DE TERMINAR LA LLAMADA AL MÉTODO Y SUI RETORNO
				// LA LÍNEA 41 ESTABLECE QUE SE CIERRE LA CONEXIÓN.
				if _, err = con.Write([]byte("GOT IT!\n")); err != nil {
					log.Printf("failed to respond to client: %v\n", err)
				}
			}
		case io.EOF:
			log.Println("client closed the connection by terminating the process")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}
	}
}*/
// ... Sugerencia creada por chat*******************

func handleClientRequest(con net.Conn, cn int) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	for {
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			if clientRequest == ":QUIT" {
				log.Println("client requested server to close the connection so closing")
				return
			} else if strings.HasPrefix(clientRequest, ":PLAY ") {
				songName := strings.TrimPrefix(clientRequest, ":PLAY ")
				log.Printf("CLIENT#", cn, " requested to play: ", songName)

				// Abre el archivo de audio
				file, err := os.Open(songName)
				if err != nil {
					log.Printf("failed to open audio file: %v\n", err)
					return
				}
				defer file.Close()

				// Lee el archivo de audio y envíalo al reproductor
				_, err = io.Copy(player, file)
				if err != nil {
					log.Printf("failed to play audio: %v\n", err)
				}

				// Responde al cliente
				if _, err := con.Write([]byte("Now playing: " + songName + "\n")); err != nil {
					log.Printf("failed to respond to client: %v\n", err)
				}
			}


// ...

/*código de juan************************************
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)


import (
	"Server/communication"
	"Server/playlist"
	"bufio"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

type MusicServer struct {
	Communication communication.Server
	Playlist      *playlist.Playlist
	// Otros campos relevantes
}

// Funciones para gestionar solicitudes y acciones del servidor
func NewMusicServer(address string) (*MusicServer, error) {
	// Crear un nuevo servidor de música
	return nil, nil
}*/

// Otras funciones relacionadas con la gestión del servidor
