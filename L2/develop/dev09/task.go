package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func download(fileName, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	url := flag.String("url", "", "укажите адрес который нужно скачать")
	flag.Parse()
	if *url == "" {
		log.Fatalln("нужен url адрес")
	}
	fileName := path.Base(*url)
	download(fileName, *url)
}
