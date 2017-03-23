package libfss

import (
  "testing"
  "fmt"
)

func Test2PartyEq(t *testing.T) {
  // Generate fss Keys on client
  fClient := ClientInitialize(6)
  // Test with if x = 10, evaluate to 2
  fssKeys := fClient.GenerateTreePF(10, 2) 
  // Simulate server
  fServer := ServerInitialize(fClient.PrfKeys, fClient.NumBits) 

  // Test 2-party Equality Function
  cases := []struct {
    in uint; want int
  }{
    {95, 0},
    {10, 2},
    {1, 0},
  }
  for _, c := range cases {
    val0 := fServer.EvaluatePF(0, fssKeys[0], c.in)
    val1 := fServer.EvaluatePF(1, fssKeys[1], c.in)
    eval := val0 + val1 
    fmt.Printf("Evaluated equality point function over value %v.\nServer 1: %v\nServer2: %v\nTotal: %v\n\n",c.in,val0,val1,eval)
    if eval != c.want {
      t.Errorf("got %v, want %v", eval, c.want)
    }
  }
}

func Test2PartyLt(t *testing.T) {
  // Generate fss Keys on client
  fClient := ClientInitialize(6)
  // Test if x < 10, evaluate to 2
  fssKeysLt := fClient.GenerateTreeLt(10, 2)
  // Simulate server
  fServer := ServerInitialize(fClient.PrfKeys, fClient.NumBits)   

  // Test 2-party Less than Function
  cases := []struct {
    in, want uint
  }{
    {95, 0},
    {10, 0},
    {1, 2},
  }
  for _, c := range cases {
    val0 := fServer.EvaluateLt(fssKeysLt[0], c.in)
    val1 := fServer.EvaluateLt(fssKeysLt[1], c.in)
    eval := val0 - val1 
    fmt.Printf("Evaluated less than point function over value %v.\nServer 1: %v\nServer2: %v\nTotal: %v\n\n",c.in,val0,val1,eval)
    if eval != c.want {
      t.Errorf("got %v, want %v", eval, c.want)
    }
  }
}

func TestMultiPartyEq(t *testing.T) {
  // Generate fss Keys on client
  fClient := ClientInitialize(6)
  // Simulate server
  fServer := ServerInitialize(fClient.PrfKeys, fClient.NumBits) 
  // Test with if x = 10, evaluate to 2  
  fssKeysEqMP := fServer.GenerateTreeEqMP(10, 2, 3)

  // Test multiparty equal function case
  cases := []struct {
    in uint; want uint32
  }{
    {95, 0},
    {10, 2},
    {1, 0},
  }
  for _, c := range cases {
    val0 := fServer.EvaluateEqMP(fssKeysEqMP[0], c.in)
    val1 := fServer.EvaluateEqMP(fssKeysEqMP[1], c.in)
    val2 := fServer.EvaluateEqMP(fssKeysEqMP[2], c.in)
    eval := val0^val1^val2
    fmt.Printf("Evaluated multi-party equality point function over value %v.\nServer 1: %v\nServer2: %v\nServer3: %v\nTotal: %v\n\n",c.in,val0,val1,val2,eval)
    if eval != c.want {
      t.Errorf("got %v, want %v", eval, c.want)
    }
  }
}