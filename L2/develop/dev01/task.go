package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	tt, err := MyTime()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Текущие время: %v\n", tt)
}

func MyTime() (time.Time, error) {
	tm, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, fmt.Errorf("Не удалось получить текущие время")
	}
	return tm, err
}
