package main

import (
	"fmt"

	"github.com/joshsoftware/curem/config"
)

func main() {
	fmt.Printf("%+v", config.Db)
	fmt.Println()
	fmt.Printf("%+v", config.Session)
	fmt.Println()
}
