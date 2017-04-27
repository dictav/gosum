package main

import (
	"os"

	"github.com/haya14busa/gosum/checker"
	"honnef.co/go/tools/lint/lintutil"
)

func main() {
	lintutil.ProcessArgs("gosumcheck", checker.NewChecker(), os.Args[1:])
}
