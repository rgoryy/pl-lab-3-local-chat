package main

import (
  "bufio"
  "fmt"
  "io"
  "net"
  "strings"
)

type Client struct {
  conn net.Conn
  nick string
}

var clients = make(map[string]*Client)

func handleConnection(conn net.Conn) {
  reader := bufio.NewReader(conn)

  fmt.Println("New connection!")

  msg, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("Error reading user nickname from msg: ", err)
    return
  }

  respMsg := "Greetings! " + strings.ToUpper(msg)
  conn.Write([]byte(respMsg + "\n"))

  client := &Client{conn: conn, nick: "@" + msg[:len(msg)-1]}

  clients[client.nick] = client

  fmt.Println(client.nick)

  go handleMessages(client)
}

func handleMessages(client *Client) {
  defer client.conn.Close()

  reader := bufio.NewReader(client.conn)
  for {
    msg, err := reader.ReadString('\n')
    if err != nil {
      if err != io.EOF {
        fmt.Println("Error reading message:", err)
        break
      }
    }

    msg = strings.TrimSpace(msg)

    splited := strings.SplitN(msg, " ", 2)
    firstPart := splited[0]

    if strings.HasPrefix(firstPart, "@") {
      if clientToSend, ok := clients[firstPart]; ok {
        msg = splited[1]
        clientToSend.conn.Write([]byte(fmt.Sprintf("%s: %s\n", client.nick, msg)))
      }
    } else {
      for _, c := range clients {

        if c.nick == client.nick {
          continue
        }

        fmt.Println("______")
        fmt.Println("SENDING FROM")
        fmt.Println(client.nick)
        fmt.Println("TO")
        fmt.Println("SENDING FROM")
        fmt.Println(c.nick)
        fmt.Println("______")

        c.conn.Write([]byte(fmt.Sprintf("%s: %s\n", client.nick, msg)))
      }
    }
    client.conn.Write([]byte(fmt.Sprintf("%s\n", "SENT")))
  }

  delete(clients, client.nick)
  fmt.Printf("User %s disconnected\n", client.nick)
}

func main() {

  fmt.Println("Launching server...")

  ln, err := net.Listen("tcp", ":8081")
  if err != nil {
    fmt.Println("Error starting server:", err)
    return
  }
  defer ln.Close()

  fmt.Println("Server started on port 8081")

  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Error accepting connection:", err)
      continue
    }

    go handleConnection(conn)
  }

}