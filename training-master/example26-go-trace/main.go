package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

func Sum(numbers ...int) int {

	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SumParallel(numbers ...int) int {
	mid := len(numbers) / 2

	ch := make(chan int)
	go func() { ch <- Sum(numbers[:mid]...) }()
	go func() { ch <- Sum(numbers[mid:]...) }()

	total := <-ch + <-ch
	return total
}

var foo = []int{4, 5, 6, 1, 3, 2, 8, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()
	fmt.Println(runtime.NumCPU())
	fmt.Println(Sum(foo...))
	fmt.Println(SumParallel(foo...))
}
