package main

import (
	"fmt"

	"github.com/zogot/go-simple-example/pkg/anothercomponent"
)

func main() {
	mRepo := anothercomponent.NewMemoryRepository()

	fmt.Println(mRepo.FindAllOfComponent())
}
