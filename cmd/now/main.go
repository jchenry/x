package main

import (
	"fmt"
	"time"

	"github.com/jchenry/jchenry/pkg/neralie"
)

func main() {
	a := neralie.FromTime(time.Now())
	fmt.Println(a)
}
