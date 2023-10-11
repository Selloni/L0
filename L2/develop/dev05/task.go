package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	A *int
	B *int
	C *int
	c *bool
	i *bool
	v *bool
	F *bool
	n *bool
}

func main() {
	fl := flags{}
	ParsFlag(&fl)

	files := flag.Args()
	out, err := ReadFile(files[1:], fl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*fl.v)
	//fmt.Printf("%v:", files[1])
	for _, i := range out {
		fmt.Println(i)
	}
}

func ParsFlag(fl *flags) {
	fl.A = flag.Int("A", 0, " печатать +N строк после совпадения")
	fl.B = flag.Int("B", 0, "печатать +N строк до совпадения")
	fl.C = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	fl.c = flag.Bool("c", false, "количество строк")
	fl.i = flag.Bool("i", false, "игнорировать регистр")
	fl.v = flag.Bool("v", false, "вместо совпадения, исключать")
	fl.F = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	fl.n = flag.Bool("n", false, "тпечатать номер строки")
	flag.Parse()

}

func ReadFile(files []string, fl flags) ([]string, error) {
	outStr := make([]string, 0)
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		pattern, err := regexp.Compile(os.Args[1])
		if err != nil {
			return nil, err
		}
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			if pattern.MatchString(scan.Text()) && !*fl.v {
				outStr = append(outStr, scan.Text())
			} else if *fl.v {
				outStr = append(outStr, scan.Text())
			}
		}
	}
	return outStr, nil
}
