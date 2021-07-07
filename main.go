package main

import (
	"log"

	"github.com/cobra/tour/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Printf("cmd.Execute err: %v", err)
	}
}
