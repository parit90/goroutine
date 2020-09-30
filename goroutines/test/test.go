package Test

import (
	"fmt"
	"time"
)

//export Hey, example for RACE condition. This Hey function should print "Hello", but it will print "Bye"
//This is a problem and should be avoided

//A race condition is when two or more routines have access to the same resource, such as a variable
//or data structure and attempt to read and write to that resource without any regard to the other routines.
func Hey() {
	var msg = "Hello"
	go func() {
		fmt.Println("msg from Hey: ", msg)
	}()

	msg = "Bye"
	//to see the output, otherwise main func will execute and exit, go routine will run in the background
	time.Sleep(100 * time.Millisecond)
}
