package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for range 5 {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func await(awaitGroup *sync.WaitGroup, funcToRun func(string), params ...string) {
	awaitGroup.Add(len(params))
	
	for _, param_ := range params {
		go func(){
			funcToRun(param_)
			awaitGroup.Done()
		}()
	}

	awaitGroup.Wait()
}

func main() {
	var waitGroup sync.WaitGroup;

	await(&waitGroup, write, "Hello", "World")

}