package main

import (
	"fmt"
)

func main() {
	fmt.Println(BruteForce("Janko"))
	fmt.Println(findCombination(10_000_000_000))
}

func findCombination(position int) []int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetSize := len(alphabet)

	position -= 1

	var length int
	for {
		numCombinations := pow(alphabetSize, length+1)
		if position < numCombinations {
			break
		}
		position -= numCombinations
		length++
	}

	result := make([]int, length+1)

	for i := length; i >= 0; i-- {
		result[i] = int(alphabet[position%alphabetSize])
		position /= alphabetSize
	}

	for index := 0; index < len(result)/2; index++ {
		result[index], result[len(result)-1-index] = result[len(result)-1-index], result[index]
	}

	return result
}

func pow(a, b int) int {
	result := 1
	for b > 0 {
		result *= a
		b--
	}
	return result
}

// 10_000_000_000
func BruteForce(password string) string {
	list := "abcdefghijklmnopqrstuvwxyz"

	block := make([]byte, 1, 4)
	indexSaver := make([]int, 1, 4)

	for index := 0; index < 1105; index++ {
		for f := 0; indexSaver[f] > len(list)-1; f++ {
			indexSaver[f] = 0

			if len(indexSaver) > f+1 { //increase first next character
				indexSaver[f+1]++
				if indexSaver[f+1] > len(list)-1 {
					block[f+1] = list[0]
					continue
				}

				block[f+1] = list[indexSaver[f+1]]
			} else { //add new character
				block = append(block, list[0])
				indexSaver = append(indexSaver, 0)
			}
		}

		block[0] = list[indexSaver[0]]

		indexSaver[0]++

		fmt.Println(index, string(block))
		if string(block) == password {
			return string(block)
		}
	}

	return ""
}
