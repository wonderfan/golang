package main

import (
  "fmt"
  "io/ioutil"
  utils "github.com/hyperledger/fabric/protos/utils"
)

func main(){
  fmt.Println("Use codes to verify fabric")
  blockFile := "/root/go/src/github.com/hyperledger/fabric/examples/e2e_cli/channel-artifacts/genesis.block"
  data, err := ioutil.ReadFile(blockFile)
  if err != nil {
		fmt.Errorf("Could not read block %s", blockFile)
  }

  fmt.Println(utils.GetChainIDFromBlockBytes(block))
  fmt.Println(" get what we want")
}
