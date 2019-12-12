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
	doTheJob(strings.Split(scanner.Text(), ","))
}

func doTheJob(valStr []string) {
	var val []int

	for i := 0; i < len(valStr); i++ {
		v, err := strconv.Atoi(valStr[i])
		if err != nil {
			log.Println(err)
		}
		val = append(val, v)
	}

	for i := 0; i < len(val); {
		var A, B, C, code, X, Y, Z int

		if val[i] == 99 {
			code = 99
		} else {
			code = val[i] % 100
			num := val[i] / 100

			C = num % 10
			num /= 10
			B = num % 10
			num /= 10
			A = num

			Z = val[i+3]
			Y = val[i+2]
			X = val[i+1]
			
			if C == 1 {
				X = i + 1
			}
			if B == 1 {
				Y = i + 2
			}
			if A == 1 {
				Z = i + 3
			}
		}

		switch code {
		case 1:
			val[Z] = val[X] + val[Y]
			i += 4
		case 2:
			val[Z] = val[X] * val[Y]
			i += 4
		case 3:
			fmt.Print("Input: ")
			var input int
			fmt.Scanln(&input)
			val[val[i+1]] = input
			i += 2
		case 4:
			fmt.Println("Output:", val[val[i+1]])
			i += 2
		case 99:
			return
		}
	}
}
