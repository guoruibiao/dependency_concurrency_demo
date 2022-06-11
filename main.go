package main

import (
	"fmt"
	"github.com/guoruibiao/hello/prc"
	"time"
)

func main() {
	fmt.Println("starting...")
	startTime := time.Now().Unix()
	prc.Process()
	seconds := time.Now().Unix() - startTime
	fmt.Println("done! cost=", seconds)
}
