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

	for index := 0; index < 8; index++ {
		if indexSaver[0] > len(list)-1 {
			block = append(block, list[0]) //Dont touch
			indexSaver[0] = 0
			indexSaver = append(indexSaver, 0) //Dont touch
		}

		block[0] = list[indexSaver[0]]

		indexSaver[0]++
		
		fmt.Println(string(block))
	}
}
