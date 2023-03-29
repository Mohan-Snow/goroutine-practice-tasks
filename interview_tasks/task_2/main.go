package main

import (
	"log"
	"time"
)

func main() {
	// в этом примере нет никакого таймаута, поэтому main горутина успевает отработать раньше, чем анонимные горутины в цикле
	for i := 0; i < 5; i++ {
		//log.Printf("MAIN goroutine I: %d", i)
		i := i
		go func() {
			log.Printf("Anonymous goroutine I: %d", i)
		}()
	}
	time.Sleep(time.Second)
}

// Недетерминированное поведение (исп time.Sleep(10000))

// 2023/03/29 19:32:12 MAIN goroutine I: 0
// 2023/03/29 19:32:12 MAIN goroutine I: 1
// 2023/03/29 19:32:12 Anonymous goroutine I: 1
// 2023/03/29 19:32:12 Anonymous goroutine I: 1
// 2023/03/29 19:32:12 MAIN goroutine I: 2
// 2023/03/29 19:32:12 Anonymous goroutine I: 2
// 2023/03/29 19:32:12 MAIN goroutine I: 3
// 2023/03/29 19:32:12 Anonymous goroutine I: 3
// 2023/03/29 19:32:12 MAIN goroutine I: 4
// 2023/03/29 19:32:12 Anonymous goroutine I: 4

// Если добавить переназначение итератора перед пробрасыванием в горутину i := i,
// можно поправить ситуацию с недетерминированным поведением
