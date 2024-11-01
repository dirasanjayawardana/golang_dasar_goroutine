package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Mutex (Mutual Exclusion) -> bisa digunakan untuk melakukan locking dan unlocking variabel
// sehingga hanya ada 1 goroutine yang bisa melakukan manipulasi data

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

// RWMutex digunakan untuk melakukan locking saat proses read dan write data
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // melakukan lock untuk proses write
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock() // melakukan lock hanya untuk proses read
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}

// Simulasi proses deadlock (go routine saling menunggu, sehingga tidak ada go routine yang berjalan)
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(pengirim *UserBalance, penerima *UserBalance, amount int) {
	pengirim.Lock()
	fmt.Println("Lock pengirim", pengirim.Name)
	pengirim.Change(-amount)

	time.Sleep(1 * time.Second)

	penerima.Lock()
	fmt.Println("Lock penerima", penerima.Name)
	penerima.Change(amount)

	time.Sleep(1 * time.Second)

	pengirim.Unlock()
	penerima.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Dira",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}