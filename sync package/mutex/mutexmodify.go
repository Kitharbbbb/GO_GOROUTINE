package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Stock represents an item in a store
type Stock struct {
	quantity int
	mu       sync.Mutex
}

// Buy simulates a customer buying 1 item
func (s *Stock) Buy(id int) {
	s.mu.Lock() // ล็อกก่อนตรวจสอบหรือแก้ไขค่า
	defer s.mu.Unlock()

	if s.quantity > 0 {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // จำลองเวลาซื้อ
		s.quantity--
		fmt.Printf("Customer %d bought an item. Remaining: %d\n", id, s.quantity)
	} else {
		fmt.Printf("Customer %d tried to buy, but OUT OF STOCK!\n", id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	store := Stock{quantity: 5} // มีสินค้าเริ่มต้น 5 ชิ้น
	var wg sync.WaitGroup

	// มีลูกค้ามา 10 คน พยายามซื้อพร้อมกัน
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			store.Buy(id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All customers finished shopping.")
}
