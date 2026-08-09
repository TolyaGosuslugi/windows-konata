package display

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Err(err error) {
	fmt.Printf("%s", color.RedString("Error: "+err.Error()))
	os.Exit(1)
}
