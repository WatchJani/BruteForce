package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const Characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	bf := New()
	start := time.Now()

	for index := range runtime.NumCPU() {
		go bf.Worker(index)
	}

	//controlling system closing
	go func() {
		<-sigs

		fmt.Println(time.Since(start))
		os.Exit(1)
	}()

	//2m22s
	go func() {
		for index := 0; index*10_000_000_000 < 1_000_000_000_000; index++ {
			fmt.Println(index * 10_000_000_000)
			bf.dataStreamCh <- DataStream{
				hash:       "JankoKondic",
				startPoint: findCombination(index * 10_000_000_000),
			}
		}

		fmt.Println(time.Since(start))
		os.Exit(1)
	}()

	fmt.Println(<-bf.responseCh)
}

type BruteForce struct {
	dataStreamCh chan DataStream
	responseCh   chan string
}

func New() BruteForce {
	return BruteForce{
		dataStreamCh: make(chan DataStream),
		responseCh:   make(chan string),
	}
}

type DataStream struct {
	hash       string //password
	startPoint []int
}

func findCombination(position int) []int {
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
func (bf *BruteForce) Worker(workerIndex int) {
	for {
		store := <-bf.dataStreamCh

		indexSaver := store.startPoint
		block := make([]byte, len(indexSaver))
		for index, value := range indexSaver {
			block[index] = Characters[value]
		}

		for index := 0; index < 10_000_000_000; index++ {
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

			if string(block) == store.hash {
				bf.responseCh <- string(block)
			}
		}
	}
}
