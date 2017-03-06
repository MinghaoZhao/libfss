package main

import (
  "os"
  "strconv"
  "bufio"
  "strings"
  "../libfss"
  "fmt"
)

// Runs FSS on (matches) the 1st element of line (key), returns the 3rd element (val).
func readFetchSmallValue(server *libfss.Fss, serverNum byte, fssKey libfss.FssKeyEq2P) string {
  var ans int = 0
  file, _ := os.Open("./routing_data/NY_edge_grid.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)

  // Read file line by line, on each line evaluate PF on node id
  for scanner.Scan() {
    line := strings.Split(scanner.Text(), " ")
    key, _ := strconv.Atoi(line[1])
    val, _ := strconv.Atoi(line[3])
    ans += server.EvaluatePF(serverNum, fssKey, uint(key))*val
  }
  return strconv.Itoa(ans)
}

// Runs FSS on (matches) the 0th element of line, returns the 1st element.
// Answer is base64-encoded int array, where each int represents an edge.
func readFetchLargeValue(server *libfss.Fss, serverNum byte, fssKey libfss.FssKeyEq2P) string {
  ans := make([]int, 1000) // TODO: make smarter
  file, _ := os.Open("./routing_data/test.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)

  // Read file line by line, on each line evaluate PF on key
  for scanner.Scan() {
    // line := strings.Split(scanner.Text(), " ")
    split := strings.SplitAfterN(scanner.Text(), " ", 2)
    fmt.Println("split: ", split)
    key, _ := strconv.Atoi(strings.TrimSpace(split[0]))

    byteArray := []byte(split[1])
    fssVal := server.EvaluatePF(serverNum, fssKey, uint(key))
    for i := range byteArray {
      ans[i] += int(byteArray[i]) * fssVal
    }

  }
  fmt.Println("answer before transmit: ", ans)
  transmit := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ans)), ","), "[]")
  return transmit
}