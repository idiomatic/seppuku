package main

import (
	"os"
	"path/filepath"

	"github.com/idiomatic/seppuku"
)

func main() {
	for _, arg := range os.Args[1:] {
		watching := filepath.SplitList(arg)
		go seppuku.Seppuku(watching)
	}
	select {}
}
