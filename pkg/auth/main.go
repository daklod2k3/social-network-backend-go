package main

import (
	"auth/cmd"
	"auth/internal/global"
	global2 "shared/global"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	global.InitGlobal()

	global2.InitGlobal(&global2.Type{
		Config: global.Config,
		Logger: global2.Logger,
	})

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
