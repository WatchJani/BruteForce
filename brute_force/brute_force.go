package brute_force

const Characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

type BruteForce struct {
	responseCh chan Response
}

type Response struct {
	iteration int
	password  string
}

func (r *Response) GetIteration() int {
	return r.iteration
}

func (r *Response) GetPassword() string {
	return r.password
}

func (bf *BruteForce) GetResponseCh() Response {
	return <-bf.responseCh
}

func New() *BruteForce {
	return &BruteForce{
		responseCh: make(chan Response),
	}
}

func FindCombination(position int) []int {
	if position == 0 {
		return []int{0}
	}

	alphabetSize := len(Characters)

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
		result[i] = position % alphabetSize
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
func (bf *BruteForce) Worker(hash string, indexSaver []int, c *CancelManager) {
	index, block := 0, make([]byte, len(indexSaver))
	for index, value := range indexSaver {
		block[index] = Characters[value]
	}

	for ; index < 10_000_000_000 && c.GetState(); index++ {
		for f := 0; indexSaver[f] > len(Characters)-1; f++ {
			indexSaver[f] = 0

			if len(indexSaver) > f+1 { //increase first next character
				indexSaver[f+1]++
				if indexSaver[f+1] > len(Characters)-1 {
					block[f+1] = Characters[0]
					continue
				}

				block[f+1] = Characters[indexSaver[f+1]]
			} else { //add new character
				block = append(block, Characters[0])
				indexSaver = append(indexSaver, 0)
			}
		}

		block[0] = Characters[indexSaver[0]]

		indexSaver[0]++

		if string(block) == hash {
			c.CancelFn()
			break
		}
	}

	bf.responseCh <- Response{
		iteration: index,
		password:  string(block),
	}
}
