package sample

import (
	"log"

	"github.com/nats-io/nats"
)

func sendMessage(URL string) error {

	nc, err := nats.Connect(URL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer nc.Close()

	nc.Publish("Hello", []byte("World"))
	nc.Flush()

	log.Printf("Published [%s] : '%s'\n", "Hello", "World")

	return nil
}
