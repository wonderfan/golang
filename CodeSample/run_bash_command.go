package main

import "os/exec" //import the exec package
import "fmt"

func main() {
	output, err := exec.Command("pwd").Output()
	if err != nil {
		panic(" can not run the command")
	}
	fmt.Println("The command result is " + string(output))
}
