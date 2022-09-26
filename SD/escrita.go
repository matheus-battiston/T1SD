package main

import (
	"fmt"
	"log"
	"os"
)

func Writeasdasd(fileName string, message []string) {

	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	for i := 0; i < len(message); i++ {
		_, err2 := f.WriteString(message[i] + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
	}
	fmt.Println("done")
}

func main() {
	x := []string{"oi", "matheus"}
	Writeasdasd("teste.txt", x)
}
