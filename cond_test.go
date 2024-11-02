package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Cond adalah implementasi locking berbasis kondisi, membutukan locker (Mutex atau RWMutex)
// Berbeda dengan locker biasa, Cond memiliki method Wait() untuk menunggu apakah perlu menunggu atau tidak
// Signal() -> untuk memberi tahu goroutine tidak perlu menunggu
// Broadcast() -> untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
// sync.NewCond(Locker) -> membuat Cond

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // akan melanjutkan satu goroutine
		}
	}()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	cond.Broadcast() // akan melanjutkan semua goroutine
	//}()

	group.Wait()
}
