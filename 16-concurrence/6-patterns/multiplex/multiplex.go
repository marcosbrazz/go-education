package main

import (
	"fmt"
	"reflect"
	"time"
)

// join two channels
func multiplex(input_chan1, input_chan2 <-chan string) <-chan string {
	out_chan := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-input_chan1:
				out_chan <- msg1
			case msg2 := <-input_chan2:
				out_chan <- msg2
			}
		}
	}()

	return out_chan
}

// multiplex n channels
func multiplexN(input_chan ...<-chan string) <-chan string {
	output := make(chan string)

	go func() {
		cases := make([]reflect.SelectCase, len(input_chan))
		for i, channel := range input_chan {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(channel)}
		}
		for {
			chosen, value, opened := reflect.Select(cases)
			fmt.Printf("chosen channel %d\n", chosen)
			// channel := input_chan[chosen]
			msg := value.String()
			if opened {
				output <- msg
			}
		}
	}()
	return output
}

func main() {
	input1 := make(chan string)
	input2 := make(chan string)

	output := multiplexN(input1, input2)

	go func() {
		for i := 0; i < 5; i++ {
			input1 <- fmt.Sprintf("input1 %d", i)
			time.Sleep(time.Second / 4)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			input2 <- fmt.Sprintf("input2 %d", i)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-output)
	}

}
