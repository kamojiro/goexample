package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("./tello", message, 0777)
	// file が存在すれば書き換え、
	// file が存在しなければ、perm の permission で新しいファイルを作る
	if err != nil {
		log.Fatal(err)
	}
}
