package main

import (
	"fmt"
	"log"
	"time"

	"github.com/raja-dettex/dap-client/client"
)

func main() {
	// test cases
	client, err := client.New("localhost", "3000")
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.HandleInsert("users", "raja", map[string]any{"name": "raja", "age": 24})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("response POST %v", res)

	time.Sleep(time.Second * 3)
	resD, err := client.HandleSelect("users", "raja")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("response GET %v", resD)
}
