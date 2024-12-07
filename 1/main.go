package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("input")
	input := string(data)

	//First column matches
	regex1, _ := regexp.Compile(`(?m:^\d{5})`)
	firstColMatch := regex1.FindAllString(input, -1)
	firstCol := make([]int, 0, len(firstColMatch))

	//Second column matches
	regex2, _ := regexp.Compile(`(?m:\d{5}$)`)
	secondColMatch := regex2.FindAllString(input, -1)
	secondCol := make([]int, 0, len(secondColMatch))

	// len(firstColMatch) = len(secondColMatch)
	for i, _ := range firstColMatch {
		el, err := strconv.Atoi(firstColMatch[i])
		if err != nil {
			fmt.Println(err)
		}
		firstCol = append(firstCol, el)

		el, err = strconv.Atoi(secondColMatch[i])
		if err != nil {
			fmt.Println(err)
		}
		secondCol = append(secondCol, el)
	}
	sort.Ints(firstCol)
	sort.Ints(secondCol)

	println(totalDistance(firstCol, secondCol))
	println(similarityScore(firstCol, secondCol))
}

func totalDistance(firstCol []int, secondCol []int) int {
	var totalDistance int
	for i, _ := range firstCol {
		if firstCol[i] > secondCol[i] {
			totalDistance += firstCol[i] - secondCol[i]
		} else {
			totalDistance += secondCol[i] - firstCol[i]
		}
	}
	return totalDistance
}
func similarityScore(firstCol []int, secondCol []int) int {
	htable := make(map[int]int)
	var score int

	for i, v := range firstCol {
		for i1, _ := range secondCol {
			if firstCol[i] == secondCol[i1] {
				htable[v]++
			}
		}
	}
	for k, v := range htable {
		score += k * v
	}

	return score

	//// Debug
	// for i, v := range firstCol {
	// fmt.Printf("I:%v, Val: %v, H Val: %v\n", i, v, htable[v])
	// for i1, _ := range secondCol {
	// if firstCol[i] == secondCol[i1] {
	// htable[v]++
	// fmt.Printf("\tI1:%v, I2:%v, 1Ar: %v, 2Ar: %v, H Val: %v\n",
	// i, i1, firstCol[i], secondCol[i1], htable[v])
	// }
	// }
	// }
}
