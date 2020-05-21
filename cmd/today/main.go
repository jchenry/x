package main

import (
	"fmt"
	"time"

	"github.com/jchenry/libs/arvelie"
)

func main() {
	a := arvelie.FromDate(time.Now())
	fmt.Println(a)
}
