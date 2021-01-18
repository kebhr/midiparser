package main

import (
	"fmt"
)

func main() {
	f, err := os.Open("midi1.mid")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
}
