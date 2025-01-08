package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
  "sync"
)

func main() {
  conn, err := net.Dial("tcp", "127.0.0.1:8081")
  if err != nil {
    fmt.Println("Error connecting:", err)
    return
  }
  defer conn.Close()

  fmt.Print("Введите свой ник: ")
  reader := bufio.NewReader(os.Stdin)
  nick, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("Error reading nickname:", err)
    return
  }
  nick = strings.TrimSpace(nick)

  sendMessage(conn, nick)

  var wg sync.WaitGroup

  wg.Add(1)
  go receiveMessage(conn, &wg)

  for {
    fmt.Print("Text to send: ")
    reader := bufio.NewReader(os.Stdin)
    msg, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("Error reading message:", err)
      return
    }
    msg = strings.TrimSpace(msg)

    if msg == "exit" {
      fmt.Println("Exiting...")
      break
    }

    sendMessage(conn, msg)
  }

  wg.Wait()
}

func sendMessage(conn net.Conn, msg string) {
  fmt.Println("Sending message:", msg)
  _, err := fmt.Fprintf(conn, msg+"\n")
  if err != nil {
    fmt.Println("Error sending message:", err)
  }
}

func receiveMessage(conn net.Conn, wg *sync.WaitGroup) {
  defer wg.Done()

  for {
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
      fmt.Println("Error reading message from server:", err)
      return
    }
    fmt.Println(message)
  }
}