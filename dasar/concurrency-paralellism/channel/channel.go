package channel

import (
	"fmt"
	"time"
)


func Demo()  {

	// messaging with channel
	fmt.Println("Start")

	// create channel
	messageCH := make(chan string)
	
	// receiver
	go func ()  {
		for i := 0; i < 5; i++ {
			fmt.Println("Receive message")
			messageData := <- messageCH
			fmt.Printf("Message received: %s\n", messageData)
		}
	}()


	// sender 
	go func () {
		fmt.Println("Sending message")
		messageCH <- "Hello, im sending message"

		// this will be sent due to matter of program stopped time
		time.Sleep(time.Second * 2) 
		messageCH <- "This is my message after 2 seconds"

		// this will be ignored
		time.Sleep(time.Second * 5) 
		messageCH <- "This is my message after 5 seconds"
		defer fmt.Println("Sender done")
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("Done")
	
}