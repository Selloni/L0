package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
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

todo:Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	k, n, r, u *bool
}

func main() {
	ff := flags{}
	ff.k = flag.Bool("k", false, "указание колонки для сортировки")
	ff.n = flag.Bool("n", false, "сортировать по числовому значению")
	ff.r = flag.Bool("r", false, "сортировать в обратном порядке")
	ff.u = flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	fileName := os.Args[len(os.Args)-1]
	list, err := ReadFile(fileName, &ff)
	if err != nil {
		log.Fatal(err)
	}
	//if *ff.n {
	//	sortNum(list)
	//}
	if *ff.r {
		sort.Sort(sort.Reverse(sort.StringSlice(list)))
	}
	for _, i := range list {
		fmt.Println(i)
	}
}

func ReadFile(fileName string, ff *flags) (buff []string, err error) {
	var numBuff []string
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
		fmt.Println(numBuff)
		buff = append(buff, numBuff...)
	}
	return buff, nil
}

func sortNum(lines string) bool {
	if len(lines) > 0 {
		tmp := rune(lines[0])
		return unicode.IsDigit(tmp)
	}
	return false
}

//func sortNum(data[]string) []string  {
//	sort.Slice(data, func(i, j int) bool {
//		return data[i] > data[j]
//	})
//	return data
//}
