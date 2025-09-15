package main

import (
	"fmt"
)

func main() {
	// สร้าง Channel ใหม่
	ch := make(chan int)

	// ส่งข้อมูลไปยัง Channel
	ch <- 10

	// รับข้อมูลจาก Channel
	v := <-ch

	// พิมพ์ข้อมูล
	fmt.Println(v)
}
