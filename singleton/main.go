package main

import (
	"fmt"
	"sync"

	"github.com/vindecodex/Aryzath/singleton/singleton"
)

var once sync.Once // thread safe
var singletonInstance *singleton.Singleton

func getInstance() *singleton.Singleton {
	once.Do(func() {
		singletonInstance = &singleton.Singleton{}
		fmt.Println("Created Instance")
	})
	return singletonInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance() // even if we run go routines it only create 1 instance
	}
	fmt.Scanln()
}
