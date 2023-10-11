package main

import (
	"reflect"
	"testing"
)

func TestAnagram_1(t *testing.T) {
	world := []string{"пятак", "пятка", "тяПка", "тяпка", "тЕрка", "терк",
		"Листок", "сЛИток", "столик", "лист", "литс"}
	expect := map[string][]string{
		"листок": {"слиток", "столик"},
		"пятак":  {"пятка", "тяпка"},
	}
	result := findAnagramme(world)
	if !(reflect.ValueOf(expect).Type() == reflect.ValueOf(result).Type()) {
		t.Errorf("не свопадают типы")
	}
	for k, v := range expect {
		for i := range v {
			if result[k][i] != v[i] {
				t.Errorf("Данные не совпали, ожидаем - %v, поулчили - %v", expect, result)
			}
		}
	}
}

func TestFindAnagram_2(t *testing.T) {
	words := []string{"cat", "dog", "tac", "god", "act"}
	expected := map[string][]string{
		"cat": {"act", "tac"},
	}
	result := findAnagramme(words)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Данные не совпали, ожидаем - %v, поулчили - %v", expected, result)
	}
}
