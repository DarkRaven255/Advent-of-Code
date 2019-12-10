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
	fmt.Println("Shortest steps to an intersection is:", dist)

}

func findClosestIntersection(a, b []string) int {
	var A, B [3]int
	var aP, bP, is [][3]int
	var counter, stepsA, stepsB int
	for i := 0; i < len(a); i++ {
		lenA, _ := strconv.Atoi(a[i][1:])
		lenB, _ := strconv.Atoi(b[i][1:])
		dirA := a[i][:1]
		dirB := b[i][:1]
		for j := 0; j < lenA; j++ {
			A = newPos(A, dirA)
			stepsA++
			A[2] = stepsA
			aP = append(aP, A)
			
		}
		for j := 0; j < lenB; j++ {
			B = newPos(B, dirB)
			stepsB++
			B[2] = stepsB
			bP = append(bP, B)
		}
	}

	for i := 0; i < len(aP); i++ {
		for j := 0; j < len(bP); j++ {
			if detected := collisionDetect(aP[i], bP[j]); detected {
				aP[i][2] += bP[j][2]
				is = append(is, aP[i])
				counter++
			}
		}
	}

	fmt.Println("Found", counter, "intersections")
	min := math.MaxInt64
	for i := 0; i < len(is); i++ {
		if is[i][2] < min {
			min = is[i][2]
		}
	}
	return min
} 

func newPos(X [3]int, x string) [3]int {

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

func collisionDetect(A, B [3]int) bool {
	if A[0] == B[0] && A[1] == B[1] {
		return true
	}
	return false
}
