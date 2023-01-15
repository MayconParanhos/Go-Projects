package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./list.txt")
	check(err)

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())

	for i := 0; i < len(lines); i++ {
		fmt.Printf("\n  %s \n", lines[i])
		time.Sleep(2 * time.Second)
		if i == (len(lines) - 1) {
			i = -1
		}
	}

}
