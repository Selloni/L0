package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
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
	out, err := ReadFile(files, fl)
	if err != nil {
		log.Fatal(err)
	}
	Output(out, fl)

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

func ReadFile(strIn []string, fl flags) ([]string, error) {
	outStr := make([]string, 0)
	var (
		fullFile []string
	)
	flagBC := math.Max(float64(*fl.C), float64(*fl.B))
	flagAC := math.Max(float64(*fl.C), float64(*fl.A))
	flagI, flagN := "", ""
	if *fl.i {
		flagI = "(?i)"
	}
	for _, path := range strIn[1:] {
		regx := strIn[0]
		NumLine := 1
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		if *fl.F {
			regx = regexp.QuoteMeta(regx)
		}
		pattern, err := regexp.Compile(flagI + regx)

		if err != nil {
			return nil, err
		}
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			if *fl.n {
				flagN = (strconv.Itoa(NumLine) + ":")
			}
			find := pattern.MatchString(scan.Text())
			if flagBC > 0 {
				fullFile = append(fullFile, scan.Text())
			}
			if find && !*fl.v {
				if flagBC > 0 {
					outStr = append(outStr, appendBefore(fullFile, int(flagBC), NumLine)...)
				}
				outStr = append(outStr, flagN+scan.Text())
				if flagAC > 0 {
					outStr = append(outStr, appendAfter(scan, int(flagAC))...)
				}
				if flagAC+flagBC > 0 {
					outStr = append(outStr, "--")
				}
			} else if !find && *fl.v {
				outStr = append(outStr, flagN+scan.Text())
			}
			NumLine++
		}
	}
	if flagAC+flagBC > 0 { // костыль
		outStr = outStr[:len(outStr)-1]
	}
	return outStr, nil
}

func Output(out []string, fl flags) {
	if *fl.c {
		fmt.Println(len(out))
	} else {
		for _, i := range out {
			fmt.Println(i)
		}
	}
}

func appendBefore(fullFile []string, countBefore int, numLine int) []string {
	var tmp []string
	for i := countBefore + 1; i > 1; i-- {
		if numLine-i > -1 {
			tmp = append(tmp, fullFile[numLine-i])
		}
	}
	return tmp
}

func appendAfter(scan *bufio.Scanner, flagAC int) []string {
	var tmp []string
	for scan.Scan() {
		if flagAC != 0 {
			tmp = append(tmp, scan.Text())
			flagAC--
			if flagAC == 0 {
				break
			}
		}
	}
	return tmp
}
