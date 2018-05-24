package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	//带缓存的chan
	naturals := make(chan int, 200)
	fmt.Println(cap(naturals))
	//不带缓存的chan
	squares := make(chan int)
	fmt.Println(cap(squares))
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

}
