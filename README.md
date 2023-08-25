# OddFlow

OddFlow is a music player application designed to cater to everyone's music needs.

## Description

OddFlow is an application that enables the management of songs using a unique architecture. One module is responsible for synchronous handling of client requests, while the other module implements a client that can potentially be instantiated multiple times.

The request-handling module efficiently manages songs locally and provides responses to the requests made by clients. The implementation of this module aims to incorporate various techniques of efficient and concurrent imperative programming. Moreover, it is essential that the resulting program can be executed on real hardware.

To achieve these goals, the request-handling module is developed using the Go programming language. This choice facilitates parallel request access from diverse clients. For the sake of simplicity and practicality, a socket-based architecture employing a straightforward TCP-IP client-server model is chosen. In this setup, the application functions as a server that addresses a range of client requests. These clients are instances of the module/application responsible for song playback and general interaction with the client. The client module is implemented in F#, following best practices and strategies of functional programming.

## License

OddFlow is licensed under the GNU Affero General Public License v3.0. 

Please refer to the [LICENSE](https://github.com/juanbnunez/OddFlow/blob/main/LICENSE) file for more details about the terms and conditions of this license.

## Features

- Efficient management of songs
- Synchronous handling of client requests
- Support for parallel client access
- Local song administration
- Playback functionality
- Interaction with clients implemented using functional programming practices

## Installation

1. Clone this repository.
   ```
   git clone https://github.com/your-username/OddFlow.git
   ```
   
2. Follow the installation instructions specific to each module (request-handling in Go and client implementation in F#)
   
## Support

For any issues or concerns, please contact us at contact@juanbnunez.info

---

Happy coding!
