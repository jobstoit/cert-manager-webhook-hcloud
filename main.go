package main

import (
	"fmt"
	"log"

	"github.com/jobstoit/tags/env"
)

// Config contains the configuration
type Config struct {
	GroupName string `env:"GROUP_NAME"`
}

func NewConfig() *Config {
	x := &Config{}
	if err := env.Parse(x); err != nil {
		log.Fatalf("Error parsing environment config: %v", err)
	}

	return x
}

func main() {
	fmt.Println("vim-go")
}
