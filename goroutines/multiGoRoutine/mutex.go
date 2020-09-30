package multiGoRoutine

import (
	"fmt"
	"sync"
)

var mycounter = 0
var Mt = sync.RWMutex{}

//sayonara export
func Mutex() {
	fmt.Println("Mutex :", mycounter)
	Mt.RUnlock()
	WG.Done()
}

//sayonara increment
func MuIncrement() {
	// Mt.Lock()
	mycounter++
	Mt.Unlock()
	WG.Done()
}
