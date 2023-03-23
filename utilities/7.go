package main

import (
	"fmt"
	"time"
)

func main() {
	//var or func(channels ...<-chan interface{}) <-chan interface{}
	//Пример использования функции:
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()

	test := merge(
		sig(20*time.Second),
		sig(50*time.Second),
		sig(12*time.Second),
		sig(6*time.Second),
		sig(1*time.Second),
	)
	<-test
	fmt.Printf("done after %v", time.Since(start))
}

func merge(cs ...<-chan interface{}) <-chan interface{} {
	test := make(chan interface{})
	go func() {
		for {
			for i := range cs {
				select {
				case <-cs[i]:
					close(test)
					return
				default:
					continue
				}
			}
		}
	}()

	return test
}
