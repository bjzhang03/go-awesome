package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1 Finish the Works!")
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2 Finish the Works!")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("OK, Finish all the Works, Release!")
}
