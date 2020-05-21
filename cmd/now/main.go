package main

import (
	"fmt"
	"time"

	"github.com/jchenry/libs/neralie"
)

func main() {
	a := neralie.FromTime(time.Now())
	fmt.Println(a)
}
