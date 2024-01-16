package main

import (
	"fmt"
	"time"

	"github.com/fprofit/EffectiveMobile/internal/entry"
)

func main() {
	time.Sleep(5 * time.Second)
	if err := entry.ComposeServer(); err != nil {
		fmt.Println(err)
	}

}
