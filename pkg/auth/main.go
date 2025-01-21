package main

import (
	"auth/cmd"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		cmd.StartGRPCServer()

	}()

	go func() {
		defer wg.Done()
		go cmd.StartAPI()

	}()

	wg.Wait()
}
