package bash

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tkdeng/goutil"
)

var cliReader = bufio.NewReader(os.Stdin)

// InputText prompts the user for text input and returns the entered string
func InputText(msg string) string {
	fmt.Printf("\n%s ", msg)

	input, err := cliReader.ReadString('\n')
	if err != nil {
		return ""
	}

	fmt.Println("")

	return goutil.Clean(strings.TrimSpace(input))
}

// InputSelect prompts the user to select from a list of options and returns the selected index
func InputSelect(msg string, opts ...string) int {
	fmt.Println("")

	for i, v := range opts {
		fmt.Printf("[%d] %s\n", i, v)
	}

	fmt.Printf("\n%s ", msg)

	input, err := cliReader.ReadString('\n')
	if err != nil {
		return 0
	}

	fmt.Println("")

	sel, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || sel < 0 || sel > len(opts)-1 {
		return 0
	}

	return sel
}

// InputYN prompts the user for a yes/no input and returns true or false
func InputYN(msg string, def ...bool) bool {
	if len(def) == 0 {
		fmt.Printf("%s (y/n)? ", msg)
	} else if def[0] {
		fmt.Printf("%s (Y/n)? ", msg)
	} else {
		fmt.Printf("%s (y/N)? ", msg)
	}

	input, err := cliReader.ReadString('\n')
	if err != nil {
		if len(def) != 0 {
			return def[0]
		}
		return false
	}

	switch strings.ToLower(strings.TrimSpace(input)) {
	case "y", "yes", "1", "t", "true":
		return true
	case "n", "no", "0", "f", "false":
		return false
	}

	if len(def) != 0 {
		return def[0]
	}
	return false
}
