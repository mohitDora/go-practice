package main

import (
	"fmt"
	"sync"
	"time"
)

func once_() {
	/*
	   once is a type that has a method Do(f func()) that takes a function as an argument and ensures that the function is executed only once, even if Do is called multiple times.
	*/

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			config := getConfig()
			fmt.Printf("Config: %v\n", *config)
		}()
	}
	wg.Wait()

}

type Config struct {
	ApiKey string
	Port   int
}

var (
	appConfig *Config
	once      sync.Once
)

func loadConfig() {
	fmt.Println("Loading config")
	appConfig = &Config{ApiKey: "123456", Port: 8080}
	time.Sleep(2 * time.Second)
	fmt.Println("Config loaded")
}

func getConfig() *Config {
	once.Do(loadConfig)
	return appConfig
}
