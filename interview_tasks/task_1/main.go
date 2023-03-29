package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// close(ch)
	}()

	for n := range ch {
		// после завершения цикла в горутине, range пытается вычитывать значения из пустого, незакрытого канала
		// вследствии этого возникает deadlock
		// чинится закрытием канала (выше в горутине)
		fmt.Println(n)
	}
}
