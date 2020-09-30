package Test

import (
	"fmt"
	"time"
)

//HeyResolve This is still unreliable solution because we are using time.Sleep method
func HeyResolve() {
	var msg = "Hello"
	go func(msg string) {
		fmt.Println("msg from HeyResolve: ", msg)
	}(msg)

	msg = "Bye"
	//to see the output, otherwise main func will execute and exit, go routine will run in the background
	time.Sleep(100 * time.Millisecond)
}
