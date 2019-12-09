package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var counter int
	data, err := os.Open("file.txt")
	if err != nil {
		log.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		counter += doTheMath(scanner.Text())
	}
	fmt.Println(counter)
}

func doTheMath(read string) int {
	value, err := strconv.Atoi(read)
	if err != nil {
		log.Println(err)
	}
	floatVal := float64(value) / float64(3)
	floatVal = math.Floor(floatVal)
	value = int(floatVal)
	return value - 2
}
