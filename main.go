package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go -client|-server")
        return
    }

    arg := os.Args[1]

    switch arg {
    case "-client":
        fmt.Println("Starting client...")
        cmd := exec.Command("go", "run", "tcp-client.go")
        cmd.Stdout = os.Stdout
        cmd.Stdin = os.Stdin
        cmd.Stderr = os.Stderr
        err := cmd.Run()
        if err != nil {
            fmt.Printf("Client exited with error: %v\n", err)
        }

    case "-server":
        fmt.Println("Starting server...")
        cmd := exec.Command("go", "run", "tcp-server.go")
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        err := cmd.Run()
        if err != nil {
            fmt.Printf("Server exited with error: %v\n", err)
        }

    default:
        fmt.Println("Invalid argument. Use -client or -server")
    }
}