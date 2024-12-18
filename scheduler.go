package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutine pertama untuk menjalankan ticker
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop() // Pastikan ticker dihentikan jika sudah tidak digunakan

		for {
			select {
			case t := <-ticker.C:
				// Task yang dijalankan setiap 5 detik
				fmt.Println("Ticker Task dijalankan pada:", t)
			}
		}
	}()

	// Goroutine kedua untuk menjalankan AfterFunc
	go func() {
		time.AfterFunc(10*time.Second, func() {
			// Task yang dijalankan setelah 10 detik
			fmt.Println("AfterFunc Task dijalankan pada:", time.Now())
		})
	}()

	// Blok utama agar program tetap hidup tanpa menggunakan time.Sleep
	select {}
}
