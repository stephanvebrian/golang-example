package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

// utility function, dont have any to do with the example,
// what you need to know is it will generating random string based on the length we pass
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result += string(charset[index.Int64()])
	}
	return result
}

// generating million of string using sync nature
// in every 100_000 process, we will sleep it around 3s
func generateMillionOfStrings() []string {
	strs := []string{}

	for i := 1; i <= 1_000_000; i++ {
		lengthGenerated := 25

		strs = append(strs, generateRandomString(lengthGenerated))

		if i%100_000 == 0 {
			time.Sleep(3 * time.Second)
		}
	}

	return strs
}

// generating million of string using async nature (goroutines)
// in this example i used maximum of 1_000 goroutines spawn, to limit the cpu usage
func generateMillionOfStringsGoroutine() []string {
	strs := []string{}
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for i := 1; i <= 1_000_000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			lengthGenerated := 25
			generatedStr := generateRandomString(lengthGenerated)
			if i%100_000 == 0 {
				time.Sleep(3 * time.Second)
			}

			mutex.Lock()
			strs = append(strs, generatedStr)
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	return strs
}

// in this example i will show you the difference processing long process using goroutines (async process)
func main() {
	fmt.Println("Starting the generation process 1 with sync nature...")
	startProcess1 := time.Now()
	process1 := generateMillionOfStrings()
	fmt.Printf("Process 1")
	fmt.Printf("  | total length generated: %d\n", len(process1))
	fmt.Printf("  | time processing: %+v (seconds) \n", time.Since(startProcess1).Seconds())

	fmt.Println("Starting the generation process 2 with async nature...")
	startProcess2 := time.Now()
	process2 := generateMillionOfStringsGoroutine()
	fmt.Printf("Process 2")
	fmt.Printf("  | total length generated: %d\n", len(process2))
	fmt.Printf("  | time processing: %+v (seconds) \n", time.Since(startProcess2).Seconds())
}
