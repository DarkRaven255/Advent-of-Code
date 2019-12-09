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
	var counter, val int64
	data, err := os.Open("file.txt")
	if err != nil {
		log.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		val, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Println(err)
		}
		counter += doTheMath(val)
	}
	fmt.Println(counter)
}

func doTheMath(value int64) int64 {
	floatVal := float64(value) / float64(3)
	floatVal = math.Floor(floatVal)
	value = int64(floatVal) - 2
	if value > 0 {
		value += doTheMath(value)
	}

	if value <= 0 {
		return 0
	}

	return value
}
