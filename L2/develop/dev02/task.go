package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var (
		tmp string
		str []rune
	)
	fmt.Println("напиши последовательность")
	fmt.Fscan(os.Stdin, &tmp)
	str = []rune(tmp)
	Parsing(str)
	//fmt.Println(str)
}

func Parsing(str []rune) {
	buff := make([]rune, len(str))
	var tmp rune
	for i := 0; i < len(str); i++ {
		n := 1
		if unicode.IsLetter(str[i]) {
			tmp = str[i]
			buff = append(buff, tmp)
		} else if unicode.IsDigit(str[i]) {
			n, _ = strconv.Atoi(string(str[i]))
			fmt.Println(n)
			for ; n > 0; n-- {
				buff = append(buff, tmp)
			}
		}

	}
	fmt.Println(string(buff))
}
