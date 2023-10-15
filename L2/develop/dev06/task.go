package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	f *int
	d *string
	s *bool
}

type Args struct {
	fl  flags
	str []string
}

func main() {
	argc := getArgc()
	fmt.Print(readString(*argc))
}

func readString(argc Args) []string {
	var output []string
	for _, str := range argc.str {
		arr := strings.Split(str, *argc.fl.d)
		output = append(output, arr[*argc.fl.f-1])
	}
	return output
}

func parsFlag(fl *flags) {
	fl.f = flag.Int("f", 0, "fields - выбрать поля (колонки)")
	fl.d = flag.String("d", "	", "delimiter - использовать другой разделитель")
	fl.s = flag.Bool("s", true, "separated - только строки с разделителем")
	flag.Parse()
}

func getArgc() *Args {
	var fl flags
	parsFlag(&fl)

	return &Args{
		fl: fl,
		//str: os.Args[len(os.Args)-1], // одну строку
		str: flag.Args(), // передаю все строки после команды и флагов

	}
}
