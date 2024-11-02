package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Map mirip seperti GoLang Map, bedanya menggunakan concurrent dengan goroutine, dan sudah aman dari race condition
// Store(key, value) -> menyimpan data ke Map
// Load(key) -> mengambil data dari Map dengan key
// Delete(key) -> menghapus data di Map sesuai key
// Range(func(key, value)) -> melakukan iterasi seluruh data di Map

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
