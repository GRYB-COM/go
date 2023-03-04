package main

import (
	"fmt"
	"log"
	"tutoriale.com/greetings"
	
)

func main() {
	
	message,err := greetings.Powitanie("Sławeczek")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message)
}
