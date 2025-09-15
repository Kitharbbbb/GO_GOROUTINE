package main

import (
	"fmt"
)

func main() {
	// สร้าง Channel ใหม่
	ch := make(chan int)

	// สร้าง Goroutine ใหม่
	go func() {
		// ส่งข้อมูลไปยัง Channel
		ch <- 1
	}()

	// รับข้อมูลจาก Channel
	v := <-ch

	// พิมพ์ข้อมูล
	fmt.Println(v)
}
