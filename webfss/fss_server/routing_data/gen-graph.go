/*
Take 3 parameters as input:
1) the file of the travel data for that region (from DIMACS site)
2) the file of the transit nodes for that region (from TRANSIT code)
3) the file you want to output zone data to
4) the file you want to output transit node shortest paths to
*/

package main

import (
  "os"
  "bufio"
  "strings"
  "github.com/gyuho/goraph"
  "strconv"
  "fmt"
  "log"
)

func main() {
  if len(os.Args) < 5 {
    fmt.Println("Missing parameter(s), provide file names!")
    return
  }

  g := makeGraph(os.Args[1])
  fmt.Println("Finished creating graph")
  tnMap := getTransitNodes(os.Args[2], g)

  tnMap = getZones(os.Args[3], tnMap, g) 
  fmt.Println(tnMap)
}

func makeGraph(travelDataFile string) goraph.Graph{
  // read in edges from file
  travel_data, err := os.Open(travelDataFile)
  checkErr(err)
  defer travel_data.Close()
  travel_reader := bufio.NewScanner(travel_data)

  // create graph from edges, using goraph
  g := goraph.NewGraph()
  for travel_reader.Scan() {
    edge := strings.Split(travel_reader.Text(), " ")
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
  return g
}

func getTransitNodes(transitNodeFile string, g goraph.Graph) map[goraph.Node]([]goraph.Node) {
  // read in transit nodes
  tn_data, err := os.Open(transitNodeFile)
  checkErr(err)
  defer tn_data.Close()
  tn_reader := bufio.NewScanner(tn_data)

  tnMap := make(map[goraph.Node]([]goraph.Node))
  for tn_reader.Scan() { 
    tn, err := g.GetNode(goraph.StringID(tn_reader.Text()))
    checkErr(err)
    tnMap[tn] = []goraph.Node{}
  }
  return tnMap
}

func getZones(loc string, tnMap map[goraph.Node]([]goraph.Node), g goraph.Graph) map[goraph.Node]([]goraph.Node) {
  // make queue for new nodes to expand
  type Next struct {
    enode goraph.Node // node to expand
    tnode goraph.Node // transit node enode corresponds to
    edge float64      // 
  }
  queue := make(chan Next, g.GetNodeCount()*2)

  // make map of node IDs to:
  //  - Node, if it hasn't been included in a zone yet
  //  - nil, if it has been included in a zone.
  nodeMap := g.GetNodes()

  // populate queue with existing transit nodes
  for tn, _ := range tnMap {
    queue <- Next{tn, tn, 0}
  }

  count := 0
  nodeCount := g.GetNodeCount()
  // repeat until all nodes have been added to zones
  for (count < nodeCount) {
    next := <- queue
    // Check if next node has been visited. If so, skip.
    if (nodeMap[next.enode.ID()] != nil) {
      // add target node to transit node zone mapping
      fmt.Println("adding node: ", next.enode, " to zone for tn: ", next.tnode," count: ", count)
      count++
      // fmt.Println("transit node mapping: ", tnMap[next.tnode])
      tnMap[next.tnode] = append(tnMap[next.tnode], next.enode)
      // get the children of target node, and add them to queue
      targetsMap, err := g.GetTargets(next.enode.ID())
      checkErr(err)
      for id, targetNode := range targetsMap {
        if (nodeMap[id] != nil) {
          queue <- Next{targetNode, next.tnode, 0}
        } // else, check if it's in the same zone (?) and add edges (?)
      }
      // mark next node as visited 
      nodeMap[next.enode.ID()] = nil
    } else {
      fmt.Println("Node ", next.enode, " has already been visited.")
    }
  }
  return tnMap
}

func checkErr(e error) {
  if e != nil {
    log.Fatal(e)
  }
}

func printMap(tnMap map[goraph.Node]([]goraph.Node)) {
  for tn, nodes := range tnMap {
    fmt.Printf("\n"+tn.String()+": ")
    for _, node := range nodes {
      fmt.Printf(node.String()+ " ")
    }
  }
  fmt.Printf("\n")
}
