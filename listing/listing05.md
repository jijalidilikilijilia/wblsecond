Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	} 
	
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Программа выведет "error", потому что она проверяет переменную err и выводит "error", если err не равна nil. В данном случае, функция test() всегда возвращает nil, поэтому условие if err != nil не выполняется, и программа выводит "error".