package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool adalah implementasi design pattern bernama object pool pattern
// design pattern Pool ini digunakan untuk menyimpan data, kemudian menggunakan datanya, setelah selesai bisa menyimpan kembali ke Pool nya
// Implementasi Pool di GoLang sudah aman dari problem race condition

func TestPool(t *testing.T) {
	// Membuat Pool dengan data defaulnya "New" ketika data Pool belum diisi, jika tidak diatur -? sync.Pool{}, defaultnya adalah nil
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Dira")
	pool.Put("Sanjaya")
	pool.Put("Wardana")

	// Karena hanya ada 10 data di Pool, sedangkan ada 10 goroutine, maka ada 7 goroutine yang tidak kebagian data (akan menerima data nil jika tidak ada defaultnya di Pool)
	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get() // ketika di get, data akan hilang dari Pool
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // Data kembali dimasukkan ke Pool
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}