package main

import (
	"log"

	"github.com/knight-zlm/eTerm/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd Execute err:%v", err)
	}
}
