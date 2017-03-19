package main

import (
  "os"
  "bufio"
  "strings"
  "fmt"
)

func main() {
  reformatEdgeGrid()
}

func reformatEdgeGrid() {
  readf, _ := os.Open("./NY_edge_grid.txt")
  defer readf.Close()
  reader := bufio.NewScanner(readf)
  writef, _ := os.Create("./NY_zones.txt")
  defer writef.Close()
  writer := bufio.NewWriter(writef)

  var edges = ""

  zone := "9" // first zone in the file
  for reader.Scan() {
    line := strings.Split(reader.Text(), " ")
    if line[0] != zone {
      writer.WriteString(zone+edges+"\n")
      zone = line[0]
      edges = ""
    } 
    edges += " "+strings.Join(line[1:4],"-")
  }
  writer.WriteString(zone+edges+"\n")
  writer.Flush()
}