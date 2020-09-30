package main

import (
	"fmt"
	"runtime"
	"time"

	mulwg "github.com/parit90/goroutines/multiGoRoutine"
	ts "github.com/parit90/goroutines/test"
)

const name = "Parit"

func main() {
	go sayMyName()
	time.Sleep(100 * time.Millisecond)
	ts.Hey()
	ts.HeyResolve()
	ts.HeyWG()
	for i := 0; i < 10; i++ {
		mulwg.WG.Add(2)
		//go routine will race with each other, and there will be no syncronization between below both go-routine
		go mulwg.Sayonara()
		go mulwg.Increment()
	}
	mulwg.WG.Wait()

	//GOMAXPROCS provides us the threads available in goroutine
	//by passing the arg to GOMAXPROCS we can restrict goroutine on number of threads
	//eg. GOMAXPROCS(1) will gives us 1 thread
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		mulwg.WG.Add(2)
		mulwg.Mt.RLock()
		go mulwg.Mutex()
		mulwg.Mt.Lock()
		go mulwg.MuIncrement()
	}
	mulwg.WG.Wait()
}

//
func sayMyName() {
	fmt.Printf("my name is :%s\n", name)
}

// go run -race main.go
// run the application to check if there is any race condition
