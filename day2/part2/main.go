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

	var i, j int64

	for i = 0; i < 100; i++ {
		for j = 0; j < 100; j++ {
			data := strings.Split(scanner.Text(), ",")
			if result, err := doTheJob(data, i, j); result == 19690720 {
				fmt.Println("Anwser is: ", 100*i+j, "i = ", i, "j = ", j)
				break
			} else if err != nil {
				log.Println(err)
			}
		}
	}
}

func doTheJob(valStr []string, a, b int64) (int64, error) {
	var val []int64

	for i := 0; i < len(valStr); i++ {
		v, err := strconv.ParseInt(valStr[i], 10, 64)
		if err != nil {
			log.Println(err)
		}
		val = append(val, v)
	}

	val[1] = a
	val[2] = b

	// fmt.Println(a,b)

	for i := 0; i < len(val); i += 4 {
		switch val[i] {
		case 1:
			val[val[i+3]] = val[val[i+1]] + val[val[i+2]]
			break
		case 2:
			val[val[i+3]] = val[val[i+1]] * val[val[i+2]]
			break
		case 99:
			// log.Println(val)
			return val[0], nil
		}
	}

	return -1, nil
}
