package main

import (
	"fmt"
)

func main() {
	// สร้าง Channel ใหม่
	ch := make(chan int, 1)

	// ส่งข้อมูลไปยัง Channel
	ch <- 1

	// รับข้อมูลจาก Channel
	v := <-ch

	// พิมพ์ข้อมูล
	fmt.Println(v)
}
