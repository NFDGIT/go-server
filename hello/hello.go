package demo

import (
	"log"
	"fmt"
	"example.com/greetings"
)

func Demo() {

	message,err := greetings.Hello("peng")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
