package main

import (
	"github.com/sadlil/tools/cmd/gouberfmt/analyzers/interfacepointer"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		interfacepointer.Analyzer,
	)
}
