package Test

import (
	"fmt"
	"sync"
)

//wait group syncronize the go-routine
var wg = sync.WaitGroup{}

//HeyWG using the wait group we can syncronize the go routine between main and HeyWG go routine
//using wg we can remove the time.sleep method
func HeyWG() {
	var msg = "Hello"

	wg.Add(1)
	go func(msg string) {
		fmt.Println("msg from HeyWG: ", msg)
		wg.Done()
	}(msg)

	msg = "Bye"
	//to see the output, otherwise main func will execute and exit, go routine will run in the background
	wg.Wait()
}
