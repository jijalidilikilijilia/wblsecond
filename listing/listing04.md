Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Случайный вывод чисел от 0 до 9 и затем deadlock
```
Программа создает канал ch, отправляет числа от 0 до 9 в этот канал, а затем пытается прочитать и вывести числа из канала. Однако она не закрывает канал после отправки значений, и это приводит к блокировке программы.

Вывод программы: программа заблокируется и не выведет ничего на экран.

Это происходит потому, что главная горутина бесконечно ждет, что будут отправлены еще значения в канал, но отправляющая горутина уже завершила выполнение и больше не отправляет значения. Чтобы избежать блокировки треюуется явно сказать, что значений больше ждать не следует 