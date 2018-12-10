package main

import (
	"fmt"
)

var squares = [...]int{4, 9, 16, 25, 36, 49, 64, 81}

func main() {
	fmt.Println(getSquareSums(5))
}

func getPairs(n int) map[int][]int {
	pairs := make(map[int][]int)
	for i := 1; i <= n; i++ {
		var ints []int
		for j := 1; j <= n; j++ {
			if i == j {
				continue
			}
			for _, square := range squares {
				if (i + j) == square {
					ints = append(ints, j)
				}
			}
		}
		pairs[i] = ints
	}
	return pairs
}

func getNextFromPair(current int, values []int) int {
	for _, value := range values {
		if current < value {
			return value
		}
	}
	return 0
}

func isFree(value int, solution []int) bool {
	for _, s := range solution {
		if value == s {
			return false
		}
	}
	return true
}

func revert(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func getSquareSums(n int) []int {
	pairs := getPairs(n)
	var solution []int
	var previous int
	firstNumber := 1
	current := firstNumber
	next := getNextFromPair(current, pairs[current])
	for firstNumber <= n && len(solution) != n {
		if next != 0 && isFree(next, solution) {
			solution = append(solution, next)
			previous = next
			next = getNextFromPair(1, pairs[next])
		} else {
			if next == 0 {
				next = solution[len(solution)-1]
				solution = solution[:len(solution)-1]
				if len(solution) > 0 {
					previous = solution[len(solution)-1]
				} else {
					firstNumber++
					previous = firstNumber
					solution = append(solution, previous)
					next = 0
				}
			}
			next = getNextFromPair(next, pairs[previous])
		}
	}
	if firstNumber > n {
		return nil
	}
	return revert(solution)
}
