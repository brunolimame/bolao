package main

import (
	"bolao/configs"
	"fmt"
)

func main() {
	configs := configs.NewConfig()
	fmt.Println(configs.GetDBDriver())
}
