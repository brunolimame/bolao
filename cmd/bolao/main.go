package main

import (
	"bolao/configs"
	"fmt"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(config.DBDriver)
}
