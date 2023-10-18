package main

import "testing"

// echo "gggj  hhhjlll" | cut
func TestWithoutFlag(t *testing.T) {
	argc := Args{
		fl: flags{
			f: 0,
			d: "\t",
			s: false,
		},
		str: []string{"gggj  hhhjlll"},
	}
	_, err := readString(&argc)
	if err == nil {
		t.Errorf("программа должна ожидать флаг")
	}
}

// echo "gg gj  hh hjll l"| cut -d " " -f 2
func TestFlagDF(t *testing.T) {
	argc := Args{
		fl: flags{
			f: 2,
			d: " ",
			s: false,
		},
		str: []string{"gg gj  hh hjll l"},
	}
	result, err := readString(&argc)
	if err != nil {
		t.Errorf("программа должна ожидать флаг")
	}
	for i := range result {
		if result[i] != "gj" {
			t.Errorf("ожидаю %s, получаю %s", result[i], "gj")
		}
	}
}

func TestSeveralStr(t *testing.T) {
	argc := Args{
		fl: flags{
			f: 2,
			d: " ",
			s: false,
		},
		str: []string{"gg gj  hh hjll l", "hh hjll l", "kkll", "Ya ko lo"},
	}
	result, err := readString(&argc)
	expect := []string{"gj", "hjll", "kkll", "ko"}
	if err != nil {
		t.Errorf("программа должна ожидать флаг")
	}
	for i := range result {
		if len(result) != len(expect) {
			t.Errorf("длина не совпадает res %d, exep %d", len(result), len(expect))
		}
		if result[i] != expect[i] {
			t.Errorf("ожидаю %s, получаю %s", result[i], expect[i])
		}
	}
}

func TestSeveralFlagS(t *testing.T) {
	argc := Args{
		fl: flags{
			f: 2,
			d: "j",
			s: true,
		},
		str: []string{"gg gjhh hjll l", "hh hjll l", "kkll", "Ya ko lo"},
	}
	result, err := readString(&argc)
	expect := []string{"hh h", "ll l"}
	if err != nil {
		t.Errorf("программа должна ожидать флаг")
	}
	for i := range result {
		if len(result) != len(expect) {
			t.Errorf("длина не совпадает res %d, exep %d", len(result), len(expect))
		}
		if result[i] != expect[i] {
			t.Errorf("ожидаю %s, получаю %s", result[i], expect[i])
		}
	}
}
