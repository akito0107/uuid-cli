package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	u, _ := uuid.NewRandom()
	fmt.Println(u)
}
