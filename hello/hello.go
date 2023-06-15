package hello

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func Demo() {

	message, err := greetings.Hello("peng")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
