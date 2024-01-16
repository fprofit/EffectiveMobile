package main

import (
	"fmt"

	"github.com/fprofit/EffectiveMobile/internal/entry"
)

func main() {
	if err := entry.ComposeServer(); err != nil {
		fmt.Println(err)
	}

}
