package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	a := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	b := strings.Split(scanner.Text(), ",")

	dist := findClosestIntersection(a, b)
	fmt.Println("Shortest Manhattan distance to an intersection is:", dist)
}

func findClosestIntersection(a, b []string) int {
	var A, B [2]int
	var aP, bP, is [][2]int
	var counter int
	for i := 0; i < len(a); i++ {
		lenA, _ := strconv.Atoi(a[i][1:])
		lenB, _ := strconv.Atoi(b[i][1:])
		dirA := a[i][:1]
		dirB := b[i][:1]

		for j := 0; j < lenA; j++ {
			A = newPos(A, dirA)
			aP = append(aP, A)
		}

		for j := 0; j < lenB; j++ {
			B = newPos(B, dirB)
			bP = append(bP, B)
		}
	}

	for i := 0; i < len(aP); i++ {
			for j := 0; j < len(bP); j++ {
				if detected := collisionDetect(aP[i], bP[j]); detected {
					is = append(is, aP[i])
					counter++
				}
			}
	}
	fmt.Println("Found", counter, "intersections")

	var manLen []int
	min := math.MaxInt64
	for i := 0; i < len(is); i++ {
		len := calcManhattan(is[i])
		manLen = append(manLen, len)
		if len < min {
			min = len
		}
	}
	return min
}

func calcManhattan(X [2]int) int {
	return abs(X[0]) + abs(X[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func newPos(X [2]int, x string) [2]int {

	switch x {
	case "L":
		X[0]--
	case "R":
		X[0]++
	case "U":
		X[1]++
	case "D":
		X[1]--
	}
	return X
}

func collisionDetect(A, B [2]int) bool {
	if A == B {
		return true
	}
	return false
}
