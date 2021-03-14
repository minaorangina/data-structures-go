package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/minaorangina/data-structures-go/linkedlist"
)

var usage = `here's some instructions lol`

func main() {
	values := os.Args[1:]

	if len(values) < 1 {
		log.Fatalf("enter at least one integer")
	}

	var numbers []int
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("could not parse %v", num)
		}

		numbers = append(numbers, num)
	}

	ll := linkedlist.NewLinkedList()

	for _, n := range numbers {
		ll.AddNode(n)
	}
	ll.Display()

	// wait for input
	inputCh := make(chan string)
	go func() {

	}()

	interruptCh := make(chan os.Signal)
	signal.Notify(interruptCh, os.Interrupt, syscall.SIGTERM)

	for {
		displayUsage(os.Stdout)
		select {
		case input := <-inputCh:
			// work out what they want
			fmt.Printf("you asked for %s\n", input)
		case sig := <-interruptCh:
			fmt.Printf("got %s signal, aborting...\n", sig)
			return
		}
	}
}

func displayUsage(w io.Writer) {
	fmt.Fprint(w, usage)
}
