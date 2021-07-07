package main

import (
	"log"
	"main/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Printf("cmd.Execute err: %v", err)
	}
}
