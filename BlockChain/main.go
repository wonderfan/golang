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
  block, err := utils.UnmarshalBlock(data)
	if err != nil {
		fmt.Errorf("error unmarshaling to block: %s", err)
  }
  fmt.Println(utils.GetChainIDFromBlock(block))
  fmt.Println(" get what we want")
}
