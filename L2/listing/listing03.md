Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
```
type error - это инфтерфейс который хранит данные о типе и значение
nil - вывод значения
err == nil, false, так как мы проверяем указывает ли err на nil, а не на содержащиеся в нем значение