package main

import (
  "fmt"
  "net"
  "bufio"
  "../libfss"
  "encoding/gob"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
  // Generate fss Keys on client
  fClient := libfss.ClientInitialize(6)
  // Test with if x = 10, evaluate to 2
  fssKeys := fClient.GenerateTreePF(10, 2)

  conn, _ := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
  enc := gob.NewEncoder(conn)
  enc.Encode(fssKeys)

/*
  // send to socket
  fmt.Fprintf(conn, "message to send\n")

  // listen for reply 
  message, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Print("Message from server: "+message)
*/
  
  // Close the connection when you're done with it.
  conn.Close()
  fmt.Println("done");
}