package main

import "fmt"

func main() {
	data := make([]byte, 20)
	BruteForce(&data)
}

func BruteForce(data *[]byte) {
	list := "abc"

	block := make([]byte, 1, 4)
	indexSaver := make([]int, 1, 4)

	for index := 0; index < 12; index++ {
		for f := 0; indexSaver[f] > len(list)-1; f++ {
			if len(indexSaver) > f+1 {
				indexSaver[f+1]++
				block[f+1] = list[indexSaver[f+1]]
				indexSaver[f] = 0
			} else {
				block = append(block, list[0]) //Dont touch
				indexSaver[f] = 0
				indexSaver = append(indexSaver, 0) //Dont touch
			}
		}

		block[0] = list[indexSaver[0]]

		indexSaver[0]++

		fmt.Println(string(block))
	}
}
