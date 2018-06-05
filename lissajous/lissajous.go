package main

import (
	"os"

	liss "github.com/jfixby/gotest/lissajous/lissfunc"
)

func main() {
	liss.Lissajous(os.Stdout)
}
