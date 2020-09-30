package multiGoRoutine

import (
	"fmt"
	"sync"
)

var counter = 0
var WG = sync.WaitGroup{}

//sayonara export
func Sayonara() {
	fmt.Println("Sayonara :", counter)
	WG.Done()
}

//sayonara increment
func Increment() {
	counter++
	WG.Done()
}
