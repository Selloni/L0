package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestWget(t *testing.T) {
	url := "https://edu.21-school.ru/my-profile"
	// вызываю оригинальную команду
	expect := exec.Command("wget", "--no-directories", "--accept", "html", url)
	//expect.Stdout = os.Stdout
	//expect.Stderr = os.Stderr
	err := expect.Run()
	if err != nil {
		t.Log("Не выполнилась bash команда")
	}
	// вызываю мою команду
	if err := download("myProfile", url); err != nil {
		t.Error(err)
	}
	// считываю данные с файлов
	wget, err := os.ReadFile("my-profile.tmp")
	if err != nil {
		log.Fatal(err)
	}
	myprog, err := os.ReadFile("myProfile")
	if err != nil {
		log.Fatal(err)
	}
	// сраниваю их
	if string(wget) != string(myprog) {
		t.Errorf("Данные отличаються")
	}
	// удаляем не нужные файлы
	err = os.Remove("my-profile.tmp")
	if err != nil {
		t.Log("Не выполнилась удаление my-profile.tmp")
	}
	err = os.Remove("myProfile")
	if err != nil {
		t.Log("Не выполнилась удаление myProfile")
	}
}
