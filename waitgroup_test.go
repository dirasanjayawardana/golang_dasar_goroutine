package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Wait group -> fitur untuk menunggu sebuah proses selsai dilakukan (mirip seperti await di javaScript)
// untuk menandai ada proses go routine dalam sebuah func, dengan menerima parameter group, dan menambahkan method di func -> Add(jumlahProses)
// setelah proses goroutine selesai, bisa menggunakan method -> Done()
// Untuk menunggu semua proses selesai, menggunakan method -> Wait()

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1) // menandakan ada satu proses yang harus ditunggu

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait() // menunggu semua proses selesai, baru melanjutkan code setelahnya
	fmt.Println("Selesai")
}