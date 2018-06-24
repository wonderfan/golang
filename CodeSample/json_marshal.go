package main

import "fmt"
import "encoding/json"
import "os"

type Comupter struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

func main() {
	fmt.Println("Hello, 世界")
	computer := Comupter{Name: "hp", Platform: "window"}
	jsonBytes, err := json.Marshal(computer)
	fmt.Println("a", "b")

	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(jsonBytes)
	fmt.Println(string(jsonBytes))
}
