package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("file.txt")
	if err != nil {
		log.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Scan()

	rangeStr := strings.Split(scanner.Text(), "-")

	numberOfResults, err := findResults(rangeStr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("There are:", numberOfResults, "results")
	}
}

func findResults(rangeStr []string) (int, error) {
	start, err := strconv.Atoi(rangeStr[0])
	if err != nil {
		return -1, err
	}
	end, err := strconv.Atoi(rangeStr[1])
	if err != nil {
		return -1, err
	}

	var results int
	for i := start; i < end; i++ {
		inc := isInc(i)
		if inc {
			if isAdj(i) {
				results++
			}
		}
	}

	return results, nil
}

func isAdj(num int) bool {
	var curr, counter int
	curr = num % 10
	for {
		num /= 10
		prev := curr
		curr = num % 10

		if prev == curr {
			counter++
			if (num/10)%10 != prev && counter < 2 {
				return true
			}
		} else {
			counter = 0
		}

		if num == 0 {
			break
		}
	}
	return false
}

func isInc(num int) bool {
	var curr int
	curr = num % 10
	num /= 10
	for i := 0; ; i++ {
		last := curr
		curr = num % 10
		num /= 10
		if last < curr {
			return false
		}
		if num == 0 {
			break
		}
	}
	return true
}
