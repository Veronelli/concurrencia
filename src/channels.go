package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "mensaje 1"
	c <- "mensaje 2"

	fmt.Println(len(c), cap(c))

	//Range y close
	close(c)
	// c <- "mensaje 3"
	for messge := range c {
		fmt.Println(messge)

	}

	// Select
	email := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1 ", email)
	go message("mensaje 2 ", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("Email recibido de email 1: ", m1)

		case m2 := <-email2:
			fmt.Println("Email recibido de email 2: ", m2)

		}

	}

}
