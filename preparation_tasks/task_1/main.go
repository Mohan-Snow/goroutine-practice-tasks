//Что будет выведено, и почему?

package main

import (
	"log"
	"time"
)

func main() {
	for _, v := range []int{1, 2, 3, 4, 5} {
		go func(v *int) {
			//println(*v)
			log.Println(*v) // log.Println(v) - принтанет хекс адресс на ячейку в памяти
		}(&v)
	}
	time.Sleep(2 * time.Second)
}
