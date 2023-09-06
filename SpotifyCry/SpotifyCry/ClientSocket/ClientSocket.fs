module SpotifyCry.ClientSocket.ClientSocket

// open System
// open System.IO
// open System.Net.Sockets
// open System.Text
// let serverAddress = "127.0.0.1"
// let serverPort = 8000
//
// let client = new TcpClient(serverAddress, serverPort)
// let stream = client.GetStream()
//
// let sendData (data: byte[]) =
//     stream.Write(data, 0, data.Length)
//
// let receiveData () =
//     let bufferSize = 1024
//     let buffer = Array.zeroCreate<byte> bufferSize
//     let mutable bytesRead = 0
//     let ms = new MemoryStream()
//
//     while (true) do
//         bytesRead <- stream.Read(buffer, 0, bufferSize)
//         if bytesRead = 0 then
//             // Exit the loop when there is no more data to read
//             ()
//         else
//             ms.Write(buffer, 0, bytesRead)
//
//     ms.ToArray()
//
// let sendRequest (request: string) =
//     let requestData = Encoding.ASCII.GetBytes(request)
//     sendData requestData
//
// let receiveResponse () =
//     let responseData = receiveData()
//     let responseText = Encoding.ASCII.GetString(responseData)
//     responseText
//
// let rec mainLoop () =
//     Console.Write("Text to send: ")
//     let request = Console.ReadLine()
//
//     if request = "quit" then
//         // Close resources and exit the loop when "quit" is entered
//         client.Close()
//     else
//         sendRequest (request + "\n")
//         let response = receiveResponse ()
//         Console.WriteLine("Received: {0}", response)
//         mainLoop ()

//Código sugerido por el chat****************************************}


open System
open System.IO
open System.Net.Sockets
open System.Text

let serverAddress = "127.0.0.1"
let serverPort = 8000

let client = new TcpClient(serverAddress, serverPort)
let stream = client.GetStream()

let sendData (data: byte[]) =
    stream.Write(data, 0, data.Length)

let receiveData () =
    let bufferSize = 1024
    let buffer = Array.zeroCreate<byte> bufferSize
    let mutable bytesRead = 0
    let ms = new MemoryStream()

    while true do
        bytesRead <- stream.Read(buffer, 0, bufferSize)
        if bytesRead = 0 then
            ()
        else
            ms.Write(buffer, 0, bytesRead)

    ms.ToArray()

let sendRequest (request: string) =
    let requestData = Encoding.ASCII.GetBytes(request)
    sendData requestData

let receiveResponse () =
    let responseData = receiveData()
    let responseText = Encoding.ASCII.GetString(responseData)
    responseText

let playSong (songName: string) =
    sendRequest (":PLAY " + songName)
    let response = receiveResponse ()
    Console.WriteLine(response)

let getSongList () =
    // Enviar una solicitud al servidor Go para obtener la lista de canciones
    sendRequest ":SONGLIST"
    let response = receiveResponse ()
    // Dividir la respuesta en una lista de canciones
    response.Split('\n') |> List.ofArray

let rec mainLoop () =
    Console.Write("Enter a song to play (or 'quit' to exit): ")
    let request = Console.ReadLine()

    if request = "quit" then
        client.Close()
    else
        if request = "songlist" then
            // Obtener y mostrar la lista de canciones
            let songList = getSongList ()
            songList |> List.iter (fun song -> Console.WriteLine(song))
        else
            playSong request
        mainLoop ()





