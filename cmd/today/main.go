package main

import (
	"fmt"
	"time"

	"github.com/jchenry/jchenry/pkg/arvelie"
)

func main() {
	a := arvelie.FromDate(time.Now())
	fmt.Println(a)
}
