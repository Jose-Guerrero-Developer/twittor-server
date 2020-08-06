package main

import (
	"log"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/core"
)

func main() {
	if err := core.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
