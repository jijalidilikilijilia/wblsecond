Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Случайный вывод эементов двух каналов и затем бесконечный вывод нулей  
```
Вывод такой потому что после того как значения из обоих каналов a и b будут прочитаны, select в функции merge будет продолжать работать в бесконечном цикле for который никогда не завершится и будет продолжать постоянно читать данные из уже закрытого канала. При чтении из закрытого канала возвращается стандартное значение типа этого канала, в данном случае 0