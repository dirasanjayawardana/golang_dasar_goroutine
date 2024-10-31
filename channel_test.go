package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // Membuat channel
	defer close(channel) // Menutup channel ketika sudah tidak digunakan

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dira Sanjaya Wardana" // mengirim data ke channel
		fmt.Println("Selesai Mengirim Data ke Channel")
	}() // tanda () langsung mengeksekusi anonymous function

	data := <-channel // mengambil data dari channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Ketika Channel digunakan sebagai parameter pada sebuah function, maka otomatis akan pass by refference, tidak perlu menggunakan pointer untuk channel
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dira Sanjaya Wardana"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// secara default ketika channel dijadikan parameter sebuah func, bisa menerima dan mengirim data ke channel
// membuat parameter channel yang hanya bisa mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Dira Sanjaya Wardana"
}

// membuat parameter channel yang hanya bisa menerima data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// secara default channel hanya bisa menerima satu data, jika menambahkan data ke 2 akan menunggu sampai data ke 1 ada yg mengambil
// Buffered channel digunakan untuk menampung data antrian di Channel
// Data yang diambil, urutannya sesuai dengan data yg dimasukkan
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // membuat channel yang bisa menampung sampai 3 data
	defer close(channel)

	fmt.Println(cap(channel)) // melihat kapasitas channel
	fmt.Println(len(channel)) // melihat banyaknya data yang sedang disimpan di channel

	go func() {
		channel <- "Dira"
		channel <- "Sanjaya"
		channel <- "Wardana"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

// Ketika data dikirim ke channel secara terus menerus, untuk melakukan pengecekan bisa menggunakan for range
// ketika channel di close(), secara otomatis perulangan tersebut akan berhenti
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

// select digunakan untuk mendapatkan data dari beberapa channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}