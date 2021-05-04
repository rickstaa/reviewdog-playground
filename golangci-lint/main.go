package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("255")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(i)
}
