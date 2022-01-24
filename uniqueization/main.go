package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data_map.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	alreadySeenData := make(map[byte]bool)
	var str string
	for {
		_, err := file.Read(data)
		for i := range data {
			if _, found := alreadySeenData[data[i]]; found {
				continue
			} else {
				alreadySeenData[data[i]] = true
				str += string(data[i])
			}
		}
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
	}
	fileNew, errNew := os.Create("out_data.txt")

	if errNew != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileNew.Close()

	for i := range str {
		fmt.Fprintln(fileNew, string(str[i]))
	}
	fmt.Println("Done")
}
