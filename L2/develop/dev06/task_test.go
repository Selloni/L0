package main

import "testing"

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
		t.Errorf("программа должна одидать флаг")
	}
}
