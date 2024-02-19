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
    HOST = "localhost"
)

func main() {
    // Connect to server
    conn, err := net.Dial(TYPE, HOST+":"+PORT)
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        os.Exit(1)
    }
    defer conn.Close()

    // Send file
    fmt.Println("Sending file...")
    sendFile(conn, "test.txt")
    fmt.Println("File sent.")
}

func sendFile(conn net.Conn, filename string) {
    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err.Error())
        return
    }
    defer file.Close()

    // Create a buffer to store read data temporarily
    buffer := make([]byte, 1024)

    // Read from file and send over connection
    for {
        n, err := file.Read(buffer)
        if err != nil {
            if err != io.EOF {
                fmt.Println("Error reading file:", err.Error())
            }
            break
        }
        if n > 0 {
            _, err := conn.Write(buffer[:n])
            if err != nil {
                fmt.Println("Error sending data:", err.Error())
                break
            }
        }
    }
}
