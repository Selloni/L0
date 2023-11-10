package main

import (
	"fmt"
	"log"
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
	buff, err := parsing(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buff))
}

// парсим строку
func parsing(str []rune) ([]rune, error) {
	var buff []rune
	var (
		tmp    rune
		err    error
		letter bool // проверка что буква встречается
	)
	for i := 0; i < len(str); i++ {
		n := 1
		if unicode.IsLetter(str[i]) {
			letter = true
			tmp = str[i] // сохраняем в переменную
			buff = append(buff, tmp)
		} else if unicode.IsDigit(str[i]) {
			if !letter {
				return nil, fmt.Errorf("не коректный ввод")
			}
			n, err = strconv.Atoi(string(str[i])) // переводим в привычный вид сроку, после в число,для цикла
			if err != nil {
				return nil, err
			}
			for ; n > 1; n-- { // изначально буква уже добавлена
				buff = append(buff, tmp) // добавляем в список количество
			}
		}
	}
	return buff, nil
}
