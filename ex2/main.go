/*
Channel Exercise

Create a program that does the following:

has a buffered channel of type string with a capacity of 5
add 5 messages to the channel
have a function that ranges through the channel printing them out
the function should close a signal channel to allow the program to exit
close the messages channel to end the go routine previously launched
*/

package main

import (
	"fmt"
)

func process(messages chan string, quit chan struct{}) {
	// Loop through your messages
	for m := range messages {
		// print the message with a for loop using range
		fmt.Println(m)
	}
	quit <- struct{}{}
}

func main() {
	// declare the messages channel of type string and capacity of 5
	messages := make(chan string, 5)
	// declare a signal channel
	sigCh := make(chan struct{})
	// launch process in a goroutine
	go process(messages, sigCh)
	// declare 5 fruits in a []string
	fruits := []string{
		"apple",
		"banana",
		"kiwi",
		"pear",
		"peach",
	}
	// loop through the fruits and send them to the messages channel
	for _, f := range fruits {
		messages <- f
	}
	// close the messages channel
	close(messages)
	// wait for everything to finish
	<-sigCh
	fmt.Println("done")
}
