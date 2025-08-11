package main

import (
	"os"
	"time"

	clockface "github.com/and-volkov/learn-go-with-tests.study/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
