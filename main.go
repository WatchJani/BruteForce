package main

import "fmt"

func main() {
	fmt.Println(BruteForceChatGPT("Janko"))
}

func BruteForce(password string) string {
	list := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	block := make([]byte, 1, 4)
	indexSaver := make([]int, 1, 4)

	for {
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

				block = append(block, list[0])     //Dont touch
				indexSaver = append(indexSaver, 0) //Dont touch
			}
		}

		block[0] = list[indexSaver[0]]

		indexSaver[0]++

		if string(block) == password {
			return string(block)
		}
	}
}

func BruteForceChatGPT(password string) string {
	list := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	block := make([]byte, 1, 4)
	indexSaver := make([]int, 1, 4)

	for {
		for i := range block {
			block[i] = list[indexSaver[i]]
		}

		if string(block) == password {
			return string(block)
		}

		indexSaver[0]++

		for i := 0; i < len(indexSaver); i++ {
			if indexSaver[i] < len(list) {
				break
			}

			indexSaver[i] = 0
			if i+1 < len(indexSaver) {
				indexSaver[i+1]++
			} else {
				indexSaver = append(indexSaver, 0)
				block = append(block, list[0])
			}
		}
	}
}
