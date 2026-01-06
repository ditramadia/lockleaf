package main

import (
	"log"

	"github.com/ditramadia/lockleaf/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
