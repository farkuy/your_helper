package main

import (
	"fmt"
	"log"
	"your_helper/internal/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
