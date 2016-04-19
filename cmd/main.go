package main

import (
	"os"

	"github.com/jduhamel/sample"
)

func main() {
	err := sample.SendMessage("nats://localhost:6222")
	if err != nil {
		os.Exit(-1)
	}
}
