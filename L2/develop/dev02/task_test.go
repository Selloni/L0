package main

import (
	"testing"
)

func TestWrongArgument(t *testing.T) {
	str := []rune{'2', 'f', '2', 'h'}
	_, err := Parsing(str)
	if err == nil {
		t.Errorf("Ошибка не отработала: %v", err)
	}
}

func TestOne(t *testing.T) {
	str := []rune{'f', '3', 'h', '3'}
	result := []rune{'f', 'f', 'f', 'h', 'h', 'h'}
	pp, _ := Parsing(str)
	for i := 0; i < len(result); i++ {
		if pp[i] != result[i] {
			t.Errorf("Нет совпадения по индексу %d, значение %v, %v", i, string(pp[i]), string(result[i]))
		}
	}
}

func TestTwo(t *testing.T) {
	str := []rune{'f', '2', 'h', '3', 'L', 'K', '1'}
	result := []rune{'f', 'f', 'h', 'h', 'h', 'L', 'K'}
	pp, _ := Parsing(str)
	for i := 0; i < len(result); i++ {
		if pp[i] != result[i] {
			t.Errorf("Нет совпадения по индексу %d, значение %v, %v", i, string(pp[i]), string(result[i]))
		}
	}
}

func TestThree(t *testing.T) {
	str := []rune{'р', '3', 'ы', 'щ', 'щ', '2'}
	result := []rune{'р', 'р', 'р', 'ы', 'щ', 'щ', 'щ'}
	pp, _ := Parsing(str)
	for i := 0; i < len(result); i++ {
		if pp[i] != result[i] {
			t.Errorf("Нет совпадения по индексу %d, значение %v, %v", i, string(pp[i]), string(result[i]))
		}
	}
}
func TestFour(t *testing.T) {
	str := []rune{'得', '到', '2', 'y', 'e', 's', '4'}
	result := []rune{'得', '到', '到', 'y', 'e', 's', 's', 's', 's'}
	pp, _ := Parsing(str)
	for i := 0; i < len(result); i++ {
		if pp[i] != result[i] {
			t.Errorf("Нет совпадения по индексу %d, значение %v, %v", i, string(pp[i]), string(result[i]))
		}
	}
}
