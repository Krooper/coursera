package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Organize() {
	file, err := os.Open("C:\\Users\\kroop\\Downloads\\coursera\\uniqueization\\data_map.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileSorted, errSorted := os.Create("C:\\Users\\kroop\\Downloads\\coursera\\uniqueization\\sorted_data.txt")
	if errSorted != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileSorted.Close()

	var str string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		str += fileScanner.Text()
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	strSorted := SortString(str)
	for i := range strSorted {
		fmt.Fprintln(fileSorted, string(strSorted[i]))
	}
}
func Uniquize() {
	file, err := os.Open("C:\\Users\\kroop\\Downloads\\coursera\\uniqueization\\sorted_data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileUnique, errUnique := os.Create("C:\\Users\\kroop\\Downloads\\coursera\\uniqueization\\uniquized_data.txt")
	if errUnique != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileUnique.Close()

	fileScanner := bufio.NewScanner(file)
	var prev string
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			panic("File not sorted")
		}
		prev = txt
		fmt.Fprintln(fileUnique, txt)
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
}

func main() {
	Organize()
	Uniquize()
	fmt.Println("Done!")
}
