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
	fmt.Println(doTheJob(strings.Split(scanner.Text(), ",")))
}

func doTheJob(valStr []string) (int64, error) {
	var val []int64

	for i := 0; i < len(valStr); i++ {
		v, err := strconv.ParseInt(valStr[i], 10, 64)
		if err != nil {
			log.Println(err)
		}
		val = append(val, v)
	}

	val[1] = 12
	val[2] = 2 

	for i := 0; i < len(val); i += 4 {
		switch val[i] {
		case 1:
			val[val[i+3]] = val[val[i+1]] + val[val[i+2]]
			break
		case 2:
			val[val[i+3]] = val[val[i+1]] * val[val[i+2]]
			break
		case 99:
			log.Println(val)
			return val[0], nil
		}
	}

	return -1, nil
}