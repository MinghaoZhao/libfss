  package main

import (
  "os"
  "bufio"
  "strings"
  "github.com/gyuho/goraph"
  "strconv"
)

func main() {
  loc := "NY"

  // read in edges
  edge_data, _ := os.Open("./"+loc+"-travel-data.txt")
  defer edge_data.Close()
  edge_reader := bufio.NewScanner(edge_data)

  // create graph from edges
  g := goraph.NewGraph()
  for edge_reader.Scan() {
    edge := strings.Split(edge_reader.Text(), " ")
    if edge[0] == "a" {
      node1, err := g.GetNode(goraph.StringID(edge[1]))
      if err != nil {
        node1 = goraph.NewNode(edge[1])
        g.AddNode(node1)
      }
      node2, err := g.GetNode(goraph.StringID(edge[2]))
      if err != nil {
        node2 = goraph.NewNode(edge[2])
        g.AddNode(node2)
      }
      weight, _ := strconv.ParseFloat(edge[3], 32)
      g.AddEdge(node1.ID(), node2.ID(), weight)
    }
  }
  fmt.Println("Finished creating graph")

  // read in transit nodes
  tn_data, _ := os.Open("./"+loc+"-transit-nodes.txt")
  defer tn_data.Close()
  tn_reader := bufio.NewScanner(tn_data)
  
  *gcopy := *g
  var m map[goraph.Node]([]goraph.Node)
  for tn_reader.Scan() { 
    tn, err := goraph.GetNode(tn_reader.Text())
    
  }


/**
  writef, _ := os.Create("./"+loc+"-join-data.txt")
  writer := bufio.NewWriter(writef)

  for edge_reader.Scan() {
    line := strings.Split(creader.Text(), " ")
    if line[0] == "p" {
      writer.WriteString(line[4]+"\n")
    } else if line[0] == "v" {
      writer.WriteString(strings.Join(line[1:]," ")+"\n")
    }
  }

  for dreader.Scan() {
    treader.Scan()
    dline := strings.Split(dreader.Text(), " ")
    tline := strings.Split(treader.Text(), " ")
    if dline[0] == "p" {
      writer.WriteString(dline[3]+"\n")
    } else if dline[0] == "a" {
      top, _ := strconv.Atoi(dline[3])
      bottom, _ := strconv.Atoi(tline[3])
      speed := (top*15)/bottom
      writer.WriteString(strings.Join(dline[1:]," ")+" 1 "+strconv.Itoa(speed)+"\n")
      fmt.Println(speed)
    }
  }

  writer.Flush()
**/

}