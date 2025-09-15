package main

import (
	"fmt"
)

func main() {
	// สร้าง Buffer channel ขนาด 10 ตัว
	ch := make(chan int, 10)

	// ส่งข้อมูลไปยัง Buffer channel
	for i := 0; i < 10; i++ {
		ch <- i

	}

	// รับข้อมูลจาก Buffer channel
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
