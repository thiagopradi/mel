package main

import (
  "fmt"
  "net"
  "os"
  "github.com/tchandy/mel/command_parser"
)

func main() {
  fmt.Println("Starting MEL server")

  listener, err := net.Listen("tcp", "0.0.0.0:2653")

  if err != nil {
    fmt.Println("Error starting the server")
    os.Exit(1)
  }

  ss := new(command_parser.ServerStorage)

  for {
    connection, err := listener.Accept()

    if err != nil {
      fmt.Println("Error accepting the socket")
      os.Exit(2)
    }

    go EchoFunction(connection, ss)
  }
}

func EchoFunction(connection net.Conn, ss *command_parser.ServerStorage) {
  defer connection.Close()

  for  { 
    buf := make([]byte, 1024)

    n, err := connection.Read(buf)

    if err != nil {
      fmt.Println("Error reading connection %s", err.Error())
      return
    }

    return_value := ss.ParseCommand(string(buf[0:n]))

    fmt.Printf("\n Received %d with content %s \n", n, buf)

    _, err = connection.Write([]byte(return_value))

    if err != nil {
      fmt.Println("Fail to reply: %s", err.Error())
    } else {
      fmt.Printf("\n Successfully replied with %s \n", return_value)
    }
  }
}
