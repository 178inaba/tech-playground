package main

import (
	"github.com/178inaba/tech-playground/gocc"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gocc.Analyzer)
}
