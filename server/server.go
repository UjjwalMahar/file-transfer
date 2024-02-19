package main

import (
    "fmt"
    "io"
    "net"
    "os"
)

const (
    PORT = "5000"
    TYPE = "tcp"
)
func main() {
    // Create a listener
    listener, err := net.Listen(TYPE, ":" + PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer listener.Close()
    fmt.Println("Server listening on : %v" ,PORT)

    // Accept connections
    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Error accepting connection:", err.Error())
        return
    }
    defer conn.Close()

    // Receive file
    fmt.Println("Receiving file...")
    receiveFile(conn)
    fmt.Println("File received.")
}

func receiveFile(conn net.Conn) {
    // Create a buffer to store received data temporarily
    buffer := make([]byte, 1024)

    // Create a new file to write to
    file, err := os.Create("recieved.txt")
    if err != nil {
        fmt.Println("Error creating file:", err.Error())
        return
    }
    defer file.Close()

    // Read data from connection and write to file
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            if err != io.EOF {
                fmt.Println("Error reading:", err.Error())
            }
            break
        }
        if n > 0 {
            _, err := file.Write(buffer[:n])
            if err != nil {
                fmt.Println("Error writing to file:", err.Error())
                break
            }
        }
    }
}
