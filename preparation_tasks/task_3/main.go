// Что выведет код?
package main

import (
	"fmt"
	"time"
)

// Если секунда успевает протикать быстрее горутины -> получаем interrupted signal и в finish будет true
// Если увеличить тикер до 2 секунд, он не успеет протикать, отработает горутина и в канал запишется false
// Во втором случае мы получим normal signal и в finish будет false

func main() {
	sync := make(chan bool, 1) // sync := make(chan bool) - делаем канал небуфферизованным для того,
	// чтобы можно было записать одно значение в неблокирующем режиме
	go func() {
		time.Sleep(time.Second)
		fmt.Println("get normal signal")
		sync <- false
	}()
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C: // если протикает 1 секунда
			fmt.Println("get interrupted signal")
			sync <- true
		case value := <-sync:
			fmt.Printf("finish %t", value)
			return
		}
	}
}
