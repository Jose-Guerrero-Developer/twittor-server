package main

import (
	"log"

	"github.com/devJGuerrero/twittor-server/galex/core"
)

func main() {
	if err := core.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
