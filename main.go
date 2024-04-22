package main

import (
	"fmt"
	"os"

	explorer "github.com/h4vismat/rangure/internal/explorer"
)

func main() {
	fmt.Fprintf(os.Stdout, "rangure - file tagger\n")

	explorer.ShowFiles()

}
