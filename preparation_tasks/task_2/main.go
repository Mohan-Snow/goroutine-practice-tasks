// Разобрать код: Что будет если GOMAXPROCS установить больше 1.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
func (c *Common) Worker(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		c.Counter++
	}
	c.Map[index] = c.Counter
}

func main() {
	runtime.GOMAXPROCS(1)
	obj := &Common{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go obj.Worker(i, wg)
	}
	wg.Wait()
	fmt.Println(obj.Counter)
}
*/

type Common struct {
	sync.Mutex
	Map     map[int]int
	Counter int
}

// добавить в функцию Worker c.Mutex.Lock() и c.Mutex.Unlock(), чтобы горутины не имели конкурентного доступа к мапе

func (c *Common) Worker(index int, wg *sync.WaitGroup) {
	defer wg.Done() // сообщаем счетчику, что эта горутиа отработала ( нужно сделать отсчет)
	c.Mutex.Lock()
	for i := 0; i < 1000; i++ {
		c.Counter++
	}
	c.Map[index] = c.Counter
	c.Mutex.Unlock()
}

func main() {
	runtime.GOMAXPROCS(3)                        // указываем количество используемых ядер
	obj := &Common{Map: make(map[int]int, 1000)} // obj := &Common{} - panic: assignment to entry in nil map
	wg := &sync.WaitGroup{}                      // работает по типу CountDownLatch
	for i := 0; i < 100; i++ {
		wg.Add(1) // инкрементим счетчик обратного отсчета на 1
		go obj.Worker(i, wg)
	}
	wg.Wait() // ждем, когда счетчик обнулится
	fmt.Println(obj.Counter)
}
