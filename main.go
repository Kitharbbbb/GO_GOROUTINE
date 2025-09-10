package main

import (
	"fmt"
	"time"
)

func main() {
	// สร้าง Goroutines ใหม่ 2 ตัว
	go doSomething1()
	go doSomething2()

	// รอให้ Goroutines ทั้งสองตัวทำงานเสร็จสิ้น
	time.Sleep(time.Second)

	// พิมพ์ข้อความ
	fmt.Println("Done")
}

func doSomething1() {
	fmt.Println("Doing something 1")
}

func doSomething2() {
	fmt.Println("Doing something 2")
}
