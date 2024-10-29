package main

func main() {
	Combination(1254)
	// fmt.Println(BruteForceChatGPT("Janko"))
}

func Combination(code int) {
	
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

				block = append(block, list[0])
				indexSaver = append(indexSaver, 0)
			}
		}

		block[0] = list[indexSaver[0]]

		indexSaver[0]++

		if string(block) == password {
			return string(block)
		}
	}
}
