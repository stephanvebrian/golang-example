package main

import (
	"fmt"
	"sync"
)

func manipulateDataWithoutDefer(mutex *sync.Mutex, data *int) {
	mutex.Lock()

	// Modify the data
	*data += 10
	fmt.Println("Data modified inside manipulateData function:", *data)
	mutex.Unlock()
}

func manipulateDataWithDefer(mutex *sync.Mutex, data *int) {
	mutex.Lock()
	defer mutex.Unlock()

	*data += 10
	fmt.Println("Data modified inside manipulateData function:", *data)
}

// defer is basicaly a statement to assign a statement that will be run after all
// statement inside a function scope is run,
// in this case i will gave an example the sampe operation, with and without defer
func main() {
	var mutex sync.Mutex
	data := 5

	// Call manipulateDataWithoutDefer
	fmt.Printf("Data current value: %d\n", data)
	manipulateDataWithoutDefer(&mutex, &data)
	fmt.Println("Data value after manipulateDataWithDefer function call:", data)

	// Call manipulateDataWithDefer with the mutex and data
	data = 5
	fmt.Printf("Data current value: %d\n", data)
	manipulateDataWithDefer(&mutex, &data)
	fmt.Println("Data value after manipulateDataWithDefer function call:", data)
}
