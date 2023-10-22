package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	n, r, u *bool
	k       *int
}

func main() {
	ff := flags{}
	ff.k = flag.Int("k", 0, "указание колонки для сортировки")
	ff.n = flag.Bool("n", false, "сортировать по числовому значению")
	ff.r = flag.Bool("r", false, "сортировать в обратном порядке")
	ff.u = flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	fileName := os.Args[len(os.Args)-1]
	list, err := ReadFile(fileName, &ff)
	if err != nil {
		log.Fatal(err)
	}
	if *ff.r {
		sort.Sort(sort.Reverse(sort.StringSlice(list)))
	}
	if *ff.k != 0 {
		SortOnIndex(list, *ff.k-1)
	}
	if *ff.u {
		list = DeleteDuplicates(list)
	}
	for _, i := range list {
		fmt.Println(i)
	}
}

func ReadFile(fileName string, ff *flags) ([]string, error) {
	var buff, numBuff []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if *ff.n && sortNum(line) {
			numBuff = append(numBuff, line)
			continue
		} else {
			buff = append(buff, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	sort.Strings(buff)
	if *ff.n {
		sort.Strings(numBuff)
		buff = append(buff, numBuff...)
	}
	return buff, nil
}

func SortOnIndex(list []string, num int) {
	sort.Slice(list, func(i, j int) bool {
		tmp1 := strings.Split(list[i], " ")
		tmp2 := strings.Split(list[j], " ")
		if num >= len(tmp1) {
			return true
		}
		if num < len(tmp2) {
			return tmp1[num] < tmp2[num]
		}
		return false
	})
}

func sortNum(lines string) bool {
	if len(lines) > 0 {
		tmp := rune(lines[0])
		return unicode.IsDigit(tmp)
	}
	return false
}

func DeleteDuplicates(line []string) []string {
	newLine := make([]string, 0)
	for i := 0; i < len(line); i++ {
		if elemExists(newLine, line[i]) == false {
			newLine = append(newLine, line[i])
		}
	}
	return newLine
}

func elemExists(newLine []string, elem string) bool {
	for _, v := range newLine {
		if v == elem {
			return true
		}
	}
	return false
}
