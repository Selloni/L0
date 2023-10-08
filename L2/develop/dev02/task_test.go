package main

import (
	"fmt"
	"testing"
)

func TestWrongArgument(t *testing.T) {
	str := []rune{'2', 'f', '2', 'h'}
	_, err := Parsing(str)
	if err != fmt.Errorf("не коректный ввод") {
		t.Error("Ошибка не отработала")
	}
}

func TestOne(t *testing.T) {
	str := []rune{'f', '3', 'h', '3'}
	result := []rune{'f', 'f', 'f', 'h', 'h', 'h'}
	pp, _ := Parsing(str)
	for i := 0; i < len(pp); i++ {
		if pp[i] != result[i] {
			t.Errorf("Нет совпадения по индексу %d, значение %v, %v", i, pp[i], result[i])
		}
	}

}
